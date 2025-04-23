package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/mitchellh/mapstructure"
)

type Character struct {
	Id                      string        `json:"id"`
	Name                    string        `json:"name"`
	Image                   string        `json:"image"`
	RuleSet                 rules.RuleSet `json:"ruleset"`
	Properties              []Element     `json:"properties"`
	propertiesValidElements []FieldType
}

func NewCharacter(r rules.RuleSet, prop []Element) *Character {

	copyElements := make([]Element, len(prop))
	copy(copyElements, prop)
	return &Character{
		uuid.New().String(),
		"New Character",
		"",
		r,
		copyElements,
		make([]FieldType, 0),
	}
}

func LoadCharacters(f rules.Folderstructure, r rules.RuleSet) ([]Character, error) {
	var cs []Character

	data, err := os.ReadFile(filepath.Join(f.Characters, "characters.json"))
	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(data, &cs)
	if err != nil {
		return cs, err
	}

	for _, c := range cs {
		if c.RuleSet != r {
			return cs, fmt.Errorf("error loading characters due to incomapatible ruleset. "+
				"Loaded ruleset: %s, requested ruleset: %s", c.RuleSet, r)
		}
	}
	return cs, nil
}

func SaveCharacters(f rules.Folderstructure, c *[]Character) error {

	bytes, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(f.Characters, "characters.json"), bytes, 0644); err != nil {
		return err
	}

	return nil
}

func LoadElements(f rules.Folderstructure) ([]Element, error) {

	e := make([]Element, 0, 20)

	data, err := os.ReadFile(filepath.Join(f.Data, "characterproperties.json"))
	if err != nil {
		return e, err
	}

	err = json.Unmarshal(data, &e)
	if err != nil {
		return e, err
	}

	return e, nil
}

func (c *Character) UnmarshalJSON(data []byte) error {

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	if s, ok := jsonData["Id"].(string); ok {
		c.Id = s
	} else {
		return fmt.Errorf("loading aborted. No field 'Id' found")
	}

	if s, ok := jsonData["Name"].(string); ok {
		c.Name = s
	} else {
		return fmt.Errorf("loading aborted. No field 'Name' founf")
	}

	err = mapstructure.Decode(jsonData["RuleSet"], &c.RuleSet)
	if err != nil {
		return err
	}

	err = mapstructure.Decode(jsonData["Properties"], &c.Properties)
	if err != nil {
		return err
	}

	return nil
}

// func (c *Character) MarshalJSON() ([]byte, error) {

// 	return json.Marshal(map[string]interface{}{
// 		"id":         c.Id,
// 		"name":       c.Name,
// 		"ruleset":    c.RuleSet,
// 		"properties": c.Properties,
// 	})
// }

// func elementsToMap(src *[]Element) map[string]Element {
// 	var result = make(map[string]Element)
// 	for _, v := range *src {
// 		result[v.Identify()] = v
// 	}
// 	return result
// }

// func mapToElements(src *map[string]Element) []Element {
// 	var result = make([]Element, 0, 10)
// 	for _, v := range *src {
// 		result = append(result, v)
// 	}
// 	return result
// }

// Func needed for RuleFact interface
// it's needed for the grule-engine validation
func (c *Character) FactKey() string {
	return "Character"
}
