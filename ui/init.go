package ui

import (
	"html/template"
	"os"
)

const (
	templatesPath  = "./templates"
	layoutGlob     = "*.go.html"
	pagesGlob      = "pages/*.go.html"
	componentsGlob = "components/*.go.html"
)

var (
	tmpl *template.Template
)

func init() {
	fsys := os.DirFS(templatesPath)
	tmpl = template.Must(template.ParseFS(fsys, layoutGlob, pagesGlob, componentsGlob))
}
