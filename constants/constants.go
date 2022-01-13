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
		"arch",
		"astra",
		"amazon-1",
		"amazon-2",
		"fedora-35",
		"centos-8",
		"debian-jessie",
		"debian-stretch",
		"debian-buster",
		"oracle-8",
		"rocky-8",
		"ubuntu-bionic",
		"ubuntu-focal",
	}
	ReleasesMatch = map[string]string{
		"arch":           "",
		"astra":          "astra",
		"amazo-1":        ".amzn1.",
		"amazon-2":       ".amzn2.",
		"fedora-35":      ".fc35.",
		"centos-8":       ".el8.",
		"debian-jessie":  "jessie",
		"debian-stretch": "stretch",
		"debian-buster":  "buster",
		"oracle-8":       ".ol8.",
		"rocky-8":        ".el8.",
		"ubuntu-bionic":  "bionic",
		"ubuntu-focal":   "focal",
	}
	DistroPack = map[string]string{
		"arch":   "pacman",
		"astra":  "debian",
		"amazon": "redhat",
		"fedora": "redhat",
		"centos": "redhat",
		"debian": "debian",
		"oracle": "redhat",
		"rocky":  "redhat",
		"ubuntu": "debian",
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
