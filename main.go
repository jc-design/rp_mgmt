package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/container"
	_ "fyne.io/fyne/v2/data/binding"
	_ "fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/rules"
)

const APPNAME = "RoleplayManagement"

var iconData []byte

func main() {

	myApp := app.NewWithID("ddd")
	x := myApp.Storage().RootURI
	fmt.Println(x().Path())

	appfolder, err := rules.NewFolderstructure(APPNAME)
	if err != nil {
		fmt.Printf("error while initialisation: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(appfolder.Basepath)

	// charactermodel, err := views.NewCharacterModel(appfolder)

	// fmt.Println(charactermodel)
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
