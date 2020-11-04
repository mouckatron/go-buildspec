package buildspec

import (
	"testing"
)

var bs = Buildspec{
	Version: "0.2",
	Env:     Environment{},
	Phases: map[string]*Phase{
		"install":    &Phase{Name: "install", Environment: map[string]string{}, Commands: []string{"echo Hello"}},
		"pre_build":  &Phase{Name: "pre_build", Environment: map[string]string{}, Commands: []string{}},
		"build":      &Phase{Name: "build", Environment: map[string]string{}, Commands: []string{}},
		"post_build": &Phase{Name: "post_build", Environment: map[string]string{}, Commands: []string{}},
	},
}

func TestLoadFromFileNoFile(t *testing.T) {
	_, err := LoadFromFile("nonexistant")

	if err == nil {
		t.Errorf("Was expecting file not found error")
	}
}

func TestLoadFromFileBuildspecYAML(t *testing.T) {

	subject, _ := LoadFromFile("../test/buildspec_test.yml")

	if bs.Version != subject.Version {
		t.Errorf("Version error: wanted %s, got %s", bs.Version, subject.Version)
	}
	if bs.Env != subject.Env {
		t.Errorf("Environment error: wanted %s, got %s", bs.Env, subject.Env)
	}
	if len(bs.Phases) != len(subject.Phases) {
		t.Errorf("Phase count error: wanted %d, got %d", len(bs.Phases), len(subject.Phases))
	}
}

func TestLoadFromFileBuildspec(t *testing.T) {

	subject, _ := LoadFromFile("../test/buildspec_test.yml")

	if len(bs.Phases) != len(subject.Phases) {
		t.Errorf("Phase count error: wanted %d, got %d", len(bs.Phases), len(subject.Phases))
	}
}
