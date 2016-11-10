package main

import (
	"fmt"
	"github.com/illyohs/forgewrapper/util"
	"os"
	"github.com/fatih/color"
)

func main() {
	if !util.IsValidPlatFrom() {
		color.Red("The forge wrapper doesn't support %s", util.GetThisPlatform())
		os.Exit(1)
	}
	util.CheckForLauncher()
	util.LaunchWithSysJava()
	fmt.Println(util.GetArch())
	fmt.Println(util.IsValidPlatFrom())
	fmt.Println(util.IsValidArch())

}
