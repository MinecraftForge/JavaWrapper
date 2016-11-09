package util

import (
	"fmt"
	"os/user"
	"runtime"
	"os"
	"github.com/fatih/color"
)

func GetThisPlatform() string {
	return runtime.GOOS
}

func GetArch() string {
	return runtime.GOARCH
}

func IsValidPlatFrom() bool {
	if GetThisPlatform() == "windows" || GetThisPlatform() == "darwin" {
		return true
	} else {
		return false
	}
}

func IsValidArch() bool {
	if GetArch() == "amd64" || GetArch() == "386" {
		return true
	} else {
		return false
	}
}

func getMcDir() string {
	usr, err := user.Current()

	home := usr.HomeDir

	if err != nil {
		fmt.Println("Unable to locate the user home directory")

	}
	if GetThisPlatform() == "windows" {
		return home + "/AppData/Roaming/.minecraft"
	} else if GetThisPlatform() == "darwin" {
		return home + "/Library/Application Support/minecraft"
	} else {
		return home + "/.minecraft"
	}
}

func checkForRuntime() {
	runDir := getMcDir() + "/runtime"
	if _ , err := os.Stat(runDir); os.IsNotExist(err){
		color.Red(runDir + ", Not found now creating")
		os.MkdirAll(runDir, os.ModePerm)
		color.Green(runDir + ", Has been created")
	}
}

func CheckForLauncher()  {
	checkForRuntime()
	jar := getMcDir() + "/launcher.jar"
	if _ , err := os.Stat(jar); os.IsNotExist(err){
		color.Red(jar + ", Not found now downloading.")
		DownloadFromUrl("http://launcher.mojang.com/mc-staging/launcher/jar/3613e45482b58d3b5214911365d13afe3e24aa33/launcher.jar.lzma", getMcDir())
		color.Green(jar + ", Has been dowbloaded.")
	}
}

func RuntimeDownloader()  {
	path := getMcDir() + "/runtime"
	darwin := "http://launcher.mojang.com/jre/osx-64/1.8.0_74/241139aa590e2aa139c0f0ede1dc98fdce3e3776/jre-osx-64-1.8.0_74.lzma"
	win_386 := "http://launcher.mojang.com/jre/win-32/1.8.0_51/9e6a4608c1116ee064d5ec4cabb9410bc4677f3c/jre-win-32-1.8.0_51.lzma"
	win_amd64 := "http://launcher.mojang.com/jre/win-32/1.8.0_51/9e6a4608c1116ee064d5ec4cabb9410bc4677f3c/jre-win-32-1.8.0_51.lzma"
	checkForRuntime()
	switch GetThisPlatform() {
	case "windows":
		if GetArch() == "amd64"{
			DownloadFromUrl(win_amd64, path)
		} else if GetArch() == "386" {
			DownloadFromUrl(win_386, path)
		} else {
			color.Red("UNABLE TO DETERMIN ARCHITECTURE NOW SHUTTING DOWN")
			os.Exit(2)
		}
	case "darwin":
		DownloadFromUrl(darwin, path)
	}
}