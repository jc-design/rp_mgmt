package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestElementJsonUnMarshal(t *testing.T) {
	jsonStr := `{
    "type": {
      "type": "baseproperty",
      "id": "race",
      "label": "Rasse",
      "description": "Rasse"
    },
    "referencetype": {
      "type": "",
      "id": "",
      "label": "",
      "description": ""
    },
    "value": {
      "type": "race",
      "id": "hu",
      "label": "Mensch",
      "description": "Mensch"
    },
    "visibility": "creation|levelup"
  }`

	e := models.Element{}
	err := json.Unmarshal([]byte(jsonStr), &e)

	assert.NoError(t, err)
}

func TestElementJsonMarshal(t *testing.T) {
	e := models.Element{
		Type: models.FieldType{
			Type:        "Test",
			Id:          "Test",
			Label:       "Test",
			Description: "Test",
		},
		ReferenceType: models.FieldType{
			Type:        "Test",
			Id:          "Test",
			Label:       "Test",
			Description: "Test",
		},
		Value:      &models.StringValue{StringValue: "Test"},
		Visibility: 5,
	}
	j, err := json.Marshal(&e)
	if err != nil {
		assert.Fail(t, "failing marshaling Element")
	}

	compare_e := models.Element{}
	err = json.Unmarshal(j, &compare_e)
	if err != nil {
		assert.Fail(t, "failing unmarshaling Element")
	}
	assert.Equal(t, e, compare_e)
}

func TestElementExecuteDiceFunction(t *testing.T) {
	jsonStr := `{
    "type": {
      "type": "baseproperty",
      "id": "race",
      "label": "Rasse",
      "description": "Rasse"
    },
    "referencetype": {
      "type": "",
      "id": "",
      "label": "",
      "description": ""
    },
    "value": {
      "dicevalue": 100,
      "dicecount": 1,
      "dicemarkup": 0,
      "value": 0,
	  "abr": "W"
    },
    "visibility": "creation|levelup"
  }`

	e := models.Element{}
	err := json.Unmarshal([]byte(jsonStr), &e)
	assert.NoError(t, err)

	e.Value.Execute()
	assert.NotEqual(t, "0", e.Value.ValueAsString())
}
