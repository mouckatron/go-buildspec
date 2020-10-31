package buildspec

import (
	"testing"
)

var bsy = BuildspecYAML{
	Version: "0.2",
	Env:     Environment{},
	Phases: map[string][]string{
		"install":    {"echo Hello"},
		"pre_build":  {},
		"build":      {},
		"post_build": {},
	},
}

var bs = Buildspec{
	Version: "0.2",
	Env:     Environment{},
	Phases: map[string]Phase{
		"install":    Phase{Name: "install", Commands: []string{"echo Hello"}},
		"pre_build":  Phase{Name: "pre_build", Commands: []string{}},
		"build":      Phase{Name: "build", Commands: []string{}},
		"post_build": Phase{Name: "post_build", Commands: []string{}},
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

	if bsy.Version != subject.Version {
		t.Errorf("Version error: wanted %s, got %s", bsy.Version, subject.Version)
	}
	if bsy.Env != subject.Env {
		t.Errorf("Environment error: wanted %s, got %s", bsy.Env, subject.Env)
	}
	if len(bsy.Phases) != len(subject.Phases) {
		t.Errorf("Phase count error: wanted %d, got %d", len(bsy.Phases), len(subject.Phases))
	}
}

func TestLoadFromFileBuildspec(t *testing.T) {

	subject, _ := LoadFromFile("../test/buildspec_test.yml")

	if len(bs.Phases) != len(subject.Phases) {
		t.Errorf("Phase count error: wanted %d, got %d", len(bs.Phases), len(subject.Phases))
	}
}

func TestbuildspecFromBuildspecYAML(t *testing.T) {

	subject := buildspecFromBuildspecYAML(&bsy)

	if len(bs.Phases) != len(subject.Phases) {
		t.Errorf("Incorrect length: wanted %d, got %d", len(bs.Phases), len(subject.Phases))
	}

	for key, phase := range bs.Phases {
		if subject.Phases[key].Name != phase.Name {
			t.Errorf("Name did not match: wanted %s, got %s", phase.Name, subject.Phases[key].Name)
		}
		if len(subject.Phases[key].Commands) != len(phase.Commands) {
			t.Errorf("Phase command length mismatch: wanted %d, got %d", len(phase.Commands), len(subject.Phases[key].Commands))
		}
	}

}
