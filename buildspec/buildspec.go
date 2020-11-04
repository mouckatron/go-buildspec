package buildspec

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Buildspec final object after changing to Phase type
type Buildspec struct {
	Version string
	Env     Environment
	Phases  map[string]*Phase
}

// SetPhaseNames A convenience to set the name of the phase within the Phase struct
func (b *Buildspec) SetPhaseNames() {
	for key := range b.Phases {
		b.Phases[key].Name = key
	}
}

// Environment object for YAML
type Environment struct {
}

// Phase object for Buildspec Phases
type Phase struct {
	Name        string
	Environment map[string]string `yaml:"environment"`
	Commands    []string          `yaml:"commands,flow"`
}

// LoadFromFile create a Buildspec object from a YAML file
func LoadFromFile(filepath string) (bs Buildspec, err error) {

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	yaml.Unmarshal(data, &bs)
	bs.SetPhaseNames()

	return
}
