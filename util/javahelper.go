package util

import (
	"github.com/fatih/color"
	"os/exec"

	"strings"
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
	color.Green("ForgeWrapper is now lauching the laucnher")

	exec.Command("java", "-jar", getMcDir() + "/launcher.jar").Run()
}

func RunVersionCheck() bool {
	out, _ := exec.Command("java", "-version").CombinedOutput()

	n := string(out)

	if strings.Contains(n, "1.8") {
		return true
	} else {
		return false
	}
}
