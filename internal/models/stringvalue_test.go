package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStringvalueJsonUnMarshal(t *testing.T) {
	jsonStr := `{
      "stringvalue": "test"
    }`

	val := models.Stringvalue{}
	err := json.Unmarshal([]byte(jsonStr), &val)

	assert.NoError(t, err)
}

func TestStringvalueJsonMarshal(t *testing.T) {
	val := models.Stringvalue{Stringvalue: "Test"}
	marshalval, err := json.Marshal(&val)
	if err != nil {
		assert.Fail(t, "failing marshaling element")
	}

	compare_val := models.Stringvalue{Stringvalue: ""}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling element")
	}
	assert.Equal(t, val.GetInfo("value"), compare_val.GetInfo("value"))
}
