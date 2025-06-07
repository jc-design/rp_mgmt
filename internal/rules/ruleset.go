package rules

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	CreationRules string = "creation"
	LevelupRules  string = "levelup"
)

type Ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func LoadRuleSet(f *Folderstructure) (Ruleset, error) {

	r := Ruleset{}

	data, err := os.ReadFile(filepath.Join(f.Rules, "ruleset.json"))
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
