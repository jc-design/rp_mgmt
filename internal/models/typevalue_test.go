package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestTypevalueJsonUnMarshal(t *testing.T) {
	jsonStr := `{
      "type": "baseproperty",
      "id": "race",
      "label": "Rasse",
      "description": "Rasse"
	}`

	val := models.Typevalue{}
	err := json.Unmarshal([]byte(jsonStr), &val)

	assert.NoError(t, err)
}

func TestTypevalueJsonMarshal(t *testing.T) {
	ft := models.Fieldtype{"fieldtype", "id", "label", "desciption"}
	fts := []models.Fieldtype{
		ft,
		ft,
		ft,
	}

	val := models.Typevalue{}
	val.SetValue(ft, fts)

	marshalval, err := json.Marshal(&val)
	if err != nil {
		assert.Fail(t, "failing marshaling element")
	}

	compare_val := models.Typevalue{}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling element")
	}

	assert.Equal(t, val.GetInfo("identify"), compare_val.GetInfo("identify"))
}
