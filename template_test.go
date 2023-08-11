/*
   Create: 2023/8/9
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"html/template"
	"os"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	t.Log("start")
	tm, err := template.ParseFiles(Template)
	if err != nil {
		t.Error(err)
	}
	tm.Execute(os.Stdout, DataModel{DirList: []Entry{
		{
			Name:    "test",
			Size:    "",
			ModTime: "",
			IsDir:   false,
		},
	}})
}
