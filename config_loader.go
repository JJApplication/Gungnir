/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import "os"

func InitConfig() {
	if ConfigGen {
		newConfig()
		os.Exit(0)
		return
	}
	cf := loadConfig()
	cf.Init()
}
