package conf

import (
	"html/template"
	"log"
	"path/filepath"
	"strings"
)

const (
	layoutPattern   = "./web/views/index.layout.html"
	templatePattern = "./web/views/**/*.tmpl.html"
	partialPattern  = "./web/views/**/_*.tmpl.html"
	viewsRootDir    = "./web/views"
)

var cache map[string]*template.Template

func BuildTemplateCache(funcs template.FuncMap) (map[string]*template.Template, error) {
	cache = map[string]*template.Template{}

	baseTemplate, err := parseLayouts(funcs)
	if err != nil {
		return nil, err
	}

	pages, err := filepath.Glob(templatePattern)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name, err := nameTemplate(page)
		if err != nil {
			return nil, err
		}

		tmpl, err := baseTemplate.Clone()
		if err != nil {
			return nil, err
		}
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	return cache, nil
}

func GetTemplates(funcs template.FuncMap) map[string]*template.Template {
	if GetEnv().AppEnv == Production {
		return cache
	}

	cache, err := BuildTemplateCache(funcs)
	if err != nil {
		log.Fatalf("could not build template cache, %s", err.Error())
	}

	return cache
}

func parseLayouts(funcs template.FuncMap) (*template.Template, error) {
	base := template.New("index.layout.html").Funcs(funcs)

	_, err := base.ParseGlob(layoutPattern)
	if err != nil {
		return nil, err
	}

	partials, err := filepath.Glob(partialPattern)
	if err != nil {
		return nil, err
	}

	if len(partials) > 0 {
		_, err = base.ParseGlob(partialPattern)
		if err != nil {
			return nil, err
		}
	}
	return base, nil
}

func nameTemplate(path string) (string, error) {
	relPath, err := filepath.Rel(viewsRootDir, path)
	if err != nil {
		return "", err
	}

	key := strings.TrimSuffix(relPath, ".tmpl.html")
	return key, err
}
