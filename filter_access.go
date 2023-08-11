/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import "strings"

// 文件黑名单防御

func filterDenyFiles(f string) bool {
	paths := strings.Split(f, "/")
	if len(paths) > 0 {
		lastUri := paths[len(paths)-1]
		for _, d := range DenyFiles {
			if d == lastUri {
				return true
			}
		}
	}

	return false
}

func filterDenyDirs(d string) bool {
	for _, dir := range DenyDirs {
		if dir == d {
			return true
		}
	}
	return false
}
