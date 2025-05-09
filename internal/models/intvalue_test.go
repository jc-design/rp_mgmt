package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestIntvalueJsonUnMarshal(t *testing.T) {
	jsonStr := `{
      "intvalue": 10
    }`

	val := models.Intvalue{}
	err := json.Unmarshal([]byte(jsonStr), &val)

	assert.NoError(t, err)
}

func TestIntvalueJsonMarshal(t *testing.T) {
	val := models.Intvalue{Intvalue: 10}
	marshalval, err := json.Marshal(&val)
	if err != nil {
		assert.Fail(t, "failing marshaling element")
	}

	compare_val := models.Intvalue{Intvalue: 10}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling element")
	}
	assert.Equal(t, val.GetInfo("value"), compare_val.GetInfo("value"))
}
