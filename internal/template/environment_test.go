package template

import (
	"testing"
)

func TestLoadingInvalidEnvironment(t *testing.T) {
	env := loadEnvironment([]byte("yaml: yaml: yaml"))
	if env != nil {
		t.Errorf("Expecting type to be nil, got %v", env)
	}
}

func TestLoadingAYAMLEnvironment(t *testing.T) {
	env := loadEnvironment([]byte("test: true"))
	if !env.(map[any]any)["test"].(bool) {
		t.Errorf("Expecting parsed YAML to be true, got %v", env.(map[any]any)["test"])
	}
	env = loadEnvironment([]byte("- a"))
	if env.([]any)[0].(string) != "a" {
		t.Errorf("Expecting parsed YAML to be %q, got %q", "a", env.([]any)[0])
	}
}

func TestLoadingJSONEnvironment(t *testing.T) {
	env := loadEnvironment([]byte(`{"test":true}`))
	if !env.(map[string]any)["test"].(bool) {
		t.Errorf("Expecting parsed JSON to be true, got %v", env.(map[string]any)["test"])
	}
	env = loadEnvironment([]byte(`["a"]`))
	if env.([]any)[0].(string) != "a" {
		t.Errorf("Expecting parsed JSON to be %q, got %q", "a", env.([]any)[0])
	}
}

func TestLoadingTOMLEnvironment(t *testing.T) {
	env := loadEnvironment([]byte(`test=true`))
	if !env.(map[string]any)["test"].(bool) {
		t.Errorf("Expecting parsed TOML to be true, got %v", env.(map[string]any)["test"])
	}
}
