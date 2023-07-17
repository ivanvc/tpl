package template

import (
	htmlTemplate "html/template"
	"strings"
	"testing"
	textTemplate "text/template"
)

var (
	validTemplate = `
{{- define "test" -}}
{{ . }}
{{- end -}}
{{ include "test" "abc" -}}
`
	invalidTemplate = `{{ include "test" "abc" }}`
)

func TestWithValidTemplate(t *testing.T) {
	tp := htmlTemplate.New("testValid")
	tpl, err := tp.Funcs(IncludeFunc(tp)).Parse(validTemplate)
	if err != nil {
		t.Errorf("Error parsing template %q", err)
	}
	var b strings.Builder
	if err := tpl.Execute(&b, nil); err != nil {
		t.Errorf("Error executing template %q", err)
	}
	if b.String() != "abc" {
		t.Errorf("Expecting output to be %q, got %q", "abc", b.String())
	}
}

func TestWithATextTemplate(t *testing.T) {
	tp := textTemplate.New("testValid")
	tpl, err := tp.Funcs(IncludeFunc(tp)).Parse(validTemplate)
	if err != nil {
		t.Errorf("Error parsing template %q", err)
	}
	var b strings.Builder
	if err := tpl.Execute(&b, nil); err != nil {
		t.Errorf("Error executing template %q", err)
	}
	if b.String() != "abc" {
		t.Errorf("Expecting output to be %q, got %q", "abc", b.String())
	}
}

func TestWithAnInvalidTemplate(t *testing.T) {
	tp := textTemplate.New("testInvalid")
	tpl, err := tp.Funcs(IncludeFunc(tp)).Parse(invalidTemplate)
	if err != nil {
		t.Errorf("Error parsing template %q", err)
	}
	var b strings.Builder
	if err := tpl.Execute(&b, nil); err == nil {
		t.Error("Expecting error, got nothing")
	}
}
