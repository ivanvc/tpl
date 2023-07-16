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
	Input            string
	Env              string
	UseStdin         bool
	SetDebugLogLevel bool
}

// Loads the configuration options.
func Load() *Config {
	c := new(Config)
	flag.StringVar(&c.Input, "input", "", "The template input to process.")
	flag.StringVar(&c.Env, "env", "", "The environment for the template (YAML, JSON or TOML).")
	flag.BoolVar(&c.UseStdin, "stdin", false, "Read template from stdin.")
	flag.BoolVar(&c.SetDebugLogLevel, "debug", false, "Set debug log level")
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

	if c.SetDebugLogLevel {
		log.Default().SetLevel(log.DebugLevel)
	}

	log.Debug("Loaded configuration", "config", c)
	return c
}
