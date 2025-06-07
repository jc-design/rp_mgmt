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

	s1, _ := val.GetInfo(models.Value)
	s2, _ := compare_val.GetInfo(models.Value)
	assert.Equal(t, s1, s2)
}
