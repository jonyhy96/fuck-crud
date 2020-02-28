package template

import (
	"html/template"
	"strings"

	"github.com/iancoleman/strcase"
)

// GetTemplate gets a template.
func GetTemplate() *template.Template {
	return template.New("t").Funcs(template.FuncMap{
		"upperFirst": func(origin string) string {
			return strings.Title(strings.ToLower(origin))
		},
		"lowerCamel": func(origin string) string {
			return strcase.ToLowerCamel(origin)
		},
		"camel": func(origin string) string {
			return strcase.ToCamel(origin)
		},
	})
}
