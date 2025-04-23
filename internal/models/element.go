package models

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/mitchellh/mapstructure"
)

type Element struct {
	Type          FieldType      `json:"type"`
	ReferenceType FieldType      `json:"referencetype"`
	Value         ValueElementer `json:"value"`
	Visibility    Visibility     `json:"visibility"`
	isValidated   bool
	isDirty       bool
	errorMsg      string
}

func (e *Element) Identify() string {
	return e.Type.Identify()
}

func (e *Element) RulesReset() {
	e.isDirty = true
	e.isValidated = false
	e.errorMsg = ""
}

func (e *Element) RulesApplied(validation bool, err string) {
	e.isDirty = false
	e.isValidated = validation
	e.errorMsg = err
}

func (e *Element) UnmarshalJSON(data []byte) error {

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	var t FieldType
	err = mapstructure.Decode(jsonData["type"], &t)
	if err == nil && t.Type != "" {
		e.Type = t
	} else if err == nil && len(strings.TrimSpace(t.Type)) == 0 {
		return errors.New("field [type] missing")
	} else {
		return errors.New("failed to create field named [type]")
	}

	var rt FieldType
	err = mapstructure.Decode(jsonData["referencetype"], &rt)
	if err == nil {
		e.ReferenceType = rt
	}

	val, ok := jsonData["value"].(map[string]interface{})
	if !ok {
		return errors.New("field [value] missing")
	}

	if val["intvalue"] != nil {
		var i IntValue
		err := mapstructure.Decode(val, &i)
		if err == nil {
			e.Value = &i
		}
	} else if val["stringvalue"] != nil {
		var i StringValue
		err := mapstructure.Decode(val, &i)
		if err == nil {
			e.Value = &i
		}
	} else if val["dicevalue"] != nil {
		var i Dice
		err := mapstructure.Decode(val, &i)
		if err == nil {
			e.Value = &i
		}
	} else if val["type"] != nil {
		var i FieldType
		err := mapstructure.Decode(val, &i)
		if err == nil {
			e.Value = &i
		}
	} else {
		return errors.New("field [value] missing")
	}

	vis, ok := jsonData["visibility"].(string)
	if ok {
		e.Visibility.FromString(vis)
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
