/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

const (
	APPName          = "Gungnir"
	DefaultRoot      = "/"
	DefaultHost      = "127.0.0.1"
	DefaultPort      = 7086
	DefaultConfig    = "gungnir.json"
	DefaultPool      = "counts.json"
	DefaultSyncCount = 60 * 60
)

var (
	DefaultDeny = []string{".git,", ".cache"}
)
