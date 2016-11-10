package util

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"strings"
)

func DownloadFromUrl(url string, path string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	color.Cyan("Downloading %s from %s", fileName, url)

	output, err := os.Create(path + "/" + fileName)
	if err != nil {
		color.Red("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		color.Red("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		color.Red("Error while downloading", url, "-", err)
		return
	}
	fmt.Println(n, "bytes downloaded.")
}
