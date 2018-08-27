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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func DownloadFromUrl(url string, path string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	// color.Cyan("Downloading %s from %s", fileName, url)
	output, err := os.Create(path + "/" + fileName)
	if err != nil {
		log.Fatalln("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Fatalln("Error while downloading", url, "-", err)
		return
	}
	log.Println("Finsihed downloading ", fileName, n, "bytes downloaded.")
}

func StringFromWebJson(url string) string {

	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != err {
		log.Fatalln(err)
		// fmt.Println(err)
	}

	return string(bytes)
}
