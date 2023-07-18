package template

import (
	"testing"
)

func TestLoadingInvalidEnvironment(t *testing.T) {
	env := loadEnvironment([]byte("yaml: yaml: yaml"), false, false, false)
	if env != nil {
		t.Errorf("Expecting type to be nil, got %v", env)
	}
}

func TestLoadingAYAMLEnvironment(t *testing.T) {
	env := loadEnvironment([]byte("test: true"), false, false, false)
	if !env.(map[any]any)["test"].(bool) {
		t.Errorf("Expecting parsed YAML to be true, got %v", env.(map[any]any)["test"])
	}
	env = loadEnvironment([]byte("- a"), false, false, false)
	if env.([]any)[0].(string) != "a" {
		t.Errorf("Expecting parsed YAML to be %q, got %q", "a", env.([]any)[0])
	}
}

func TestLoadingJSONEnvironment(t *testing.T) {
	env := loadEnvironment([]byte(`{"test":true}`), false, false, false)
	if !env.(map[string]any)["test"].(bool) {
		t.Errorf("Expecting parsed JSON to be true, got %v", env.(map[string]any)["test"])
	}
	env = loadEnvironment([]byte(`["a"]`), false, false, false)
	if env.([]any)[0].(string) != "a" {
		t.Errorf("Expecting parsed JSON to be %q, got %q", "a", env.([]any)[0])
	}
}

func TestLoadingTOMLEnvironment(t *testing.T) {
	env := loadEnvironment([]byte(`test=true`), false, false, false)
	if !env.(map[string]any)["test"].(bool) {
		t.Errorf("Expecting parsed TOML to be true, got %v", env.(map[string]any)["test"])
	}
}

func TestForceLoadingInvalidEnvironments(t *testing.T) {
	tt := [][]any{
		[]any{"", true, false, false},
		[]any{":", true, false, false},
		[]any{"{'json':true}", true, false, false},
		[]any{":", false, true, false},
		[]any{"foo", false, true, false},
		[]any{":", false, false, true},
	}
	for _, c := range tt {
		if e := loadEnvironment([]byte(c[0].(string)), c[1].(bool), c[2].(bool), c[3].(bool)); e != nil {
			t.Errorf("Expecting loaded environment to be nil, got %v", e)
		}
	}
}

func TestForceLoadingJSONEnvironment(t *testing.T) {
	if e := loadEnvironment([]byte("[1,2,3]"), true, false, false); e == nil {
		t.Errorf("Expecting loaded environment to not be nil, got %v", e)
	}
}

func TestForceLoadingYAMLEnvironment(t *testing.T) {
	if e := loadEnvironment([]byte(`{"not":"json"}`), false, false, true); e == nil {
		t.Errorf("Expecting loaded environment to not be nil, got %v", e)
	}
}

func TestForceLoadingTOMLEnvironment(t *testing.T) {
	if e := loadEnvironment([]byte("[1]"), false, true, false); e == nil {
		t.Errorf("Expecting loaded environment to not be nil, got %v", e)
	}
}
