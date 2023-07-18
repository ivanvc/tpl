package template

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"text/template"
)

func TestFileSize(t *testing.T) {
	tt := [][]string{
		[]string{"test", "0 B"},
		[]string{"-4096", "0 B"},
		[]string{"4096", "4.1 kB"},
		[]string{strconv.FormatUint(math.MaxUint64, 10), "18 EB"},
	}
	for _, c := range tt {
		if res := fileSize(c[0]); res != c[1] {
			t.Errorf("Expecting %q, got %q", c[1], res)
		}
	}
}

func TestFileSizeInTemplate(t *testing.T) {
	tpl, err := template.New("fileSize").
		Funcs(FileSizeFunc).
		Parse(`{{ fileSize 4096 }},{{ fileSize "4096" }},{{ fileSize .input }}`)
	if err != nil {
		t.Errorf("Error parsing template %q", err)
	}
	var b strings.Builder
	if err := tpl.Execute(&b, map[string]any{"input": "1024"}); err != nil {
		t.Errorf("Error executing template %q", err)
	}
	if expected := "4.1 kB,4.1 kB,1.0 kB"; b.String() != expected {
		t.Errorf("Expecting output to be %q, got %q", expected, b.String())
	}
}
