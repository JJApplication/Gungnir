/*
   Create: 2023/8/3
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package main

import "gungnir"

func main() {
	gungnir.InitLog()
	gungnir.InitFlag()
	gungnir.InitConfig()
	gungnir.Echo()
	gungnir.Run(gungnir.Root)
}
