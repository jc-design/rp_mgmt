package models

import (
	"strings"

	"slices"

	"github.com/google/uuid"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

type Character struct {
	Id            string         `json:"id"`
	Name          string         `json:"name"`
	Image         string         `json:"image"`
	Level         int            `json:"level"`
	Exp           int            `json:"exp"`
	RuleSet       rules.Ruleset  `json:"ruleset"`
	Properties    []*Element     `json:"properties"`
	Status        Activationmode `json:"-"`
	Deleted       bool           `json:"-"`
	Allfieldtypes []*Fieldtype   `json:"-"`
}

// Create a new Character
func NewCharacter(r rules.Ruleset, prop []*Element, types []*Fieldtype) *Character {

	c := Character{}
	copyElements := make([]*Element, len(prop))
	copy(copyElements, prop)

	c.Id = uuid.New().String()
	c.Name = "New Character"
	c.Image = ""
	c.Level = 1
	c.Exp = 0
	c.RuleSet = r
	c.Properties = copyElements
	c.Status = Activationmode(Creation)
	c.Allfieldtypes = types

	return &c
}

// Func needed for RuleFact interface
// it's needed for the grule-engine validation
func (c *Character) FactKey() string {
	return "Character"
}

func (c *Character) GetElement(ident string) *Element {
	for i, e := range c.Properties {
		if e.Fieldtype.Identify() == ident {
			return c.Properties[i]
		}
	}
	return nil
}

func (c *Character) IsElementDirty(ident string) bool {
	e := c.GetElement(ident)
	if e == nil {
		return false
	}

	return e.isDirty
}

func (c *Character) GetValueInfo(ident, key string) string {
	e := c.GetElement(ident)
	if e == nil {
		return ""
	}

	return e.Value.GetInfo(key)
}

func (c *Character) GetValueAsInt(ident string) int {
	e := c.GetElement(ident)
	if e == nil {
		return 0
	}

	switch ass := e.Value.(type) {
	case *Intvalue:
		return ass.Intvalue
	case *Dice:
		return ass.Value
	}
	return 0
}

func (c *Character) IsValueInRange(ident string, min, max int64) bool {
	e := c.GetElement(ident)
	if e == nil {
		return false
	}

	switch ass := e.Value.(type) {
	case *Intvalue:
		if ass.Intvalue >= int(min) && ass.Intvalue <= int(max) {
			return true
		}
	case *Dice:
		if ass.Value >= int(min) && ass.Value <= int(max) {
			return true
		}
	}
	return false
}

func (c *Character) IsValueInList(ident string, list string) bool {
	e := c.GetElement(ident)
	if e == nil {
		return false
	}
	l := strings.Split(list, ";")
	if slices.Contains(l, e.Value.GetInfo(id)) || slices.Contains(l, e.Value.GetInfo(value)) {
		return true
	}

	return false
}

func (c *Character) SetValueFromList(ident, fieldtype, list string) {
	e := c.GetElement(ident)
	if e == nil {
		return
	}

	l := strings.Split(list, ";")
	types := make([]*Fieldtype, 0)
	for _, field := range c.Allfieldtypes {
		if fieldtype == field.Type && slices.Contains(l, field.Id) {
			types = append(types, field)
		}
	}

	switch ass := e.Value.(type) {
	case *Stringvalue:
		for _, field := range types {
			if field.Label == ass.Stringvalue {
				return
			}
		}
		e.SetValue(types[0].Label)
	case *Typevalue:
		ass.Validvalues = types

		for _, field := range types {
			if field.Id == ass.GetInfo(id) {
				return
			}
		}

		e.SetValue(types[0])
	}
}

func (c *Character) SetDiceProperties(ident string, dicevalue, dicecount int64, dicemarkup float64) {
	e := c.GetElement(ident)
	if e == nil {
		return
	}

	arr := []int{int(dicevalue), int(dicecount), int(dicemarkup)}
	e.SetValue(arr)
}
