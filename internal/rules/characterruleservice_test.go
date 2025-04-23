package rules_test

import (
	"testing"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/stretchr/testify/assert"
)

func TestExecuteRule(t *testing.T) {
	rule := []byte(`rule TestRule {
  when 
    Character.Name == "New Character"
  then
    Character.Name = "12345";
    Retract("TestRule");
}`)

	ruleset := rules.RuleSet{
		Name:    "Midgard",
		Version: "M5",
	}
	c := models.NewCharacter(ruleset, []models.Element{})

	assert.Equal(t, c.Name, "New Character")

	client, err := rules.NewInputOnlyRuleService(rule, "creation", "v1")
	assert.NoError(t, err)

	err = client.ApplyRules(c)
	assert.NoError(t, err)

	assert.Equal(t, "12345", c.Name)
}
