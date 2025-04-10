package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const APPNAME = "RoleplayManagement"

var iconData []byte

func main() {

	// appfolder, err := rules.NewFolderstructure(APPNAME)
	// if err != nil {
	// 	fmt.Printf("error while initialisation: %s\n", err)
	// 	os.Exit(1)
	// }

	// ruleset, err := rules.LoadRuleSet(appfolder)
	// if err != nil {
	// 	fmt.Printf("error while loading ruleset: %s\n", err)
	// }

	// _, err = models.LoadTypes(appfolder)
	// if err != nil {
	// 	fmt.Printf("error while loading standardtypes: %s\n", err)
	// }

	// r, err := rules.LoadRules(appfolder, "creationrules.grl")
	// if err != nil {
	// 	fmt.Printf("error while loading rules: %s\n", err)
	// }
	// c := models.NewCharacter(ruleset, []models.Element{})

	// client, err := models.NewCharacterRuleService(r, "creation", "v1")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// if err := client.ApplyRules(&c); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(c)
	// }

	myApp := app.New()
	w := myApp.NewWindow("Two Way")

	str := binding.NewString()
	str.Set("Hi!")

	w.SetContent(container.NewVBox(
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	))

	w.ShowAndRun()
}
