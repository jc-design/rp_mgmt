package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestElementJsonUnMarshal(t *testing.T) {
	jsonStr := `{
    "Type": {
      "Type": "baseproperty",
      "Id": "race",
      "Label": "Rasse",
      "Description": "Rasse"
    },
    "ReferenceType": {
      "Type": "",
      "Id": "",
      "Label": "",
      "Description": ""
    },
    "Value": {
      "Type": "race",
      "Id": "hu",
      "Label": "Mensch",
      "Description": "Mensch"
    },
    "Visibility": "Creation|Levelup"
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
		Value:      "Test",
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
