package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

type CharacterController struct {
	Model  *CharacterModel
	App    fyne.App
	Window fyne.Window

	Elechan chan *models.Character

	// TODO:
	// check if mutex and bindings are still necessary
	Log func(string, error) `json:"-"`
}

type CharacterModel struct {
	allTypes            []*models.Fieldtype
	allElements         []*models.Element
	ruleset             rules.Ruleset
	folderstr           *rules.Folderstructure
	creationRuleservice rules.RulesApplier

	Characters        []*models.Character
	SelectedCharacter *models.Character
	Log               func(string, error) `json:"-"`
}

// func needed for RuleFact interface
// it's needed for the grule-engine validation
// not used for the moment
func (c *CharacterModel) FactKey() string {
	return "Model"
}

// func needed for RuleFact interface
// it's needed for the grule-engine validation
// not used for the moment
func (c *CharacterController) FactKey() string {
	return "Controller"
}

func NewCharacterController(f *rules.Folderstructure, log func(string, error), windowtitel string) (*CharacterController, error) {
	ctrl := CharacterController{}

	// TODO
	// resized window, check minsize func
	ctrl.App = app.New()
	ctrl.Window = ctrl.App.NewWindow(windowtitel)
	ctrl.Window.Resize(fyne.NewSize(640, 480))
	ctrl.Window.CenterOnScreen()

	// TODO
	// create own custom theme
	ctrl.App.Settings().SetTheme(theme.LightTheme())

	ctrl.Elechan = make(chan *models.Character)
	go func() {
		for {
			select {
			case c := <-ctrl.Elechan:
				ctrl.Model.creationRuleservice.ApplyRules(c)
				fyne.Do(func() {
					// refresh complete window content
					// check if savebutton is activated as well
					ctrl.Window.Content().Refresh()
				})
			}
		}
	}()

	ctrl.Log = log

	// error will be logged in main.go
	var err error
	ctrl.Model, err = NewCharacterModel(f, log)
	if err != nil {
		return &ctrl, fmt.Errorf("error while creating new charactermodel: %w\n", err)
	}

	return &ctrl, nil
}

func NewCharacterModel(f *rules.Folderstructure, log func(string, error)) (*CharacterModel, error) {
	var err error
	cm := CharacterModel{}

	//first load active ruleset
	//the ruleset is need for validation
	cm.ruleset, err = rules.LoadRuleSet(f)
	if err != nil {
		return nil, err
	}

	cm.folderstr = f
	cm.Log = log

	//prepare waitgroup and channel for following go-routines
	wg := &sync.WaitGroup{}
	respch := make(chan error, 3)

	//add counter for 3 go-routines
	wg.Add(3)

	//spin up go-routine
	//load type definitions
	go func(cm *CharacterModel, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		err := cm.LoadTypes()
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, respch, wg)

	//spin up go-routine
	//load element definitions
	go func(cm *CharacterModel, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		err := cm.LoadElements()
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, respch, wg)

	//spin up go-routine
	//create ruleservice with creation rules
	go func(cm *CharacterModel, respch chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		cm.creationRuleservice, err = cm.newRuleservice("createcharacter", "creation", "version")
		if err != nil {
			respch <- err
			return
		}
		respch <- nil
	}(&cm, respch, wg)

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

	err = cm.LoadCharacters()
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			cm.Log("error loading characters", err)
		}
	}

	if len(cm.Characters) > 0 {
		cm.SelectedCharacter = cm.Characters[0]
	}
	return &cm, nil
}

func (cm *CharacterModel) LoadTypes() error {

	data, err := os.ReadFile(filepath.Join(cm.folderstr.Data, "types.json"))
	if err != nil {
		return fmt.Errorf("error loading types:  %w\n", err)
	}

	err = json.Unmarshal(data, &cm.allTypes)
	if err != nil {
		return fmt.Errorf("error unmarshaling types:  %w\n", err)
	}

	return nil
}

func (cm *CharacterModel) LoadElements() error {
	data, err := os.ReadFile(filepath.Join(cm.folderstr.Data, "characterproperties.json"))
	if err != nil {
		return fmt.Errorf("error loading elements:  %w\n", err)
	}

	err = json.Unmarshal(data, &cm.allElements)
	if err != nil {
		return fmt.Errorf("error unmarshaling elements:  %w\n", err)
	}

	// pass log func to Element, so errors could be logged where they occur
	for _, e := range cm.allElements {
		e.Log = cm.Log
	}

	return nil
}

func (cm *CharacterModel) LoadCharacters() error {

	data, err := os.ReadFile(filepath.Join(cm.folderstr.Characters, "characters.json"))
	if err != nil {
		return fmt.Errorf("error loading characters:  %w\n", err)
	}

	err = json.Unmarshal(data, &cm.Characters)
	if err != nil {
		return fmt.Errorf("error unmarshaling characters:  %w\n", err)
	}

	// check if character ruleset is equal to defined ruleset
	for _, c := range cm.Characters {
		if c.RuleSet != cm.ruleset {
			return fmt.Errorf("error loading characters due to incomapatible ruleset. "+
				"Loaded ruleset: %s, requested ruleset: %s", c.RuleSet, cm.ruleset)
		}
		c.Allfieldtypes = cm.allTypes
		c.Status = models.Activationmode(models.Levelup)
		c.ClassifyProperties()

		// pass log func to Element, so errors could be logged where they occur
		for _, e := range c.Properties {
			e.Log = cm.Log
		}
	}
	return nil
}

func (cm *CharacterModel) SaveCharacters() error {

	bytes, err := json.MarshalIndent(cm.Characters, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling charatcers:  %w\n", err)

	}

	if err := os.WriteFile(filepath.Join(cm.folderstr.Characters, "characters.json"), bytes, 0644); err != nil {
		return fmt.Errorf("error saving characters:  %w\n", err)
	}

	return nil
}

func (cm *CharacterModel) NewCharacter() {
	c := models.NewCharacter(cm.ruleset, cm.allElements, cm.allTypes, cm.Log)
	cm.Characters = append(cm.Characters, c)

	cm.SelectedCharacter = c
	for _, e := range c.Properties {
		e.RulesReset()
	}

	// TODO
	// Access CharacterController
	cm.ApplyCreationRules()
}

func (cm *CharacterModel) RemoveCharacter(index int) error {
	if index < 0 || index > len(cm.Characters) {
		return fmt.Errorf("index is invalid")
	}
	cm.Characters = slices.Delete(cm.Characters, index, index+1)
	return nil
}

func (cm *CharacterModel) ApplyCreationRules() {
	cm.creationRuleservice.ApplyRules(cm.SelectedCharacter)
}

func (cm *CharacterModel) newRuleservice(rulefn, name, version string) (rules.RulesApplier, error) {

	entries, err := os.ReadDir(cm.folderstr.Rules)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	for _, v := range entries {
		if v.IsDir() {
			continue
		}
		if filepath.Ext(v.Name()) != ".grl" {
			continue
		}
		if !strings.HasPrefix(v.Name(), rulefn) {
			continue
		}

		file_data, err := os.ReadFile(filepath.Join(cm.folderstr.Rules, v.Name())) // For read access.
		if err != nil {
			return nil, err
		}
		data = append(data, file_data...)
	}

	client, err := rules.NewInputOnlyRuleService(data, name, version)
	if err != nil {
		return nil, err
	}
	return client, nil
}
