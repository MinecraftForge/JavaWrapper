/*
 * Minecraft Forge
 * Copyright (c) 2016.
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation version 2.1
 * of the License.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
 */
package util

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/user"
	"runtime"
)

func GetThisPlatform() string {
	return runtime.GOOS
}

func GetArch() string {
	return runtime.GOARCH
}

func IsValidPlatFrom() bool {
	if GetThisPlatform() == "windows" || GetThisPlatform() == "darwin" || GetThisPlatform() == "linux" {
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
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		color.Red(runDir + ", Not found now creating")
		os.MkdirAll(runDir, os.ModePerm)
		color.Green(runDir + ", Has been created")
	}
}

//TODO boolean
func CheckForLauncher() {
	checkForRuntime()
	jar := getMcDir() + "/launcher.jar"
	if _, err := os.Stat(jar); os.IsNotExist(err) {
		color.Red(jar + ", Not found now downloading.")
		DownloadFromUrl("http://launcher.mojang.com/mc-staging/launcher/jar/3613e45482b58d3b5214911365d13afe3e24aa33/launcher.jar.lzma", getMcDir())
		color.Green(jar + ", Has been dowbloaded.")
		color.Green("decompressing launcher.jar.lzma")
		//DecompressLzma(getMcDir()+"/launcher.jar.lzma", getMcDir()+"/launcher.jar")
		//DecompFile(getMcDir() +"/launcher.jar.lzma")
	}
}

func RuntimeDownloader() {
	path := getMcDir() + "/runtime"
	darwin := "http://launcher.mojang.com/jre/osx-64/1.8.0_74/241139aa590e2aa139c0f0ede1dc98fdce3e3776/jre-osx-64-1.8.0_74.lzma"
	win_386 := "http://launcher.mojang.com/jre/win-32/1.8.0_51/9e6a4608c1116ee064d5ec4cabb9410bc4677f3c/jre-win-32-1.8.0_51.lzma"
	win_amd64 := "http://launcher.mojang.com/jre/win-64/1.8.0_51/3cb2e56d3f00a8a9fe1ca7e0e74380fdf7556cb0/jre-win-64-1.8.0_51.lzma"
	checkForRuntime()
	switch GetThisPlatform() {
	case "windows":
		if GetArch() == "amd64" {
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
