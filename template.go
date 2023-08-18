/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type DataModel struct {
	EnableUpload bool
	DirList      []Entry
}

type Entry struct {
	Name    string
	Size    string
	ModTime string
	IsDir   bool
}

type FileInfo struct {
	EnableUpload bool
	FileName     string
	FileUrl      string
	Size         string
	ModTime      string
	Ext          string
	MD5          string
	Counts       int
}

const (
	TemplateIndex    = "gungnir_index.tmpl"
	TemplateFile     = "gungnir_file.tmpl"
	TemplateDeny     = "gungnir_403.tmpl"
	TemplateNotFound = "gungnir_404.tmpl"
	TemplateError    = "gungnir_500.tmpl"
)

func (d *DataModel) add(entry Entry) {
	d.DirList = append(d.DirList, entry)
}

func serveTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Printf("serve template [%s] error: %s\n", tmpl, err)
	}
}

func serveIndex(w http.ResponseWriter, data DataModel) {
	serveTemplate(w, TemplateIndex, data)
}

func serveFilePage(w http.ResponseWriter, stat fs.FileInfo, uri string) {
	data := FileInfo{
		EnableUpload: EnableUpload,
		FileName:     stat.Name(),
		FileUrl:      uri,
		Size:         calcFileSize(stat.Size()),
		ModTime:      getModTime(stat.ModTime()),
		Ext:          getExt(stat.Name()),
		Counts:       getPool(uri).Count,
	}
	serveTemplate(w, TemplateFile, data)
}

func serveDeny(w http.ResponseWriter, data any) {
	serveTemplate(w, TemplateDeny, data)
}

func serveNotFound(w http.ResponseWriter) {
	serveTemplate(w, TemplateNotFound, nil)
}

func serveError(w http.ResponseWriter, data any) {
	serveTemplate(w, TemplateError, data)
}
