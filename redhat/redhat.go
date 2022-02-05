package redhat

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/packagefoundation/yap/constants"
	"github.com/packagefoundation/yap/pack"
	"github.com/packagefoundation/yap/set"
	"github.com/packagefoundation/yap/utils"
)

type Redhat struct {
	Pack         *pack.Pack
	redhatDir    string
	buildDir     string
	buildRootDir string
	rpmsDir      string
	sourcesDir   string
	specsDir     string
	srpmsDir     string
}

func (r *Redhat) getDepends() (err error) {
	if len(r.Pack.MakeDepends) == 0 {
		return
	}

	args := []string{
		"-y",
		"install",
	}
	args = append(args, r.Pack.MakeDepends...)

	err = utils.Exec("", "yum", args...)

	if err != nil {
		return
	}

	return
}

func (r *Redhat) getFiles() (files []string, err error) {
	backup := set.NewSet()
	paths := set.NewSet()

	for _, path := range r.Pack.Backup {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}

		backup.Add(path)
	}

	output, err := utils.ExecOutput(r.Pack.PackageDir, "find", ".", "-printf", "%P\n")
	if err != nil {
		return
	}

	for _, path := range strings.Split(output, "\n") {
		if len(path) < 1 || strings.Contains(path, ".build-id") {
			continue
		}

		paths.Remove(filepath.Dir(path))
		paths.Add(path)
	}

	for pathInf := range paths.Iter() {
		path := pathInf.(string)

		if backup.Contains(path) {
			path = `%config "/` + path + `"`
		} else {
			path = `"/` + path + `"`
		}

		files = append(files, path)
	}

	return files, err
}

func (r *Redhat) createSpec(files []string) (err error) {
	path := filepath.Join(r.specsDir, r.Pack.PkgName+".spec")

	release := "%{?dist}"

	switch {
	case r.Pack.Distro == "amazon":
		release = ".amzn" + r.Pack.Release
	case r.Pack.Distro == "centos":
		release = ".el" + r.Pack.Release
	case r.Pack.Distro == "oracle":
		release = ".ol" + r.Pack.Release
	case r.Pack.Distro == "rocky":
		release = ".el" + r.Pack.Release
	}

	data := ""
	data += fmt.Sprintf("Name: %s\n", r.Pack.PkgName)
	data += fmt.Sprintf("Summary: %s\n", r.Pack.PkgDesc)
	data += fmt.Sprintf("Version: %s\n", r.Pack.PkgVer)
	data += fmt.Sprintf("Release: %s%s\n", r.Pack.PkgRel, release)
	data += fmt.Sprintf("Group: %s\n", ConvertSection(r.Pack.Section))
	data += fmt.Sprintf("URL: %s\n", r.Pack.URL)
	data += fmt.Sprintf("License: %s\n", r.Pack.License)
	data += fmt.Sprintf("Packager: %s\n", r.Pack.Maintainer)

	for _, pkg := range r.Pack.Provides {
		data += fmt.Sprintf("Provides: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Conflicts {
		data += fmt.Sprintf("Conflicts: %s\n", pkg)
	}

	for _, pkg := range r.Pack.Depends {
		data += fmt.Sprintf("Requires: %s\n", pkg)
	}

	for _, pkg := range r.Pack.MakeDepends {
		data += fmt.Sprintf("BuildRequires: %s\n", pkg)
	}

	data += "\n"
	data += "%global _build_id_links none\n"
	data += "%global _python_bytecompile_extra 0\n"
	data += "%global _python_bytecompile_errors_terminate_build 0\n"
	data += "%undefine __brp_python_bytecompile"
	data += "\n"

	if len(r.Pack.PkgDescLong) > 0 {
		data += "%description\n"
		for _, line := range r.Pack.PkgDescLong {
			data += line + "\n"
		}

		data += "\n"
	}

	data += "%install\n"
	data += fmt.Sprintf("rsync -a -A %s/ $RPM_BUILD_ROOT/\n",
		r.Pack.PackageDir)
	data += "\n"
	data += "%files\n"

	for _, line := range files {
		data += line + "\n"
	}

	data += "\n"

	if len(r.Pack.PreInst) > 0 {
		data += "%pre\n"
		for _, line := range r.Pack.PreInst {
			data += line + "\n"
		}

		data += "\n"
	}

	if len(r.Pack.PostInst) > 0 {
		data += "%post\n"
		for _, line := range r.Pack.PostInst {
			data += line + "\n"
		}

		data += "\n"
	}

	if len(r.Pack.PreRm) > 0 {
		data += "%preun\n"
		data += "if [[ \"$1\" -ne 0 ]]; then exit 0; fi\n"

		for _, line := range r.Pack.PreRm {
			data += line + "\n"
		}

		data += "\n"
	}

	if len(r.Pack.PostRm) > 0 {
		data += "%postun\n"
		data += "if [[ \"$1\" -ne 0 ]]; then exit 0; fi\n"

		for _, line := range r.Pack.PostRm {
			data += line + "\n"
		}
	}

	err = utils.CreateWrite(path, data)
	if err != nil {
		return
	}

	fmt.Println(data)

	return err
}

func (r *Redhat) rpmBuild() (err error) {
	err = utils.Exec(r.specsDir, "rpmbuild", "--define",
		"_topdir "+r.redhatDir, "-bb", r.Pack.PkgName+".spec")
	if err != nil {
		return
	}

	return
}

func (r *Redhat) Prep() (err error) {
	err = r.getDepends()
	if err != nil {
		return
	}

	return
}

func (r *Redhat) Update() (err error) {
	return
}

func (r *Redhat) makeDirs() (err error) {
	r.redhatDir = filepath.Join(r.Pack.Root, "redhat")
	r.buildDir = filepath.Join(r.redhatDir, "BUILD")
	r.buildRootDir = filepath.Join(r.redhatDir, "BUILDROOT")
	r.rpmsDir = filepath.Join(r.redhatDir, "RPMS")
	r.sourcesDir = filepath.Join(r.redhatDir, "SOURCES")
	r.specsDir = filepath.Join(r.redhatDir, "SPECS")
	r.srpmsDir = filepath.Join(r.redhatDir, "SRPMS")

	for _, path := range []string{
		r.redhatDir,
		r.buildDir,
		r.buildRootDir,
		r.rpmsDir,
		r.sourcesDir,
		r.specsDir,
		r.srpmsDir,
	} {
		err = utils.ExistsMakeDir(path)
		if err != nil {
			return
		}
	}

	return
}

func (r *Redhat) clean() (err error) {
	if !constants.CleanPrevious {
		return
	}

	pkgPaths, err := utils.FindExt(r.Pack.Home, ".rpm")
	if err != nil {
		return
	}

	match, ok := constants.ReleasesMatch[r.Pack.FullRelease]
	if !ok {
		fmt.Printf("redhat: Failed to find match for '%s'\n",
			r.Pack.FullRelease)

		return
	}

	for _, pkgPath := range pkgPaths {
		if strings.Contains(filepath.Base(pkgPath), match) {
			_ = utils.Remove(pkgPath)
		}
	}

	return
}

func (r *Redhat) copy(destination string) (err error) {
	archs, err := ioutil.ReadDir(r.rpmsDir)
	if err != nil {
		fmt.Printf("redhat: Failed to find rpms from '%s'\n",
			r.rpmsDir)

		return
	}

	for _, arch := range archs {
		rpmDestination := filepath.Join(r.Pack.Home, destination)
		err = utils.ExistsMakeDir(rpmDestination)
		if err != nil {
			return
		}
		err = utils.CopyFiles(filepath.Join(
			r.rpmsDir,
			arch.Name(),
		), rpmDestination, false)
		if err != nil {
			return
		}
	}

	return
}

func (r *Redhat) remDirs() {
	os.RemoveAll(r.redhatDir)
}

func (r *Redhat) Build(outputDir string) ([]string, error) {
	err := r.makeDirs()
	if err != nil {
		return nil, err
	}

	defer r.remDirs()

	files, err := r.getFiles()
	if err != nil {
		return nil, err
	}

	err = r.createSpec(files)
	if err != nil {
		return nil, err
	}

	err = r.rpmBuild()
	if err != nil {
		return nil, err
	}

	err = r.clean()
	if err != nil {
		return nil, err
	}

	err = r.copy(outputDir)
	if err != nil {
		return nil, err
	}

	pkgs, err := utils.FindExt(r.Pack.Home, ".rpm")
	if err != nil {
		return nil, err
	}

	r.remDirs()
	err = r.makeDirs()

	if err != nil {
		return nil, err
	}

	return pkgs, err
}

func (r *Redhat) Install() error {
	pkgs, err := utils.FindExt(r.Pack.Home, ".rpm")
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		absPath, err := filepath.Abs(pkg)
		if err != nil {
			return err
		}

		if err := utils.Exec("", "yum", "install", "-y", absPath); err != nil {
			return err
		}
	}

	return nil
}
