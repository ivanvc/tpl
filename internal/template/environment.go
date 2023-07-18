package template

import (
	"encoding/json"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v2"

	"github.com/ivanvc/tpl/internal/io"
)

type environment any

func loadEnvironment(input []byte, forceJSON, forceTOML, forceYAML bool) environment {
	log.Default().SetLevel(log.DebugLevel)
	var env environment

	if len(input) > 0 && input[0] == '@' {
		f := string(input[1:])
		if _, err := os.Stat(f); err == nil {
			if input, err = io.ReadFile(f); err != nil {
				log.Warn("Error reading file", "error", err, "file", f)
			}
		} else {
			log.Debug("Error opening file", "error", err, "file", f)
		}
	}

	log.Debug("Loading environment", "input", string(input))
	if forceJSON {
		loadJSONEnvironment(input, &env)
		return env
	}
	if forceTOML {
		loadTOMLEnvironment(input, &env)
		return env
	}
	if forceYAML {
		loadYAMLEnvironment(input, &env)
		return env
	}

	if err := loadJSONEnvironment(input, &env); err == nil {
		return env
	}
	if err := loadTOMLEnvironment(input, &env); err == nil {
		return env
	}
	if err := loadYAMLEnvironment(input, &env); err == nil {
		return env
	}

	return env
}

func loadJSONEnvironment(input []byte, env *environment) error {
	if err := json.Unmarshal(input, &env); err == nil {
		log.Debug("Loaded JSON environment", "env", env)
		return nil
	} else {
		log.Debug("Error parsing JSON", "error", err)
		return err
	}
}

func loadTOMLEnvironment(input []byte, env *environment) error {
	if err := toml.Unmarshal(input, &env); err == nil {
		log.Debug("Loaded TOML environment", "env", env)
		return nil
	} else {
		log.Debug("Error parsing TOML", "error", err)
		return err
	}
}

func loadYAMLEnvironment(input []byte, env *environment) error {
	if err := yaml.Unmarshal(input, &env); err == nil {
		log.Debug("Loaded YAML environment", "env", env)
		return nil
	} else {
		log.Debug("Error parsing YAML", "error", err)
		return err
	}
}
