/**
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
	"os"
	"os/user"
	"runtime"

	"github.com/fatih/color"
)

func GetThisPlatform() string {
	return runtime.GOOS
}

func GetThisArch() string {
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
	if GetThisArch() == "amd64" || GetThisArch() == "386" {
		return true
	} else {
		return false
	}
}

func checkForMcdir() {
	if _, err := os.Stat(getMcDir()); os.IsNotExist(err) {
		color.Red(getMcDir() + ", Not found")
		color.Yellow("Now setting up runtime")
		os.MkdirAll(getMcDir(), os.ModePerm)
		color.Green(getMcDir() + ", Has been created")
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

//I'm going to need a gradle task for this for this one
func InstallerVersion() string {
	return "@VERSION@"
}

func getRuntimeJREDir() string {
	return string(getMcDir() + "/runtime/" + GetJREVersion())
}

func checkForRuntime() {
	// ver := GetJREVersion()
	runDir := getRuntimeJREDir()
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		color.Red(runDir + ", Not found")
		color.Yellow("Now setting up runtime")
		os.MkdirAll(runDir, os.ModePerm)
		runtimeDownloader()
		color.Green(runDir + ", Has been created")
	}
}
func CheckForLauncher() {
	jar := getMcDir() + "/launcher.jar"
	if _, err := os.Stat(jar); os.IsNotExist(err) {
		color.Red(jar + ", Not found now downloading.")
		DownloadFromUrl(GetLauncherUrl(), getMcDir())
		color.Green(jar + ", Has been downloaded.")
		color.Green("decompressing launcher.jar.lzma")
		DecompLauncher()
	}
}

func runtimeDownloader() {
	platform, arch, version, url := GetJreInfo()
	path := getMcDir() + "/runtime/" + version

	color.Green("Downloading Jre version %s for %s%s", version, platform, arch)
	DownloadFromUrl(url, path)
	DecompJRE(version)
}
