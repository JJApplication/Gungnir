/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handleFs(root string) *http.ServeMux {
	f := fileHandler{rootPath: root, root: http.Dir(root)}
	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
	})

	// method POST
	s.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			authFunc(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	s.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			uploadFunc(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	})

	return s
}

type authBody struct {
	Token string `json:"token"`
}

// auth
// 基于token的认证
func authFunc(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("check auth failed: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var token authBody
	if err = json.Unmarshal(data, &token); err != nil {
		log.Printf("parse token failed: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token.Token != "" && token.Token == Token {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	return
}

// file upload
func uploadFunc(w http.ResponseWriter, r *http.Request) {
	if !EnableUpload {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	file, header, err := r.FormFile("file")
	log.Printf("upload file [%s], error: %v\n", header.Filename, err)
	if _, ok := safe(header.Filename); !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if err == nil {
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			log.Printf("save file error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fileName, _ := safe(header.Filename)
		err = os.WriteFile(filepath.Join(Root, fileName), data, 0644)
		if err != nil {
			log.Printf("save file error: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusForbidden)
	return
}

func safe(f string) (string, bool) {
	if strings.Contains(f, "./") {
		return "", false
	}
	if strings.Contains(f, "../") {
		return "", false
	}
	f = strings.ReplaceAll(f, "/", "-")
	return f, true
}
