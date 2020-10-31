package buildspec

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// BuildspecYAML root object for YAML
type BuildspecYAML struct {
	Version string
	Env     Environment
	Phases  map[string][]string `yaml:",flow"`
}

// Buildspec final object after changing to Phase type
type Buildspec struct {
	Version string
	Env     Environment
	Phases  map[string]Phase
}

// Environment object for YAML
type Environment struct {
}

// Phase object for Buildspec Phases
type Phase struct {
	Name     string
	Commands []string
}

// LoadFromFile create a Buildspec object from a YAML file
func LoadFromFile(filepath string) (bs Buildspec, err error) {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	bsy := BuildspecYAML{}
	yaml.Unmarshal(data, &bsy)

	bs = buildspecFromBuildspecYAML(&bsy)

	return
}

func buildspecFromBuildspecYAML(bsy *BuildspecYAML) (bs Buildspec) {
	bs.Version = bsy.Version
	bs.Env = bsy.Env
	bs.Phases = make(map[string]Phase)

	for key, commands := range bsy.Phases {
		bs.Phases[key] = Phase{
			Name:     key,
			Commands: commands,
		}
	}
	return
}
