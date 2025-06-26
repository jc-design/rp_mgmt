package models

import (
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

type elementgroup struct {
	Group    Fieldtype
	Elements []*Element
}

type Character struct {
	Id            string              `json:"id"`
	Name          string              `json:"name"`
	Image         string              `json:"image"`
	Level         int                 `json:"level"`
	Exp           int                 `json:"exp"`
	RuleSet       rules.Ruleset       `json:"ruleset"`
	Properties    []*Element          `json:"properties"`
	Status        Activationmode      `json:"-"`
	Allfieldtypes []*Fieldtype        `json:"-"`
	PropsGrouped  []*elementgroup     `jsonn:"-"`
	Log           func(string, error) `json:"-"`
}

// Create a new Character
func NewCharacter(r rules.Ruleset, prop []*Element, types []*Fieldtype, log func(string, error)) *Character {

	c := Character{}
	copyElements := make([]*Element, len(prop))
	for i, ref := range prop {
		e := ref.Clone()
		copyElements[i] = e
	}
	// copy(copyElements, prop)

	c.Id = uuid.New().String()
	c.Name = "New Character"
	c.Image = ""
	c.Level = 1
	c.Exp = 0
	c.RuleSet = r
	c.Properties = copyElements
	c.Status = Activationmode(Creation)
	c.Allfieldtypes = types
	c.Log = log

	c.ClassifyProperties()
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

	return !(!e.isDirty && e.isValidated)
}

func (c *Character) GetValueInfo(ident, key string) string {
	e := c.GetElement(ident)
	if e == nil {
		return ""
	}
	return e.GetValueInfo(key)
}

func (c *Character) GetValueAsInt(ident string) int {
	e := c.GetElement(ident)
	if e == nil {
		return 0
	}

	return e.GetValueAsInt()
}

func (c *Character) IsValueInRange(ident string, min, max float64) bool {
	e := c.GetElement(ident)
	if e == nil {
		return false
	}

	i := e.GetValueAsInt()
	if i >= int(min) && i <= int(max) {
		return true
	}

	return false
}

func (c *Character) IsValueInList(ident string, list string) bool {
	e := c.GetElement(ident)
	if e == nil {
		return false
	}

	l := strings.Split(list, ";")
	i := e.GetValueInfo(Id)
	v := e.GetValueInfo(Value)
	if slices.Contains(l, i) || slices.Contains(l, v) {
		return true
	}

	return false
}

func (c *Character) SetValueInt(ident string, value float64) {
	e := c.GetElement(ident)
	if e == nil {
		return
	}
	e.SetValue(int(value))

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
		if len(types) > 0 {
			e.SetValue(types[0].Label)
		} else {
			e.SetValue("")
		}
	case *Typevalue:
		ass.Validvalues = types

		for _, field := range types {
			i := e.GetValueInfo(Id)
			if field.Id == i {
				return
			}
		}
		if len(types) > 0 {
			e.SetValue(types[0])
		} else {
			e.SetValue(nil)
		}

	}
}

func (c *Character) SetDiceProperties(ident string, dicevalue, dicecount, dicemarkup float64) {
	e := c.GetElement(ident)
	if e == nil {
		return
	}

	arr := []int{int(dicevalue), int(dicecount), int(dicemarkup)}
	e.SetValue(arr)
}

func (c *Character) SetSkillProperties(ident string, dicevalue, dicecount, dicemarkup, dicebonusmarkup float64) {
	e := c.GetElement(ident)
	if e == nil {
		return
	}

	arr := []int{int(dicevalue), int(dicecount), int(dicemarkup), int(dicebonusmarkup)}
	e.SetValue(arr)
}
func (c *Character) IsAllValid() bool {
	for _, e := range c.Properties {
		if !e.isValidated {
			return false
		}
	}
	return true
}

func (c *Character) ClassifyProperties() {
	c.PropsGrouped = []*elementgroup{}
	if c.Properties != nil {
		// first get all fieldtypes with type == "property"
		// populate map and create empty slices
		for i := range c.Allfieldtypes {
			if c.Allfieldtypes[i].Type == "property" {
				g := &elementgroup{
					Group:    *c.Allfieldtypes[i],
					Elements: []*Element{},
				}
				c.PropsGrouped = append(c.PropsGrouped, g)
			}
		}

		// loop through characterproperties and add elements to map
		for _, p := range c.Properties {
			for _, g := range c.PropsGrouped {
				if g.Group.Id == p.Fieldtype.Type {
					g.Elements = append(g.Elements, p)
				}
			}
		}

	}
}
