/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"log"
)

func Echo() {
	log.Printf("========== %s ==========\n", APPName)
	log.Printf("%s start serve [%s]\n", APPName, Root)
	log.Printf("%s deny dirs %v\n", APPName, DenyDirs)
	log.Printf("%s deny files %v\n", APPName, DenyFiles)
	log.Printf("%s deny ips %v\n", APPName, DenyIPs)
	log.Printf("%s start at [%s:%d]\n", APPName, Host, Port)
}
