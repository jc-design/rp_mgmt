package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jc-design/rp_mgmt/internal/rules"
)

type FieldType struct {
	Type        string
	Id          string
	Label       string
	Description string
}

func (e *FieldType) Identify() string {
	return fmt.Sprintf("%s|%s", e.Type, e.Id)
}

func LoadTypes(f rules.Folderstructure) ([]FieldType, error) {

	t := make([]FieldType, 0, 20)

	data, err := os.ReadFile(filepath.Join(f.Data, "fields.json"))
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(data, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
