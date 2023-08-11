/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import "net/http"

func modifyHeaders(w http.ResponseWriter, headers map[string]string) {
	for key, val := range headers {
		w.Header().Set(key, val)
	}
}

func customHeaders(w http.ResponseWriter) {
	modifyHeaders(w, Headers)
}
