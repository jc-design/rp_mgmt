package models

import "github.com/jc-design/rp_mgmt/internal/rules"

type CharacterRuleService interface {
	ApplyRules(c *Character) error
}

type CharacterContext struct {
	RuleInputFact *Character
}

type CharacterRuleServiceClient struct {
	ruleEngineSvc *rules.RuleEngineService
}

func (c *CharacterContext) RuleFacts() []rules.RuleFact {
	var f = make([]rules.RuleFact, 0)
	f = append(f, c.RuleInputFact)
	return f
}

func NewCharacterRuleService(rulesAsBytes []byte, name string, version string) (CharacterRuleService, error) {
	ruleEngineSvc, err := rules.NewRuleEngineSvc(rulesAsBytes, name, version)
	if err != nil {
		return nil, err
	}

	return &CharacterRuleServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}, nil
}

func newCharacterContext(c *Character) *CharacterContext {
	return &CharacterContext{
		RuleInputFact: c,
	}
}

func (svc CharacterRuleServiceClient) ApplyRules(c *Character) error {
	context := newCharacterContext(c)

	err := svc.ruleEngineSvc.Execute(context)
	if err != nil {
		return err
	}

	return nil
}
