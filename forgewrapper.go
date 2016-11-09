package main

import (
	"fmt"
	"github.com/illyohs/forgewrapper/util"
	"github.com/fatih/color"
	"os"
)

func main() {
	if !util.IsValidPlatFrom() {
		color.Red("The forge wrapper doesn't support %s", util.GetThisPlatform())
		os.Exit(1)
	}
	fmt.Println(util.GetArch())
	fmt.Println(util.IsValidPlatFrom())
	fmt.Println(util.IsValidArch())
	util.IsJavaInstalled()
	util.RunVersionCheck()
	util.RuntimeDownloader()
}
