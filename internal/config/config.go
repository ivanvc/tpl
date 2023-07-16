package config

import (
	"flag"
	"io"
	"os"

	"github.com/charmbracelet/log"

	intio "github.com/ivanvc/tpl/internal/io"
)

// Config represents the configuration options for the program.
type Config struct {
	Input    string
	JSON     string
	TOML     string
	YAML     string
	UseStdin bool
}

// Loads the configuration options.
func Load() *Config {
	c := new(Config)
	flag.StringVar(&c.Input, "input", "", "The template input to process.")
	flag.StringVar(&c.JSON, "json", "", "The JSON environment for the template.")
	flag.StringVar(&c.TOML, "toml", "", "The TOML environment for the template.")
	flag.StringVar(&c.YAML, "yaml", "", "The YAML environment for the template.")
	flag.BoolVar(&c.UseStdin, "stdin", false, "Read template from stdin.")
	flag.Parse()

	if len(c.Input) == 0 {
		c.Input = string(intio.ReadFile(flag.Arg(0)))
	}
	if c.UseStdin {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Could not read input template")
		}
		c.Input = string(b)
	}

	return c
}
