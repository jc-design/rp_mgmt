package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

type CharacterModel struct {
	allTypes            []*models.Fieldtype
	allElements         []*models.Element
	ruleset             rules.Ruleset
	Characters          []*models.Character
	SelectedCharacter   *models.Character
	creationRuleservice rules.RulesApplier
}

type CharacterView struct {
	widget   *widget.Entry
	onSubmit func()
}

type CharacterController struct {
	Model *CharacterModel
	View  *CharacterView
}

// func needed for RuleFact interface
// it's needed for the grule-engine validation
func (c *CharacterModel) FactKey() string {
	return "Model"
}

// func needed for RuleFact interface
// it's needed for the grule-engine validation
func (c *CharacterController) FactKey() string {
	return "Controller"
}

func NewCharacterModel(f rules.Folderstructure) (*CharacterModel, error) {
	var err error
	cm := CharacterModel{}

	//first load active ruleset
	//the ruleset is need for validation
	cm.ruleset, err = rules.LoadRuleSet(f)
	if err != nil {
		return nil, err
	}

	//prepare waitgroup and channel for following go-routines
	wg := &sync.WaitGroup{}
	respch := make(chan error, 3)

	//add counter for 3 go-routines
	wg.Add(3)

	//spin up go-routine
	//load type definitions
	go func(cm *CharacterModel, f rules.Folderstructure, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		err := cm.LoadTypes(f)
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, f, respch, wg)

	//spin up go-routine
	//load element definitions
	go func(cm *CharacterModel, f rules.Folderstructure, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		err := cm.LoadElements(f)
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, f, respch, wg)

	//spin up go-routine
	//create ruleservice with creation rules
	go func(cm *CharacterModel, f rules.Folderstructure, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		cm.creationRuleservice, err = cm.newRuleservice(f, "createcharacter.grl", "creation", "version")
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, f, respch, wg)

	wg.Wait()
	close(respch)
	var errs []error
	for res := range respch {
		if res != nil {
			errs = append(errs, res)
		}
	}
	if len(errs) > 0 {
		// Join returns a single `error`.
		// Underlying, the error contains all the errors we add.
		return nil, errors.Join(errs...)
	}

	err = cm.LoadCharacters(f)
	if err != nil {
		fmt.Printf("error while loaded characters.json: %v\n", err)
	}

	if len(cm.Characters) > 0 {
		cm.SelectedCharacter = cm.Characters[0]
	}
	return &cm, nil
}

func (cm *CharacterModel) LoadTypes(f rules.Folderstructure) error {

	data, err := os.ReadFile(filepath.Join(f.Data, "types.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cm.allTypes)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CharacterModel) LoadElements(f rules.Folderstructure) error {
	data, err := os.ReadFile(filepath.Join(f.Data, "characterproperties.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cm.allElements)
	if err != nil {
		return err
	}

	return nil
}

func (cm *CharacterModel) LoadCharacters(f rules.Folderstructure) error {

	data, err := os.ReadFile(filepath.Join(f.Characters, "characters.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cm.Characters)
	if err != nil {
		return err
	}

	for _, c := range cm.Characters {
		if c.RuleSet != cm.ruleset {
			return fmt.Errorf("error loading characters due to incomapatible ruleset. "+
				"Loaded ruleset: %s, requested ruleset: %s", c.RuleSet, cm.ruleset)
		}
		c.Allfieldtypes = cm.allTypes
		c.Status = models.Activationmode(models.Levelup)
	}
	return nil
}

func (cm *CharacterModel) SaveCharacters(f rules.Folderstructure) error {

	// Create slice with length of 0 from the given slice
	tmp := cm.Characters[:0]
	for _, c := range cm.Characters {
		// if character is not marked as [Deleted], then copy it to slice
		for _, e := range c.Properties {
			if !e.GetValidation() {
				continue
			}
		}
		if !c.Deleted {
			tmp = append(tmp, c)
		}
	}

	bytes, err := json.MarshalIndent(tmp, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(f.Characters, "characters.json"), bytes, 0644); err != nil {
		return err
	}

	return nil
}

func (cm *CharacterModel) NewCharacter() {
	c := models.NewCharacter(cm.ruleset, cm.allElements, cm.allTypes)
	cm.Characters = append(cm.Characters, c)

	cm.SelectedCharacter = c
	for _, e := range c.Properties {
		e.RulesReset()
	}
	cm.ApplyCreationRules()
}

func (cm *CharacterModel) ApplyCreationRules() {
	cm.creationRuleservice.ApplyRules(cm.SelectedCharacter)
}

func (cm *CharacterModel) newRuleservice(f rules.Folderstructure, rulefn, name, version string) (rules.RulesApplier, error) {
	data, err := os.ReadFile(filepath.Join(f.Rules, rulefn))
	if err != nil {
		return nil, err
	}
	client, err := rules.NewInputOnlyRuleService(data, name, version)
	if err != nil {
		return nil, err
	}
	return client, nil
}
