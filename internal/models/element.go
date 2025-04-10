package models

import (
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

type Element struct {
	Type          FieldType
	ReferenceType FieldType
	Value         any
	Visibility    Visibility
	isValidated   bool
	isDirty       bool
	errorMsg      string
}

func (e *Element) Identify() string {
	return e.Type.Identify()
}

func (e *Element) UnmarshalJSON(data []byte) error {

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	var t FieldType
	err = mapstructure.Decode(jsonData["Type"], &t)
	if err == nil && t.Type != "" {
		e.Type = t
	} else {
		return errors.New("failed to create field named [Type]")
	}

	var rt FieldType
	err = mapstructure.Decode(jsonData["ReferenceType"], &rt)
	if err == nil {
		e.ReferenceType = rt
	}

	switch jsonData["Value"].(type) {
	case int:
		e.Value = jsonData["Value"].(int)
	case int8:
		e.Value = int(jsonData["Value"].(int8))
	case int16:
		e.Value = int(jsonData["Value"].(int16))
	case int32:
		e.Value = int(jsonData["Value"].(int32))
	case int64:
		e.Value = int(jsonData["Value"].(int64))
	case uint:
		e.Value = jsonData["Value"].(uint)
	case uint8:
		e.Value = int(jsonData["Value"].(uint8))
	case uint16:
		e.Value = int(jsonData["Value"].(uint16))
	case uint32:
		e.Value = int(jsonData["Value"].(uint32))
	case uint64:
		e.Value = int(jsonData["Value"].(uint64))
	case float32:
		e.Value = int(jsonData["Value"].(float32))
	case float64:
		e.Value = int(jsonData["Value"].(float64))
	case string:
		e.Value = jsonData["Value"].(string)
	case map[string]interface{}:
		var d Dice
		err := mapstructure.Decode(jsonData["Value"], &d)
		if err == nil {
			e.Value = d
			return nil
		}
		var f FieldType
		err = mapstructure.Decode(jsonData["Value"], &f)
		if err == nil {
			e.Value = f
			return nil
		}
		e.Value = nil
		return nil
	default:
		e.Value = nil
	}

	vis, ok := jsonData["Visibility"].(string)
	if ok {
		e.Visibility.FromString(vis)
	}

	return nil
}
