package views

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

type CharacterModel struct {
	allTypes          []models.FieldType
	allElements       []models.Element
	ruleset           rules.RuleSet
	Characters        []models.Character
	SelectedCharacter models.Character
}

type CharacterView struct {
	widget   *widget.Entry
	onSubmit func()
}

type CharacterController struct {
	Model *CharacterModel
	View  *CharacterView
}

// func neede for RuleFact interface
// it's neede for the grule-engine validation
func (c *CharacterModel) FactKey() string {
	return "Model"
}

// func neede for RuleFact interface
// it's neede for the grule-engine validation
func (c *CharacterController) FactKey() string {
	return "Controller"
}

func NewCharacterModel(f rules.Folderstructure) (*CharacterModel, error) {
	var err error
	cm := CharacterModel{}

	respch := make(chan bool, 3)
	wg := &sync.WaitGroup{}

	cm.ruleset, err = rules.LoadRuleSet(f)
	if err != nil {
		return nil, err
	}

	wg.Add(3)

	go loadtypes(&cm, f, respch, wg)
	go loadelements(&cm, f, respch, wg)
	go loadcharacters(&cm, f, respch, wg)

	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	if len(cm.Characters) > 0 {
		cm.SelectedCharacter = cm.Characters[0]
	}
	return &cm, nil
}

func loadruleset(cm *CharacterModel, f rules.Folderstructure, respch chan bool, wg *sync.WaitGroup) {
	ruleset, err := rules.LoadRuleSet(f)
	if err != nil {
		fmt.Printf("error while loading ruleset: %s\n", err)
		respch <- false
		wg.Done()
		return
	}
	cm.ruleset = ruleset
	respch <- true
	wg.Done()
}

func loadtypes(cm *CharacterModel, f rules.Folderstructure, respch chan bool, wg *sync.WaitGroup) {
	types, err := models.LoadTypes(f)
	if err != nil {
		fmt.Printf("error while loading fieldtypes: %s\n", err)
		respch <- false
		wg.Done()
		return
	}
	cm.allTypes = types
	respch <- true
	wg.Done()
}

func loadelements(cm *CharacterModel, f rules.Folderstructure, respch chan bool, wg *sync.WaitGroup) {
	elements, err := models.LoadElements(f)
	if err != nil {
		fmt.Printf("error while loading elements: %s\n", err)
		respch <- false
		wg.Done()
		return
	}
	cm.allElements = elements
	respch <- true
	wg.Done()
}

func loadcharacters(cm *CharacterModel, f rules.Folderstructure, respch chan bool, wg *sync.WaitGroup) {
	characters, err := models.LoadCharacters(f, cm.ruleset)
	if err != nil {
		fmt.Printf("error while loading characters: %s\n", err)
		respch <- false
		wg.Done()
		return
	}
	cm.Characters = characters
	respch <- true
	wg.Done()
}
