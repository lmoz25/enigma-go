package enigma

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/monzo/terrors"

	"github.com/lmoz25/enigma-go/git_projects/enigma-go/components"
)

type Enigma struct {
	encoders []components.Encoder
}

type Config struct {
	Components []ComponentConfig `json:components`
}

type ComponentConfig struct {
	Name        string        `json:name`
	Type        string        `json:type`
	Model       string        `json:model,omitempty`
	WiringTable map[rune]rune `json:wiring_table,omitempty`
}

func (e *Enigma) Encode(message string) (encoded string) {
	message = strings.ToUpper(message)
	spaceCounter := 1
	for _, char := range message {
		for _, encoder := range e.encoders {
			char = encoder.Encode(char)
		}
		if spaceCounter%5 == 0 {
			encoded = encoded + " "
		}
		encoded = encoded + string(char)
		spaceCounter++
	}
	return
}

func NewEnigma(configPath string) (*Enigma, error) {
	errParams := map[string]string{
		"filename": configPath,
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, terrors.Augment(err, "failed to read enigma config file", errParams)
	}
	config := Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		errParams["config"] = string(file)
		return nil, terrors.Augment(err, "failed to parse enigma config file", errParams)
	}

	encoders := []components.Encoder{}
	for _, component := range config.Components {
		encoders = append(encoders, NewEncoder(component))
	}
}

func NewEncoder(component ComponentConfig) components.Encoder {
	switch component.Type:
	case "rotor"
}