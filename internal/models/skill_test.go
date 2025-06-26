package models_test

import (
	"encoding/json"
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSkillJsonMarshal(t *testing.T) {
	skill := models.Skill{

		Skilltype:       "innate",
		DiceValue:       20,
		DiceCount:       1,
		DiceMarkup:      -2,
		DiceBonusMarkup: 0,
	}
	marshalval, err := json.Marshal(&skill)
	if err != nil {
		assert.Fail(t, "failing marshaling skill")
	}

	compare_val := models.Skill{}
	err = json.Unmarshal(marshalval, &compare_val)
	if err != nil {
		assert.Fail(t, "failing unmarshaling skill")
	}

	assert.Equal(t, skill, compare_val)
}
