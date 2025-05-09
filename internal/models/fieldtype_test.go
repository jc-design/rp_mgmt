package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSingleFieldtypeJsonUnMarshal(t *testing.T) {
	jsonStr := `{
      "type": "baseproperty",
      "id": "race",
      "label": "Rasse",
      "description": "Rasse"
    }`

	val := models.Fieldtype{}
	err := json.Unmarshal([]byte(jsonStr), &val)

	assert.NoError(t, err)
}

func TestFieldtypeJsonMarshal(t *testing.T) {
	val := models.Fieldtype{"fieldtype", "id", "label", "desciption"}
	marshalval, err := json.Marshal(&val)
	if err != nil {
		assert.Fail(t, "failing marshaling element")
	}

	compare_val := models.Fieldtype{}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling element")
	}
	assert.Equal(t, val, compare_val)
}
