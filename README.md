# JavaWrapper

Simple java wrapper using Mojangs JREs

## Compiling
__Requirements__
  * Java 8
  * [Gradle](https://gradle.org/)
  * [GoLang](https://golang.org/) (while not needed it is __HIGHLY__ encouraged)

__Instructions for the default build__
  * Run the command `gradle build`
  * build outputs to `./build/out/`

__Instructions for packaging the default build__
  * Run the command `gradle pkg`
  * Build outputs to `./build/pkg/`

## Contributing
__Requirements__
  * [GoLang](https://golang.org/)

__Code Style__
  * License header under the package name of every go file
  * Run the command `go fmt path/to/go/file`
    * Golang will autoformat by using this command

__PR Guidelines__
  * Target the devel branch
  * One feature per PR
  * Sign off commits
