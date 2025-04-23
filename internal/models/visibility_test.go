package models_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestVisbilityJsonUnMarshal(t *testing.T) {
	jsonStr := `{"Visibility":"creation|levelup"}`

	e := struct {
		Visibility models.Visibility
	}{}
	err := json.Unmarshal([]byte(jsonStr), &e)
	if err != nil {
		assert.Fail(t, "failing unmarshaling Visibility")
	}
	fmt.Println(e)
	assert.Equal(t, 3, int(e.Visibility))
}

func TestVisbilityJsonMarshal(t *testing.T) {
	jsonStr := `{"Visibility":"creation|levelup"}`

	e := struct {
		Visibility models.Visibility
	}{Visibility: 3}
	json, err := json.Marshal(&e)
	if err != nil {
		assert.Fail(t, "failing marshaling Visibility")
	}
	fmt.Println(string(json))
	assert.Equal(t, jsonStr, string(json))
}
