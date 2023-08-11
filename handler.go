/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"net/http"
)

func handleFs(root string) *http.ServeMux {
	f := fileHandler{rootPath: root, root: http.Dir(root)}
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
	})

	return s
}
