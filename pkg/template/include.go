package template

import (
	"bytes"
	"io"
)

// Template defines the interface which expects the ExecuteTemplate function
// to be defined, so it can be used by the include function.
type Template interface {
	ExecuteTemplate(io.Writer, string, any) error
}

// Returns the function map defining the include template function. Which allows
// to call defined templates from within a template.
func IncludeFunc(tpl Template) map[string]any {
	return map[string]any{
		"include": func(name string, data any) (string, error) {
			buf := bytes.NewBufferString("")
			if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
				return "", err
			}
			return buf.String(), nil
		},
	}
}
