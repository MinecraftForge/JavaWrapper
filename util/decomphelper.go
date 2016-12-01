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
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/ulikunitz/xz/lzma"
)

func DecompLauncher() {
	DecompLzma(getMcDir()+"/launcher.jar.lzma", getMcDir()+"/launcher.jar")
}

func DecompJRE(version string) {
	var targetName string

	//TODO get from json
	if GetThisPlatform() == "darwin" {
		targetName = "jre-osx-64-" + version + ".lzma"
	} else if GetThisPlatform() == "windows" {
		if GetThisArch() == "amd64" {
			targetName = "jre-win-64-" + version + ".lzma"
		} else if GetThisArch() == "386" {
			targetName = "jre-win-32-" + version + ".lzma"
		}
	}
	DecompLzma(getRuntimeJREDir()+"/"+targetName, getRuntimeJREDir()+"/jre.zip")
	// println("test")
	unzip(getRuntimeJREDir()+"/jre.zip", getRuntimeJREDir())

}

func DecompLzma(archive, target string) {
	f, err := os.Open(archive)

	if err != nil {
		fmt.Println(err)
	}

	r, err := lzma.NewReader(bufio.NewReader(f))

	if err != nil {
		fmt.Println(err)
	}

	output, err := os.Create(target)

	if err != nil {
		color.Red("error %s", err)
	}

	cop, err := io.Copy(output, r)
	fmt.Println(cop, "creaded")
	os.Remove(archive)

}

//Taken from https://gist.github.com/svett/424e6784facc0ba907ae#file-extract-go
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)

	if err != nil {
		return err
	}

	defer  reader.Close()

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fileReader.Close()
			return err
		}

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			fileReader.Close()
			targetFile.Close()
			return err
		}

		fileReader.Close()
		targetFile.Close()
	}

	return os.Remove(archive)
}
