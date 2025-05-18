package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/container"
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

func main() {

	// os.Setenv("FYNE_SCALE", "1.5")
	appfolder, err := rules.NewFolderstructure(APPNAME)
	if err != nil {
		fmt.Printf("error while initialisation: %s\n", err)
		os.Exit(1)
	}

	ctrl, err := views.NewCharacterController(appfolder, APPNAME)
	if err != nil {
		fmt.Printf("error while initialisation: %s\n", err)
		os.Exit(1)
	}

	ctrl.CreateCharacterList()
	ctrl.Model.NewCharacter()
	ctrl.Model.SelectedCharacter = ctrl.Model.Characters[0]
	ctrl.Model.ApplyCreationRules()

	e1 := ctrl.Model.SelectedCharacter.GetElement("baseproperty|race")
	e2 := ctrl.Model.SelectedCharacter.GetElement("baseproperty|class")

	hb := container.NewVBox(
		views.NewTypevalueItem(ctrl, e1),
		views.NewTypevalueItem(ctrl, e2),
	)
	// e := ctrl.Model.SelectedCharacter.GetElement("baseproperty|race")
	// item := views.NewDiceItem(ctrl, e)
	// item := views.NewTypevalueItem(ctrl, e)
	ctrl.Window.SetContent(hb)
	ctrl.Window.Content().Refresh()
	ctrl.Window.ShowAndRun()

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
