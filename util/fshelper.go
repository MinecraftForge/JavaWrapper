package util

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

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func GetThisPlatform() string {
	return runtime.GOOS
}

func IsValidPlatFrom() bool {
	if GetThisPlatform() == "windows" || GetThisPlatform() == "darwin" || GetThisPlatform() == "linux" {
		return true
	} else {
		return false
	}
}

func getArchWindows() string {
	out, _ := exec.Command("wmic", "os", "get", "osarchitecture").CombinedOutput()

	if strings.Contains(string(out), "32-bit") {
		return "32"
	} else if strings.Contains(string(out), "64-bit") {
		return "64"
	}
	return ""
}

func getArchLinux() string {
	out, _ := exec.Command("uname", "-a").CombinedOutput()

	if strings.Contains(string(out), "i386") || strings.Contains(string(out), "i868") {
		return "32"
	} else if strings.Contains(string(out), "x86_64") {
		return "64"
	}
	return ""
}

func GetSysArch() string {

	switch runtime.GOOS {
	case "windows":
		return getArchWindows()
	case "linux":
		return getArchLinux()
	case "darwin":
		return "64"
	}
	return ""
}

func IsValidArch() bool {
	if GetSysArch() == "64" || GetSysArch() == "32" {
		return true
	} else {
		return false
	}
}

func checkForMcdir() {
	if _, err := os.Stat(getMcDir()); os.IsNotExist(err) {
		log.Fatalln(getMcDir() + ", Not found")
		log.Println("Creating Mc directory")
		os.MkdirAll(getMcDir(), os.ModePerm)
		log.Println(getMcDir() + ", Has been created")
	}
}

func getMcDir() string {
	usr, err := user.Current()

	home := usr.HomeDir

	if err != nil {
		log.Panicln("Unable to locate the user home directory")

	}
	if GetThisPlatform() == "windows" {
		return home + "/AppData/Roaming/.minecraft"
	} else if GetThisPlatform() == "darwin" {
		return home + "/Library/Application Support/minecraft"
	} else {
		return home + "/.minecraft"
	}
}

func getRuntimeJREDir() string {
	return string(getMcDir() + "/runtime/" + GetJREVersion())
}

func checkForRuntime() {
	// ver := GetJREVersion()
	runDir := getRuntimeJREDir()
	if _, err := os.Stat(runDir); os.IsNotExist(err) {
		log.Fatalln(runDir + ", Not found")
		log.Println("Now setting up runtime")
		os.MkdirAll(runDir, os.ModePerm)
		runtimeDownloader()
		log.Println(runDir + ", Has been created")
	}
}

func CheckForLauncher() {
	checkForMcdir()
	jar := getMcDir() + "/launcher.jar"
	if _, err := os.Stat(jar); os.IsNotExist(err) {
		log.Println(jar + ", Not found now downloading.")
		DownloadFromUrl(GetLauncherUrl(), getMcDir())
		log.Println(jar + ", Has been downloaded.")
		log.Println("decompressing launcher.jar.lzma")
		DecompLauncher()
	}
}

func runtimeDownloader() {
	platform, arch, version, url := GetJreInfo()
	path := getMcDir() + "/runtime/" + version

	log.Printf("Downloading Jre version %s for %s%s", version, platform, arch)
	DownloadFromUrl(url, path)
	DecompJRE(version)
}
