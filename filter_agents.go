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

// 基于user-agent的过滤

func filterAgents(r *http.Request) bool {
	agent := r.UserAgent()
	for _, a := range DenyAgents {
		if strings.Contains(agent, a) {
			return true
		}
	}
	return false
}
