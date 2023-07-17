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

type environment any

func loadEnvironment(cfg *config.Config) environment {
	var env environment
	input := []byte(cfg.Env)

	if len(cfg.Env) > 0 && cfg.Env[0] == '@' {
		f := cfg.Env[1:]
		if _, err := os.Stat(f); err == nil {
			if input, err = io.ReadFile(f); err != nil {
				log.Warn("Error reading file", "error", err, "file", f)
			}
		} else {
			log.Debug("Error opening file", "error", err, "file", f)
		}
	}

	log.Debug("Loading environment", "input", string(input))
	if err := json.Unmarshal(input, &env); err == nil {
		log.Debug("Loaded JSON environment", "env", env)
		return env
	} else {
		log.Debug("Error parsing JSON", "error", err)
	}
	if err := toml.Unmarshal(input, &env); err == nil {
		log.Debug("Loaded TOML environment", "env", env)
		return env
	} else {
		log.Debug("Error parsing TOML", "error", err)
	}
	if err := yaml.Unmarshal(input, &env); err != nil {
		log.Debug("Error parsing YAML", "error", err)
	}

	log.Debug("Loaded YAML environment", "env", env)
	return env
}
