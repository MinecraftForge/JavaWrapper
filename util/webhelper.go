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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type LauncherJson struct {
	Java struct {
		Lzma struct {
			Sha1 string `json:"sha1"`
			URL  string `json:"url"`
		} `json:"lzma"`
		Sha1 string `json:"sha1"`
	} `json:"java"`
	Osx struct {
		Arch64 struct {
			Jdk struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jdk"`
			Jre struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jre"`
		} `json:"64"`
		Apphash      string `json:"apphash"`
		Applink      string `json:"applink"`
		Downloadhash string `json:"downloadhash"`
	} `json:"osx"`
	Windows struct {
		Arch32 struct {
			Jdk struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jdk"`
			Jre struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jre"`
		} `json:"32"`
		Arch64 struct {
			Jdk struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jdk"`
			Jre struct {
				Sha1    string `json:"sha1"`
				URL     string `json:"url"`
				Version string `json:"version"`
			} `json:"jre"`
		} `json:"64"`
		Apphash string `json:"apphash"`
		Applink string `json:"applink"`
	} `json:"windows"`
}

func GetJreInfo() (string, string, string, string) {
	ljString := StringFromWebJson("http://launchermeta.mojang.com/mc-staging/launcher.json")
	buf := bytes.NewBufferString(ljString)
	var ljObj LauncherJson

	err := json.NewDecoder(buf).Decode(&ljObj)
	if err != nil {
		panic(err)
	}

	var platform string
	var arch string
	var version string
	var url string

	switch GetThisPlatform() {
	case "windows":
		platform = "windows"
		if GetThisArch() == "amd64" {
			arch = "64"
			version = ljObj.Windows.Arch64.Jre.Version
			url = ljObj.Windows.Arch64.Jre.URL
		} else if GetThisArch() == "386" {
			arch = "32"
			version = ljObj.Windows.Arch32.Jre.Version
			url = ljObj.Windows.Arch32.Jre.URL
		} else {
			color.Red("UNABLE TO DETERMINE ARCHITECTURE NOW SHUTTING DOWN")
			os.Exit(2)
		}
	case "darwin":
		platform = "osx"
		arch = "64"
		version = ljObj.Osx.Arch64.Jre.Version
		url = ljObj.Osx.Arch64.Jre.URL
	case "linux":
		color.Red("Sorry Mojang has not build a JRE for this platfrom please update to your java " +
			"go to http://openjdk.java.net/install/ or " +
			"http://www.oracle.com/technetwork/java/javase/downloads/index.html to download the latest java.")
		os.Exit(3)
	}

	return platform, arch, version, url

}

func GetJREVersion() string {
	_, _, version, _ := GetJreInfo()
	return version
}

func GetLauncherUrl() string {
	ljString := StringFromWebJson("http://launchermeta.mojang.com/mc-staging/launcher.json")
	buf := bytes.NewBufferString(ljString)
	var ljObj LauncherJson

	err := json.NewDecoder(buf).Decode(&ljObj)
	if err != nil {
		fmt.Println(err)
	}

	return ljObj.Java.Lzma.URL
}
