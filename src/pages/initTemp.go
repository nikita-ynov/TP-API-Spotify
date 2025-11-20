package pages

import (
	"html/template"
	"os"
	"path/filepath"
)

var Temp *template.Template

func add1(x int) int {
	return x + 1
}

func Init() {
	Temp = template.New("").Funcs(template.FuncMap{
		"add1": add1,
	})

	err := filepath.Walk("pages", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, e := Temp.ParseFiles(path)
			return e
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
