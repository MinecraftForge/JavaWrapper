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
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	//"fmt"
	//"bufio"
	//"github.com/ulikunitz/xz/lzma"
)

func DecompLauncher() {
	DecompLzma("launcher.lzma", getMcDir()+"/launcher.jar")
}

func DecompJRE() {
	var targetName string

	//TODO get from json
	if GetThisPlatform() == "darwin" {
		targetName = "jre-osx-64-1.8.0_74.lzma"
	} else if GetThisPlatform() == "windows" {
		if GetArch() == "amd64" {
			targetName = "jre-win-64-1.8.0_51.lzma"
		} else if GetArch() == "386" {
			targetName = "jre-win-32-1.8.0_51.lzma"
		}
	}

	DecompLzma(targetName, getMcDir()+"/runtime/")

}

func DecompLzma(file, target string) {
	//f , err := os.Open(file)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//r , err := lzma.NewReader(bufio.NewReader(f))
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

}

//Taken from https://gist.github.com/svett/424e6784facc0ba907ae#file-extract-go
func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

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
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
