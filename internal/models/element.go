package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Element struct {
	Fieldtype   Fieldtype      `json:"type"`
	Value       ValueElementer `json:"value"`
	Visibility  Activationmode `json:"visibility"`
	Editable    Activationmode `json:"editable"`
	ErrorMsg    string         `json:"-"`
	isValidated bool
	isDirty     bool
}

func (e *Element) RulesReset() {
	e.isDirty = true
}

func (e *Element) RulesApplied(validation bool, err string) {
	e.isDirty = false
	e.isValidated = validation
	e.ErrorMsg = err
}

func (e *Element) Execute() {
	retval, err := e.Value.Execute()
	if err != nil {
		e.RulesApplied(false, fmt.Sprintf("%v", err))
		return
	}
	if retval != nil {
		e.SetValue(retval)
	}
	e.RulesReset()
}

func (e *Element) SetValue(input ...any) {
	e.Value.SetValue(input...)
	e.RulesReset()
}

func (e *Element) GetValidation() bool {
	return e.isValidated
}

func (e *Element) UnmarshalJSON(data []byte) error {

	var jsonData map[string]interface{}

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	var t Fieldtype
	err = mapstructure.Decode(jsonData["type"], &t)
	if err == nil && t.Type != "" {
		e.Fieldtype = t
	} else if err == nil && len(strings.TrimSpace(t.Type)) == 0 {
		return errors.New("[type] must be set")
	} else {
		return errors.New("failed to create [type]")
	}

	val, ok := jsonData["value"].(map[string]interface{})
	if !ok {
		return errors.New("failed to create [value]")
	}

	if val["intvalue"] != nil {
		var i Intvalue
		err := mapstructure.Decode(val, &i)
		if err == nil {
			e.Value = &i
		}
	} else if val["stringvalue"] != nil {
		var i Stringvalue
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
		var i = Typevalue{}
		err := mapstructure.Decode(val, &i.Fieldvalue)
		if err == nil {
			e.Value = &i
		}
	} else {
		return errors.New("[value] must be set")
	}

	vis, ok := jsonData["visibility"].(string)
	if ok {
		e.Visibility.FromString(vis)
	}
	edit, ok := jsonData["editable"].(string)
	if ok {
		e.Editable.FromString(edit)
	}

	return nil
}
