package template

import (
	"html/template"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/charmbracelet/log"

	"github.com/ivanvc/tpl/internal/config"
	tpl "github.com/ivanvc/tpl/pkg/template"
)

// Runs the template to generate the output.
func Run(cfg *config.Config) {
	env := loadEnvironment(cfg)
	input := cfg.Input
	t := template.New("input")
	tpl := template.Must(t.Funcs(sprig.TxtFuncMap()).Funcs(tpl.IncludeFunc(t)).Parse(input))
	if err := tpl.Execute(os.Stdout, env); err != nil {
		log.Fatal("Error executing template", "error", err)
	}
}
