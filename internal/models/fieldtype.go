package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jc-design/rp_mgmt/internal/rules"
)

type FieldType struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func (e *FieldType) Identify() string {
	return fmt.Sprintf("%s|%s", e.Type, e.Id)
}

func (i *FieldType) SetValue(input any) {
	assert, ok := input.(FieldType)
	if ok {
		i.Type = assert.Type
		i.Id = assert.Id
		i.Label = assert.Label
		i.Description = assert.Description
	}
}

func (i *FieldType) ValueAsString() string {
	return i.Label
}

func (i *FieldType) AdditionalValueAsString() string {
	return i.Description
}

func (i *FieldType) Execute() {}

func LoadTypes(f rules.Folderstructure) ([]FieldType, error) {

	t := make([]FieldType, 0, 20)

	data, err := os.ReadFile(filepath.Join(f.Data, "types.json"))
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
