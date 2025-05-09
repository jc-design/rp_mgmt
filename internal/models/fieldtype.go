package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jc-design/rp_mgmt/internal/rules"
)

type Fieldtype struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func (f *Fieldtype) Identify() string {
	return fmt.Sprintf("%s|%s", f.Type, f.Id)
}

func LoadTypes(f rules.Folderstructure) ([]Fieldtype, error) {

	t := make([]Fieldtype, 0, 20)

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
