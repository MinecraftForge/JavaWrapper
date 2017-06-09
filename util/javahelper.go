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

/*
 * Launches with system jre
 * target: the target jar
 * args: the jar args
 *
 */
func GenericSysJavaLauncher(target string) ([]byte, error) {
	// color.Yellow("Forge wrapper is now launching %s", target)
	color.Yellow("Now running from System JRE")
	out, err := exec.Command("java", "-jar", target).CombinedOutput()

	if err != nil {
		println(err)
		// panic(err)
	}

	return out, err

}

/*
 * Launches the mojang jre
 * target: the target jar
 * args: the jar args
 *
 */
func GenericMojangJavaLauncher(target string) ([]byte, error) {

	checkForRuntime()

	darwinJRE := getRuntimeJREDir() + "/bin/java"
	winJRE := getRuntimeJREDir() + "/bin/java.exe"
	color.Yellow("Now running from Mojang JRE")

	switch GetThisPlatform() {
	case "windows":
		out, err := exec.Command(winJRE, "-jar", target).CombinedOutput()
		return out, err
	case "darwin":
		out, err := exec.Command(darwinJRE, "-jar", target).CombinedOutput()
		return out, err
	case "linux":
		color.Red("Sorry Mojang does not currently distribute a JRE for linux please download a JRE from" +
			"http://openjdk.java.net/install/ or " +
			"http://www.oracle.com/technetwork/java/javase/downloads/index.html")
		os.Exit(3)
	}
	return nil, nil
}

/*
 * Check for java 8
 */
func IsJavaVersionValid() bool {

	//Anoying check due to mac being tempermental
	if GetThisPlatform() == "darwin" {
		return macHasJ8()
	} else {
		out, _ := exec.Command("java", "-version").CombinedOutput()

		if strings.Contains(string(out), "1.8") {
			return true
		} else {
			return false
		}
	}
}

func Wrapper(jar string) ([]byte, error) {

	if IsJavaVersionValid() {
		out, err := GenericSysJavaLauncher(jar)
		return out, err
	} else {
		out, err := GenericMojangJavaLauncher(jar)
		return out, err
	}
}

/*
 * Launches either the wrapper or an appended jar
 */
func ModedLauncher() {
	binName, _ := os.Executable()

	CheckForLauncher()
	_, err := Wrapper(binName)

	if err != nil {
		color.Yellow("Now launching in launcher mode")
		CheckForLauncher()
		nout, nerr := Wrapper(getMcDir() + "/launcher.jar")
		if nerr != nil {

			print(nerr)
		}
		print(string(nout))

	} else {
		color.Yellow("Launching in Generic Jar Wrapper mode")
	}

}
