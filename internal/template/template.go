package template

import (
	"io"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/charmbracelet/log"

	"github.com/ivanvc/tpl/internal/config"
	intio "github.com/ivanvc/tpl/internal/io"
	tpl "github.com/ivanvc/tpl/pkg/template"
)

// Excutes the template to generate the output.
func Execute(cfg *config.Config) {
	input, err := loadInput(cfg)
	if err != nil {
		log.Fatal("Error loading input template", "error", err)
	}
	if len(input) == 0 {
		log.Warn("Empty input template")
	}
	log.Debug("Loaded input template", "tpl", input)

	env := loadEnvironment([]byte(cfg.Env))

	t := template.New("__tpl_input")
	tpl := template.Must(
		t.Funcs(sprig.TxtFuncMap()).
			Funcs(tpl.IncludeFunc(t)).
			Funcs(tpl.FileSizeFunc).
			Parse(input),
	)
	if err := tpl.Execute(os.Stdout, env); err != nil {
		log.Fatal("Error executing template", "error", err)
	}
}

func loadInput(cfg *config.Config) (string, error) {
	var input string
	if len(cfg.InputFile) > 0 {
		b, err := intio.ReadFile(cfg.InputFile)
		if err != nil {
			return input, err
		}
		input = string(b)
	} else if cfg.UseStdin {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return input, err
		}
		input = string(b)
	} else {
		input = cfg.Input
	}
	return input, nil
}
