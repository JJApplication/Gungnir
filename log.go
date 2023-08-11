/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"fmt"
	"log"
)

func InitLog() {
	log.SetPrefix(fmt.Sprintf("[%s] ", APPName))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
