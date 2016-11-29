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
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func IsJavaInstalled() bool {
	jv, err := exec.LookPath("java")

	if err != nil {
		color.Red("Java has not been found!")
		return false
	} else {
		color.Green("Java has been found at: %s", jv)
		return true
	}
}

func LaunchWithSysJava() {
	color.Green("ForgeWrapper is now lauching the laucnher with system JRE")

	out, err := exec.Command("java", "-jar", getMcDir()+"/launcher.jar").CombinedOutput()

	if err != nil {
		println(err)
	}

	println(out)
}

func LaunchWithMojangJava() {
	darwinJRE := GetRuntimeJRE() + "/runtime/bin/java"
	winJRE := GetRuntimeJRE() + "/runtime/bin/java.exe"
	color.Green("Now running the Launcher from Mojang JRE")

	if GetThisPlatform() == "windows" {
		exec.Command(winJRE, "-jar", getMcDir()+"/launcher.jar").Run()
	} else if GetThisPlatform() == "darwin" {
		exec.Command(darwinJRE, "-jar", getMcDir()+"/launcher.jar").Run()
	} else if GetThisPlatform() == "linux" {
		color.Red("Sorry Mojang has not built a JRE for linux please download from go to " +
			"http://openjdk.java.net/install/ or " +
			"http://www.oracle.com/technetwork/java/javase/downloads/index.html to download the latest java 8.")
		os.Exit(3)
	}

}

func IsJavaVersionValid() bool {
	out, _ := exec.Command("java", "-version").CombinedOutput()

	if strings.Contains(string(out), "1.8") {
		return true
	} else {
		return false
	}
}

func JreLauncher() {
	if IsValidPlatFrom() {
		CheckForLauncher()
		if IsJavaVersionValid() {
			LaunchWithSysJava()
		} else {
			LaunchWithMojangJava()
		}
	}
}
