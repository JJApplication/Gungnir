/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

func InitConfig() {
	if ConfigGen {
		newConfig()
		return
	}
	cf := loadConfig()
	cf.Init()
}
