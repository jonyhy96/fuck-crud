// Package generator go files.
package generator

import (
	"io"

	"github.com/jonyhy96/fuck-crud/pkg/transform"
	templateutil "github.com/jonyhy96/fuck-crud/pkg/util/template"
)

// Generator generate.
type Generator interface {
	Generate(io.Writer) error
}

// Generate generate.
func Generate(source transform.Config, tpl string, out io.Writer) error {
	t := templateutil.GetTemplate()
	t, err := t.Parse(tpl)
	if err != nil {
		return err
	}
	return t.Execute(out, source)
}
