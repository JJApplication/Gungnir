/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"net/http"
	"strings"
)

// ip黑名单过滤

func filterIPs(r *http.Request) bool {
	addr := r.RemoteAddr
	for _, ip := range DenyIPs {
		if strings.Contains(addr, ip) {
			return true
		}
	}
	return false
}
