package rules

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var knowledgeLibrary = *ast.NewKnowledgeLibrary()

// Rule fact object
type RuleFact interface {
	FactKey() string
}

// Configs associated with each rule
type RuleConfig interface {
	RuleFacts() []RuleFact
}

type RuleEngineService struct {
	name    string
	version string
}

func NewRuleEngineSvc(rulesAsBytes []byte, name string, version string) (*RuleEngineService, error) {

	if err := buildRuleEngine(rulesAsBytes, name, version); err != nil {
		return nil, err
	}
	return &RuleEngineService{
		name:    name,
		version: version,
	}, nil
}

func buildRuleEngine(rulesAsBytes []byte, name string, version string) error {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibrary)

	// Read rule from file and build rules
	ruleFile := pkg.NewBytesResource(rulesAsBytes)
	err := ruleBuilder.BuildRuleFromResource(name, version, ruleFile)
	if err != nil {
		return err
	}

	return nil
}

func (svc *RuleEngineService) Execute(ruleConf RuleConfig) error {
	// Get KnowledgeBase instance to execute particular rule
	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance(svc.name, svc.version)
	if err != nil {
		return err
	}

	dataCtx := ast.NewDataContext()

	for _, f := range ruleConf.RuleFacts() {
		// Add fact data context
		err = dataCtx.Add(f.FactKey(), f)
		if err != nil {
			return err
		}
	}

	// Create rule engine and execute on provided data and knowledge base
	ruleEngine := engine.NewGruleEngine()
	err = ruleEngine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return err
	}
	return nil
}
