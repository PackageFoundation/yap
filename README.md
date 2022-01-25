# Yap

![yap-logo](https://raw.githubusercontent.com/M0Rf30/yap/main/images/yap.png)

Yap allows building packages for multiple linux distributions with a
consistent package spec format. Currently `deb`, `rpm` and `pacman` packages
are available for several linux distributions.

Builds are done on Docker
containers without needing to setup any virtual machines or install any
software other than Docker.

All packages are built using a simple format that
is similar to [PKGBUILD](https://wiki.archlinux.org/index.php/PKGBUILD) from
Arch Linux.

Each distribution is different and will still require different
build instructions, but a consistent build process and format can be used for
all builds.

Docker only supports 64-bit containers, Yap can't be used to
build packages 32-bit packages.

## Initialize

It is recommended to build the Docker images locally instead of pulling each
image from the Docker Hub. A script is located in the docker directory to
assist with this. Always run the `clean.sh` script to clear any existing yap
images. Building the images can take several hours.

```sh
cd ~/go/src/github.com/packagefoundation/yap/docker
sh clean.sh
sh build.sh
```

## Format

```sh
key="example string"
key=`example "quoted" string`
key=("list with one element")
key=(
    "list with"
    "multiple elements"
)
key="example ${variable} string"
key:ubuntu="this will apply only to Ubuntu  builds"
```

## Builtin Variables

| key         | value                                                             |
| ----------- | ----------------------------------------------------------------- |
| `${srcdir}` | `Source` directory where all sources are downloaded and extracted |
| `${pkgdir}` | `Package` directory for the root of the package                   |

## Spec file - the PKGBUILD

| key                | type     | value                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `targets`          | `list`   | List of build targets only used for projects. Prefix a `!` to ignore target.                                                                                                                                                                                                                                                                                                                                                                               |
| `pkgname`          | `string` | Package name                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `pkgver`           | `string` | Package version                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `pkgrel`           | `string` | Package release number                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `pkgdesc`          | `string` | Short package description                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `pkgdesclong`      | `list`   | List of lines for package description                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `maintainer`       | `string` | Package maintainer                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `arch`             | `string` | Package architecture, can be `all` or `amd64`                                                                                                                                                                                                                                                                                                                                                                                                              |
| `license`          | `list`   | List of licenses for packaged software                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `section`          | `string` | Section for package. Built in sections available: `admin` `localization` `mail` `comm` `math` `database` `misc` `debug` `net` `news` `devel` `doc` `editors` `electronics` `embedded` `fonts` `games` `science` `shells` `sound` `graphics` `text` `httpd` `vcs` `interpreters` `video` `web` `kernel` `x11` `libdevel` `libs` |
| `priority`         | `string` | Package priority, only used for Debian packages                                                                                                                                                                                                                                                                                                                                                                                                            |
| `url`              | `string` | Package url                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `depends`          | `list`   | List of package dependencies                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `optdepends`       | `list`   | List of package optional dependencies                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `makedepends`      | `list`   | List of package build dependencies                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `provides`         | `list`   | List of packages provided                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `conflicts`        | `list`   | List of packages conflicts                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `sources`          | `list`   | List of packages sources. Sources can be url or paths that are relative to the PKGBUILD                                                                                                                                                                                                                                                                                                                                                                    |
| `debconf_config`   | `string` | File used as debconf config, only used for Debian packages                                                                                                                                                                                                                                                                                                                                                                                                 |
| `debconf_template` | `string` | File used as debconf template, only used for Debian packages                                                                                                                                                                                                                                                                                                                                                                                               |
| `hashsums`         | `list`   | List of `sha256`/`sha512` hex hashes for sources, hash type is determined by the length of the hash. Use `SKIP` to ignore hash check                                                                                                                                                                                                                                                                                                          |
| `backup`           | `list`   | List of config files that shouldn't be overwritten on upgrades                                                                                                                                                                                                                                                                                                                                                                                             |
| `build`            | `func`   | Function to build the source, starts in srcdir                                                                                                                                                                                                                                                                                                                                                                                                             |
| `package`          | `func`   | Function to package the source into the pkgdir, starts in srcdir                                                                                                                                                                                                                                                                                                                                                                                           |
| `preinst`          | `func`   | Function to run before installing                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `postinst`         | `func`   | Function to run after installing                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `prerm`            | `func`   | Function to run before removing                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `postrm`           | `func`   | Function to run after removing                                                                                                                                                                                                                                                                                                                                                                                                                             |

### Build targets

| target           | value                    |
| ---------------- | ------------------------ |
| `arch`           | all Arch Linux releases   |
| `astra`          | all Astra Linux releases  |
| `amazon`         | all Amazon Linux releases |
| `centos`         | all CentOS releases      |
| `debian`         | all Debian releases      |
| `fedora`         | all Fedora releases      |
| `oracle`         | all Oracle Linux releases |
| `ubuntu`         | all Ubuntu releases      |
| `amazon-1`       | Amazon Linux 1            |
| `amazon-2`       | Amazon Linux 2            |
| `centos-8`       | CentOS 8                 |
| `debian-jessie`  | Debian Jessie            |
| `debian-stretch` | Debian Stretch           |
| `debian-buster`  | Debian Buster            |
| `fedora-32`      | Fedora 32                |
| `fedora-33`      | Fedora 33                |
| `oracle-8`       | Oracle Linux 8            |
| `rocky-8`        | Rocky Linux 8             |
| `ubuntu-bionic`  | Ubuntu Bionic            |
| `ubuntu-focal`   | Ubuntu Focal             |

### Directives

Directives are used to specify variables that only apply to a limited set of
build targets.

All variables can use directives including user defined
variables.

To use directives include the directive after a
variable separated by a colon such as
`pkgdesc:ubuntu="This description will only apply to Ubuntu packages"`.

The directives above are sorted from lowest to the highest priority.

| directive        | value                    |
| ---------------- | ------------------------ |
| `apt`            | all deb packages         |
| `pacman`         | all pkg packages         |
| `yum`            | all rpm packages         |
| `arch`           | all Arch Linux releases   |
| `amazon`         | all Amazon Linux releases |
| `centos`         | all CentOS releases      |
| `debian`         | all Debian releases      |
| `fedora`         | all Fedora releases      |
| `oracle`         | all Oracle Linux releases |
| `ubuntu`         | all Ubuntu releases      |
| `amazon-1`       | Amazon Linux 1            |
| `amazon-2`       | Amazon Linux 2            |
| `centos-8`       | CentOS 8                 |
| `debian-jessie`  | Debian Jessie            |
| `debian-stretch` | Debian Stretch           |
| `debian-buster`  | Debian Buster            |
| `fedora-32`      | Fedora 32                |
| `fedora-33`      | Fedora 33                |
| `oracle-8`       | Oracle Linux 8            |
| `rocky-8`        | Rocky Linux 8             |
| `ubuntu-bionic`  | Ubuntu Bionic            |
| `ubuntu-focal`   | Ubuntu Focal             |

## Examples

Please have a look under the `examples` folder.

You'll find:

* [the project definition](examples/yap.json)
* [the spec file](examples/yap/PKGBUILD)

## License

See [LICENSE](LICENSE) file for details.

## Credits

[Zachary Huff](https://github.com/zachhuff386), for his work on
Pacur, on which Yap is based on.
