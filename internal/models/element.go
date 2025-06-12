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

	OnValidated func(bool)          `json:"-"`
	Log         func(string, error) `json:"-"`
}

func (e *Element) RulesReset() {
	e.isDirty = true
}

func (e *Element) RulesApplied(validation bool, err string) {
	e.isDirty = !validation
	e.isValidated = validation
	e.ErrorMsg = err
	if e.OnValidated != nil {
		// e.OnValidated(validation)
	}
}

func (e *Element) Execute() {
	retval, err := e.Value.Execute()
	if err != nil {
		if e.Log != nil {
			e.Log(fmt.Sprintf("error @Element|Execute - type: %s, id: %s", e.Fieldtype.Type, e.Fieldtype.Id), err)
		}
		e.RulesApplied(false, fmt.Sprintf("%v", err))
		return
	}
	if retval != nil {
		e.SetValue(retval)
	}
}

func (e *Element) SetValue(input ...any) {
	err := e.Value.SetValue(input...)
	if err != nil && e.Log != nil {
		e.Log(fmt.Sprintf("error @Element|SetValue - type: %s, id: %s", e.Fieldtype.Type, e.Fieldtype.Id), err)
	}
	e.RulesReset()
}

func (e *Element) GetValueInfo(key string) string {
	s, err := e.Value.GetInfo(key)
	if err != nil && e.Log != nil {
		e.Log(fmt.Sprintf("error @Element|GetValue - type: %s, id: %s, key: %s", e.Fieldtype.Type, e.Fieldtype.Id, key), err)
	}

	return s
}

func (e *Element) GetValueAsInt() int {
	switch ass := e.Value.(type) {
	case *Intvalue:
		return ass.Intvalue
	case *Dice:
		return ass.Value
	}
	return 0
}

func (e *Element) GetValidation() bool {
	return e.isValidated
}

func (e *Element) Clone() *Element {

	newele := &Element{
		Fieldtype:   e.Fieldtype,
		Visibility:  e.Visibility,
		Editable:    e.Editable,
		ErrorMsg:    e.ErrorMsg,
		isValidated: e.isValidated,
		isDirty:     e.isDirty,
		OnValidated: e.OnValidated,
		Log:         e.Log,
	}

	switch ass := e.Value.(type) {
	case *Dice:
		newele.Value = CloneAny(ass)
	case *Intvalue:
		newele.Value = CloneAny(ass)
	case *Stringvalue:
		newele.Value = CloneAny(ass)
	case *Typevalue:
		newele.Value = CloneAny(ass)
	case *Skill:
		newele.Value = CloneAny(ass)
	default:
	}
	return newele
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
	} else if val["skilltype"] != nil {
		var i Skill
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
