package util

import (
	//"bufio"
	//"fmt"
	"github.com/fatih/color"
	"os/exec"

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

func RunVersionCheck() {
	out, err := exec.Command("java", "-version").CombinedOutput()

	if err != nil {
	}

	//DownloadFromUrl()

	color.Green("The installed java version is\n%s", out)
}
