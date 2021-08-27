package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

type TemplateManager interface {
	ParseTemplates() (map[string]*template.Template, error)
	ExecuteTemplateToString(tmpl *template.Template, templateData interface{}) (string, error)
}

type StandardTemplatesManager struct {
	path string
}

func NewStandardTemplatesManager(path string) *StandardTemplatesManager {
	return &StandardTemplatesManager{path: path}
}

type ConfirmEmailData struct {
	Host        string
	ConfirmLink string
}

func (t *StandardTemplatesManager) ParseTemplates() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", t.path))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}

func (t *StandardTemplatesManager) ExecuteTemplateToString(tmpl *template.Template, templateData interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, templateData); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
