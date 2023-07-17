package config

import (
	"flag"
	"fmt"
	"os"
)

// Config represents the configuration options for the program.
type Config struct {
	Input            string
	InputFile        string
	Env              string
	UseStdin         bool
	SetDebugLogLevel bool
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			`Usage: %s [options] [template input]:

Executes the template from the input applying the environment passed in options.

If you specify the input template via an option flag (-input), then it will
read it inline.
If you specify that the template comes from the stdin by setting the option flag
(-stdin), then it will read it inline from the stdin.
If you set it as the first argument, it assumes that it is a file.

For the environment (-env) it will expect it as inline data. However, if you
start it with @, it will assume it is a file.

The output is sent to stdout.

Options:
`,
			os.Args[0],
		)
		flag.PrintDefaults()
	}
}

// Loads the configuration options.
func Load() *Config {
	c := new(Config)
	flag.StringVar(&c.Input, "input", "", "The template input to process.")
	flag.StringVar(&c.Env, "env", "", "The environment for the template (YAML, JSON or TOML).")
	flag.BoolVar(&c.UseStdin, "stdin", false, "Read template from stdin.")
	flag.BoolVar(&c.SetDebugLogLevel, "debug", false, "Set log level to debug.")
	flag.Parse()

	c.InputFile = flag.Arg(0)

	return c
}
