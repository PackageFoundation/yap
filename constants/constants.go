package constants

import (
	"strings"

	"github.com/packagefoundation/yap/set"
)

const (
	DockerOrg = "yap/"
)

var (
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorWhite  = "\033[37m"
)

var (
	Releases = [...]string{
		"archlinux",
		"amazonlinux-1",
		"amazonlinux-2",
		"fedora-32",
		"fedora-33",
		"centos-7",
		"centos-8",
		"debian-jessie",
		"debian-stretch",
		"debian-buster",
		"oraclelinux-7",
		"oraclelinux-8",
		"ubuntu-trusty",
		"ubuntu-xenial",
		"ubuntu-bionic",
		"ubuntu-eoan",
		"ubuntu-focal",
		"ubuntu-groovy",
	}
	ReleasesMatch = map[string]string{
		"archlinux":      "",
		"amazonlinux-1":  ".amzn1.",
		"amazonlinux-2":  ".amzn2.",
		"fedora-32":      ".fc32.",
		"fedora-33":      ".fc33.",
		"centos-7":       ".el7.",
		"centos-8":       ".el8.",
		"debian-jessie":  "jessie",
		"debian-stretch": "stretch",
		"debian-buster":  "buster",
		"oraclelinux-7":  ".ol7.",
		"oraclelinux-8":  ".ol8.",
		"ubuntu-trusty":  "trusty",
		"ubuntu-xenial":  "xenial",
		"ubuntu-bionic":  "bionic",
		"ubuntu-eoan":    "eoan",
		"ubuntu-focal":   "focal",
		"ubuntu-groovy":  "groovy",
	}
	DistroPack = map[string]string{
		"archlinux":   "pacman",
		"amazonlinux": "redhat",
		"fedora":      "redhat",
		"centos":      "redhat",
		"debian":      "debian",
		"oraclelinux": "redhat",
		"ubuntu":      "debian",
	}
	Packagers = [...]string{
		"apt",
		"pacman",
		"yum",
	}

	ReleasesSet    = set.NewSet()
	Distros        = []string{}
	DistrosSet     = set.NewSet()
	DistroPackager = map[string]string{}
	PackagersSet   = set.NewSet()
	CleanPrevious  = false
)

func init() {
	for _, release := range Releases {
		ReleasesSet.Add(release)
		distro := strings.Split(release, "-")[0]
		Distros = append(Distros, distro)
		DistrosSet.Add(distro)
	}

	for _, distro := range Distros {
		packager := ""

		switch DistroPack[distro] {
		case "debian":
			packager = "apt"
		case "pacman":
			packager = "pacman"
		case "redhat":
			packager = "yum"
		default:
			panic("Failed to find packager for distro")
		}

		DistroPackager[distro] = packager
	}

	for _, packager := range Packagers {
		PackagersSet.Add(packager)
	}
}
