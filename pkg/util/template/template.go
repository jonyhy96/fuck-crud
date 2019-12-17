package template

import (
	"html/template"
	"strings"
)

// GetTemplate gets a template.
func GetTemplate() *template.Template {
	return template.New("t").Funcs(template.FuncMap{
		"upperFirst": func(origin string) string {
			return strings.Title(strings.ToLower(origin))
		},
	})
}
