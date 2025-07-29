package main

import (
	"html/template"
	"path/filepath"

	"github.com/Infamous003/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpls, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}
		tmpls, err = tmpls.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		tmpls, err = tmpls.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = tmpls
	}
	return cache, nil
}
