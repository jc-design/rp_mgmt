package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/jc-design/rp_mgmt/internal/views"
)

const (
	APPNAME string = "RoleplayManagement"
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
		if ctrl == nil {
			os.Exit(1)
		}
		l := container.NewScroll(widget.NewRichTextFromMarkdown(
			fmt.Sprintf("# Es is leider ein Fehler aufgetreten:\n"+
				"⚠️ `%s`\n\n"+
				"#\n"+
				"## HowTo\n"+
				"Schau dir bitte die Anleitung und Beschreibung auf [github/jc-design/rp_mgmt](https://github.com/jc-design/rp_mgmt/) an.\n\n"+
				"Dort wird beschrieben, welche Schritte unternommen und Dateien erzeugt werden müssen.\n\n"+
				"Viel Spaß!", err),
		))
		ctrl.Window.SetContent(l)

	} else {
		l := container.NewStack(views.NewCharacterlist(ctrl))
		ctrl.Window.SetContent(l)
	}

	// ctrl.Model.NewCharacter()
	// ctrl.Model.SelectedCharacter = ctrl.Model.Characters[0]
	// ctrl.Model.ApplyCreationRules()

	// e1 := ctrl.Model.SelectedCharacter.GetElement("baseproperty|race")
	// e2 := ctrl.Model.SelectedCharacter.GetElement("baseproperty|class")

	// hb := container.NewVBox(
	// 	views.NewTypevalueItem(ctrl, e1),
	// 	views.NewTypevalueItem(ctrl, e2),
	// )
	// e := ctrl.Model.SelectedCharacter.GetElement("baseproperty|race")
	// item := views.NewDiceItem(ctrl, e)
	// item := views.NewTypevalueItem(ctrl, e)

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
