package template

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v2"

	"github.com/ivanvc/tpl/internal/config"
	"github.com/ivanvc/tpl/internal/io"
)

type environment map[string]any

func loadEnvironment(cfg *config.Config) environment {
	var (
		env   environment
		input []byte
	)
	if len(cfg.YAML) > 0 {
		if cfg.YAML[0] == '@' {
			input = io.ReadFile(cfg.YAML[1:])
		} else {
			input = []byte(cfg.YAML)
		}
		if err := yaml.Unmarshal(input, &env); err != nil {
			log.Fatal("Error parsing YAML", "error", err)
		}
	}

	if len(cfg.JSON) > 0 {
		if cfg.JSON[0] == '@' {
			input = io.ReadFile(cfg.JSON[1:])
		} else {
			input = []byte(cfg.JSON)
		}
		if err := json.Unmarshal(input, &env); err != nil {
			log.Fatal("Error parsing JSON", "error", err)
		}
	}

	if len(cfg.TOML) > 0 {
		if cfg.TOML[0] == '@' {
			input = io.ReadFile(cfg.TOML[1:])
		} else {
			input = []byte(cfg.TOML)
		}
		if err := toml.Unmarshal(input, &env); err != nil {
			log.Fatal("Error parsing TOML", "error", err)
		}
	}

	return env
}
