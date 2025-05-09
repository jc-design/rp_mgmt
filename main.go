package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jc-design/rp_mgmt/internal/models"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/jc-design/rp_mgmt/internal/views"
)

const (
	APPNAME     string = "RoleplayManagement"
	description string = "description"
	id          string = "id"
	identify    string = "identify"
	value       string = "value"
)

var iconData []byte

func main() {

	appfolder, err := rules.NewFolderstructure(APPNAME)
	if err != nil {
		fmt.Printf("error while initialisation: %s\n", err)
		os.Exit(1)
	}

	charactermodel, err := views.NewCharacterModel(appfolder)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		switch input {
		case "exit":
			goto Exit
		case "new":
			charactermodel.NewCharacter()
			fmt.Println(charactermodel.SelectedCharacter.Name)
		case "dice":
			charactermodel.SelectedCharacter.GetElement("baseproperty|st").Execute()
			fmt.Println(charactermodel.SelectedCharacter.GetElement("baseproperty|st").Value.GetInfo(value))
		case "race":
			r := charactermodel.SelectedCharacter.GetElement("baseproperty|race").Value.(*models.Typevalue).Validvalues[1]
			charactermodel.SelectedCharacter.GetElement("baseproperty|race").SetValue(r)
			fmt.Println(charactermodel.SelectedCharacter.GetElement("baseproperty|class").Value.GetInfo(id))
			fmt.Println(charactermodel.SelectedCharacter.GetElement("baseproperty|classtype").Value.GetInfo(value))
			charactermodel.ApplyCreationRules()
			val := charactermodel.SelectedCharacter.GetElement("baseproperty|gs")
			fmt.Println(val)
		}
	}

Exit:

	// client, err := rules.NewInputOnlyRuleService(r, "creation", "v1")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if err := client.ApplyRules(c); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(c)
	// }

	// myApp := app.New()
	// w := myApp.NewWindow("Two Way")

	// str := binding.NewString()
	// str.Set("Hi!")

	// w.SetContent(container.NewVBox(
	// 	widget.NewLabelWithData(str),
	// 	widget.NewEntryWithData(str),
	// ))

	// w.ShowAndRun()
}
