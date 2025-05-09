package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestCharacterJsonMarshal(t *testing.T) {
	elements := []*models.Element{
		{
			Fieldtype: models.Fieldtype{
				Type:        "baseproperty",
				Id:          "race",
				Label:       "Rasse",
				Description: "Rasse",
			},
			Value: &models.Typevalue{
				Fieldvalue: models.Fieldtype{
					Type:        "race",
					Id:          "hu",
					Label:       "Mensch",
					Description: "Mensch",
				},
			},
			Visibility: models.Activationmode(5),
		},
		{
			Fieldtype: models.Fieldtype{
				Type:        "baseproperty",
				Id:          "st",
				Label:       "Stärke",
				Description: "Stärke",
			},
			Value: &models.Dice{
				DiceValue:  100,
				DiceCount:  1,
				DiceMarkup: 0,
				Value:      0,
				Abr:        "W",
			},
			Visibility: models.Activationmode(5),
		},
		{
			Fieldtype: models.Fieldtype{
				Type:        "baseproperty",
				Id:          "age",
				Label:       "Alter",
				Description: "Alter",
			},
			Value: &models.Intvalue{
				Intvalue: 35,
			},
			Visibility: models.Activationmode(5),
		},
		{
			Fieldtype: models.Fieldtype{
				Type:        "baseproperty",
				Id:          "origin",
				Label:       "Herkunft",
				Description: "Herkunft",
			},
			Value: &models.Stringvalue{
				Stringvalue: "Alba",
			},
			Visibility: models.Activationmode(5),
		},
	}
	r := rules.Ruleset{
		Name:    "Midgard",
		Version: "M5",
	}
	c := models.Character{
		Id:         "id",
		Name:       "New Character",
		Image:      "",
		RuleSet:    r,
		Properties: elements,
	}
	marshalval, err := json.Marshal(&c)
	if err != nil {
		assert.Fail(t, "failing marshaling element")
	}

	compare_val := models.Character{}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling element")
	}

	assert.Equal(t, c, compare_val)
}
