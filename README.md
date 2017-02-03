# JavaWrapper

Simple java wrapper using Mojangs JREs

## Compileing
__Requirements__
  * Java 8
  * [Gradle](https://gradle.org/)
  * [GoLang](https://golang.org/) (while not needed it is __HIGHLY__ encouraged)

__Instructions for the default build__
  * Run the command `gradle build`
  * build outputs to `./build/out/`

__Instructions for packaging the default build__
  * Run the command `gradle pkgNormal`
  * Build outputs to `./build/pkg/`

__Instructions for packaging the forge installer__
  * set the environment variable `INSTALLER_VERSION`
    * For Windows `set INSTALLER_VERSION=VERSION`
      * e.g. `set INSTALLER_VERSION=VERSION=1.11-13.20.0.2226`
    * For Linux/Mac `export INSTALLER_VERSION=VERSION`
      * e.g `export INSTALLER_VERSION=1.11-13.20.0.2226`
  * Run the command `gradle installer`
  * Build outputs `./build/pkginstaller/`
  * NOTE: If the `INSTALLER_VERSION` is not set the installer defaults to
    forge version `1.11-13.19.1.2199`

##Contributing
__Requirements__
  * [GoLang](https://golang.org/)

__Code Style__
  * License header under the package name of every go file
  * Run the command `go fmt path/to/go/file`
    * Golang will autoformat by using this command

__PR Guidelines__
  * One feature per PR
  * Sign off commits
