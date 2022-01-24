package debian

import (
	"path/filepath"

	"github.com/packagefoundation/yap/constants"
	"github.com/packagefoundation/yap/utils"
)

type Project struct {
	Name       string
	Root       string
	MirrorRoot string
	BuildRoot  string
	Path       string
	Distro     string
	Release    string
}

func (p *Project) getBuildDir() (path string, err error) {
	path = filepath.Join(p.BuildRoot, p.Distro+"-"+p.Release)

	err = utils.MkdirAll(path)
	if err != nil {
		return
	}

	return
}

func (p *Project) Prep() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	keyPath := filepath.Join(p.Path, "..", "sign.key")
	exists, err := utils.Exists(keyPath)

	if err != nil {
		return
	}

	if exists {
		err = utils.CopyFile("", keyPath, buildDir, true)
		if err != nil {
			return
		}
	}

	err = utils.RsyncExt(p.Path, buildDir, ".deb")
	if err != nil {
		return
	}

	return
}

func (p *Project) Create() (err error) {
	buildDir, err := p.getBuildDir()
	if err != nil {
		return
	}

	err = utils.Exec("", "podman", "run", "--rm", "-t", "-v",
		buildDir+":/yap:Z", constants.DockerOrg+p.Distro+"-"+p.Release,
		"create", p.Distro+"-"+p.Release, p.Name)
	if err != nil {
		return
	}

	err = utils.Rsync(filepath.Join(buildDir, "apt"),
		filepath.Join(p.MirrorRoot, "apt"))
	if err != nil {
		return
	}

	return
}
