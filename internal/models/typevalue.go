package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

var _ ValueElementer = (*Typevalue)(nil)
var _ Cloner[*Typevalue] = (*Typevalue)(nil)

type Typevalue struct {
	Fieldvalue  Fieldtype    `json:"fieldvalue"`
	Validvalues []*Fieldtype `json:"-"`
}

// ValueSetter interface
func (i *Typevalue) SetValue(input ...any) error {
	for _, val := range input {
		switch ass := val.(type) {
		case string:
			for _, v := range i.Validvalues {
				if ass == v.Label {
					i.Fieldvalue = *v
					return nil
				}
			}
		case *Fieldtype:
			i.Fieldvalue = *ass
			return nil
		case []*Fieldtype:
			i.Validvalues = ass
			return nil
		default:
			return fmt.Errorf("invalid type %T for SetValue @Typevalue. Value of input: %s", ass, ass)
		}
	}
	return fmt.Errorf("unkown error for SetValue @Typevalue")
}

// Informer interface
func (t *Typevalue) GetInfo(key string) (string, error) {
	switch strings.ToLower(key) {
	case Description:
		return t.Fieldvalue.Description, nil
	case Id:
		return t.Fieldvalue.Id, nil
	case Identify:
		return t.Fieldvalue.Identify(), nil
	case Value:
		return t.Fieldvalue.Label, nil
	default:
		return "", fmt.Errorf("wrong key (%s) for GetInfo @Typevalue", key)
	}
}

// Executor interface
func (t *Typevalue) Execute() (any, error) {
	return nil, nil
}

// Cloner interface
func (t *Typevalue) Clone() *Typevalue {
	return &Typevalue{
		Fieldvalue:  t.Fieldvalue,
		Validvalues: t.Validvalues,
	}
}

func (t *Typevalue) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Fieldvalue)
}

func (t *Typevalue) UnmarshalJSON(data []byte) error {

	var jsonData map[string]string

	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	t.Fieldvalue = Fieldtype{
		Type:        jsonData["type"],
		Id:          jsonData["id"],
		Label:       jsonData["label"],
		Description: jsonData["description"],
	}
	return nil
}
