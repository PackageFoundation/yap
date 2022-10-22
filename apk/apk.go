package apk

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/packagefoundation/yap/pack"
	"github.com/packagefoundation/yap/utils"
)

type Apk struct {
	Pack   *pack.Pack
	apkDir string
}

func (a *Apk) convertApk() {
	for index, arch := range a.Pack.Arch {
		switch arch {
		case "all":
			a.Pack.Arch[index] = "any"
		case "amd64":
			a.Pack.Arch[index] = "x86_64"
		default:
			a.Pack.Arch[index] = arch
		}
	}
}

func (a *Apk) getDepends() error {
	var err error
	if len(a.Pack.MakeDepends) == 0 {
		return err
	}

	args := []string{
		"add",
	}
	args = append(args, a.Pack.MakeDepends...)

	err = utils.Exec("", "apk", args...)
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) getUpdates() error {
	err := utils.Exec("", "apk", "update")
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) createInstall() error {
	path := filepath.Join(a.apkDir, a.Pack.PkgName+".install")

	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file
	defer file.Close()

	// create new buffer
	writer := io.Writer(file)

	tmpl := template.New(".install")
	tmpl.Funcs(template.FuncMap{
		"join": func(strs []string) string {
			return strings.Trim(strings.Join(strs, ", "), " ")
		},
		"multiline": func(strs string) string {
			ret := strings.ReplaceAll(strs, "\n", "\n ")

			return strings.Trim(ret, " \n")
		},
	})

	template.Must(tmpl.Parse(postInstall))

	if err != nil {
		log.Fatal(err)
	}

	// DEBUG
	// err = tmpl.Execute(os.Stdout, a)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = tmpl.Execute(writer, a)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (a *Apk) createMake() error {
	path := filepath.Join(a.apkDir, "APKBUILD")
	file, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file
	defer file.Close()

	// create new buffer
	writer := io.Writer(file)

	tmpl := template.New("apkbuild")
	tmpl.Funcs(template.FuncMap{
		"join": func(strs []string) string {
			return strings.Trim(strings.Join(strs, ", "), " ")
		},
		"multiline": func(strs string) string {
			ret := strings.ReplaceAll(strs, "\n", "\n ")

			return strings.Trim(ret, " \n")
		},
	})

	template.Must(tmpl.Parse(specFile))

	if err != nil {
		log.Fatal(err)
	}

	// DEBUG
	// err = tmpl.Execute(os.Stdout, a)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = tmpl.Execute(writer, a)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (a *Apk) apkBuild() error {
	err := utils.ChownR(a.apkDir, "nobody", "nobody")
	if err != nil {
		return err
	}

	err = utils.ChownR(a.Pack.PackageDir, "nobody", "nobody")
	if err != nil {
		return err
	}

	err = utils.Exec(a.apkDir, "sudo", "-u", "root", "abuild-keygen", "-n", "-i", "-a")
	if err != nil {
		return err
	}

	err = utils.Exec(a.apkDir, "sudo", "-u", "root", "abuild", "-F", "-K")
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) Prep() error {
	err := a.getDepends()
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) Update() error {
	err := a.getUpdates()
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) makePackerDir() error {
	err := utils.ExistsMakeDir(a.apkDir)
	if err != nil {
		return err
	}

	err = utils.ExistsMakeDir(a.apkDir + "/pkg/" + a.Pack.PkgName)
	if err != nil {
		return err
	}

	return err
}

func (a *Apk) Build() ([]string, error) {
	a.apkDir = filepath.Join(a.Pack.Root, "apk")

	err := utils.RemoveAll(a.apkDir)
	if err != nil {
		return nil, err
	}

	err = a.makePackerDir()
	if err != nil {
		return nil, err
	}

	a.convertApk()

	err = a.createMake()
	if err != nil {
		return nil, err
	}

	err = a.createInstall()
	if err != nil {
		return nil, err
	}

	err = a.apkBuild()
	if err != nil {
		return nil, err
	}

	pkgs, err := utils.FindExt("/root/packages", ".apk")

	if err != nil {
		return nil, err
	}

	return pkgs, nil
}

func (a *Apk) Install() error {
	pkgs, err := utils.FindExt("/root/packages", ".apk")
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		if err := utils.Exec("", "sudo", "-u", "root", "apk", "add", "--allow-untrusted", pkg); err != nil {
			return err
		}
	}

	return nil
}
