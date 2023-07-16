package config

import (
	"flag"
)

// Config represents the configuration options for the program.
type Config struct {
	Input            string
	InputFile        string
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

	c.InputFile = flag.Arg(0)

	return c
}
