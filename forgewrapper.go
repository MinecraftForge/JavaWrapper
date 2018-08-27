//go:generate goversioninfo -icon=resources/icon.ico -manifest=javawrapper.exe.manifest

// Minecraft Forge
// Copyright (c) 2018.

// This library is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation version 2.1
// of the License.

// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.

// You should have received a copy of the GNU Lesser General Public
// License along with this library; if not, write to the Free Software
// Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
package main

import (
	"log"
	"os"

	"github.com/MinecraftForge/JavaWrapper/util"
)

func main() {

	log.Println("Platfrom: " + util.GetThisPlatform())
	log.Println("Arch: " + util.GetSysArch())
	if !util.IsValidPlatFrom() {
		log.Printf("The forge wrapper doesn't support %s", util.GetThisPlatform())
		os.Exit(1)
	}

	util.ModedLauncher()

}
