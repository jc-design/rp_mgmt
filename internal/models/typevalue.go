package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Typevalue struct {
	Fieldvalue  Fieldtype    `json:"fieldvalue"`
	Validvalues []*Fieldtype `json:"-"`
}

func (i *Typevalue) SetValue(input ...any) {
	for _, val := range input {
		switch ass := val.(type) {
		case string:
			for _, v := range i.Validvalues {
				if ass == v.Label {
					i.Fieldvalue = *v
					return
				}
			}
		case *Fieldtype:
			i.Fieldvalue = *ass
		case []*Fieldtype:
			i.Validvalues = ass
		default:
			fmt.Printf("could not work with input of type %T", ass)
		}
	}
}

func (t *Typevalue) GetInfo(key string) string {
	switch strings.ToLower(key) {
	case description:
		return t.Fieldvalue.Description
	case id:
		return t.Fieldvalue.Id
	case identify:
		return t.Fieldvalue.Identify()
	case value:
		return t.Fieldvalue.Label
	default:
		return ""
	}
}

func (t *Typevalue) Execute() (any, error) {
	return nil, nil
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
