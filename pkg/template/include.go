package template

import (
	"bytes"
	"io"
)

type Template interface {
	ExecuteTemplate(io.Writer, string, any) error
}

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
