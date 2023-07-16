package template

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v2"

	"github.com/ivanvc/tpl/internal/config"
	"github.com/ivanvc/tpl/internal/io"
)

type environment map[string]any

func loadEnvironment(cfg *config.Config) environment {
	var env environment
	input := []byte(cfg.Env)

	if cfg.Env[0] == '@' {
		if _, err := os.Stat(cfg.Env[1:]); err == nil {
			input = io.ReadFile(cfg.Env[1:])
		} else {
			log.Debug("Error opening file", "error", err, "file", cfg.Env[1:])
		}
	}

	if err := json.Unmarshal(input, &env); err == nil {
		return env
	} else {
		log.Debug("Error parsing JSON", "error", err)
	}
	if err := yaml.Unmarshal(input, &env); err == nil {
		return env
	} else {
		log.Debug("Error parsing YAML", "error", err)
	}
	if err := toml.Unmarshal(input, &env); err != nil {
		log.Debug("Error parsing TOML", "error", err)
	}

	return env
}
