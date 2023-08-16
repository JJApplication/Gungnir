/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

var (
	Root       string
	Host       string
	Port       int
	Pool       string
	DenyDirs   []string          // 禁止访问的目录
	DenyFiles  []string          // 禁止访问的文件
	DenyIPs    []string          // 禁止访问的IP
	DenyAgents []string          // 禁止访问的user-agent
	Headers    map[string]string // 自定义headers
	SyncCount  int
)
