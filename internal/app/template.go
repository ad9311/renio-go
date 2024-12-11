package app

import (
	"fmt"
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

var (
	cache     map[string]*template.Template
	tmplFuncs template.FuncMap
)

func BuildTemplateCache() (map[string]*template.Template, error) {
	cache = map[string]*template.Template{}
	tmplFuncs = generateTmplFunctions()

	baseTemplate, err := parseLayouts(tmplFuncs)
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

func GetTemplates() map[string]*template.Template {
	if GetEnv().AppEnv == Production {
		return cache
	}

	cache, err := BuildTemplateCache()
	if err != nil {
		log.Fatalf("could not build template cache, %s", err.Error())
	}

	return cache
}

// --- Helpers --- //

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

func generateTmplFunctions() template.FuncMap {
	funcs := template.FuncMap{}

	funcs["formatCurrency"] = func(value float32) string {
		return fmt.Sprintf("$%.2f", value)
	}

	funcs["entryClassColor"] = func(entryClassUID string) string {
		switch entryClassUID {
		case "banking":
			return "#FF5733"
		case "clothing":
			return "#2ECC71"
		case "entertainment":
			return "#2980B9"
		case "extra":
			return "#FF33A1"
		case "foodDelivery":
			return "#8E44AD"
		case "groceries":
			return "#27AE60"
		case "home":
			return "#F39C12"
		case "insurance":
			return "#E74C3C"
		case "onlineShopping":
			return "#3498DB"
		case "other":
			return "#9B59B6"
		case "restaurants":
			return "#1ABC9C"
		case "savings":
			return "#D35400"
		case "subscriptions":
			return "#C0392B"
		case "transportation":
			return "#16A085"
		case "utilities":
			return "#8E44AD"
		case "wages":
			return "#2980B9"
		default:
			return "#000000"
		}
	}

	return funcs
}
