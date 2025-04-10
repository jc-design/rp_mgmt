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

type RuleSet struct {
	Name    string
	Version string
}

func LoadRuleSet(f Folderstructure) (RuleSet, error) {

	r := RuleSet{}

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

func LoadRules(f Folderstructure, rule string) ([]byte, error) {

	data, err := os.ReadFile(filepath.Join(f.Rules, rule))
	if err != nil {
		return nil, err
	}

	return data, nil
}
