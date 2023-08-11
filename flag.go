/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"flag"
)

var (
	Config    string
	ConfigGen bool
)

func InitFlag() {
	flag.StringVar(&Root, "root", DefaultRoot, "serve root dir")
	flag.StringVar(&Host, "host", DefaultHost, "serve listen host")
	flag.StringVar(&Config, "conf", DefaultConfig, "serve config path")
	flag.IntVar(&Port, "port", DefaultPort, "serve listen port")
	flag.BoolVar(&ConfigGen, "new", false, "new server config")

	flag.Parse()
}
