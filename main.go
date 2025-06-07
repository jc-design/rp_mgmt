package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/rules"
	"github.com/jc-design/rp_mgmt/internal/views"
)

const (
	APPNAME string = "RoleplayManagement"
)

func main() {

	//os.Setenv("FYNE_SCALE", "1.5")
	// create folderstructer
	// ~/.config/RoleplayManagement
	// ~/.config/RoleplayManagement/characters
	// ~/.config/RoleplayManagement/data
	// ~/.config/RoleplayManagement/logfiles
	// ~/.config/RoleplayManagement/rules
	// ~/.config/RoleplayManagement/settings

	appfolder, err := rules.NewFolderstructure(APPNAME)
	if err != nil {
		log.Fatalf("error while initialisation: %s\n", err)
	}

	// create logfile ~/.config/RoleplayManagement/logfiles/log_yyyy.MM.dd.json
	// create a new file for each day
	l, err := os.OpenFile(filepath.Join(appfolder.Logfiles, "log_"+time.Now().Format("2006.01.02")+".json"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer l.Close()

	// create logger and log func
	// log func will be passed to different structs
	logger := slog.New(slog.NewJSONHandler(l, nil))
	log := func(msg string, err error) {
		logger.Error(msg, "error", err)
	}

	// create charactercontroller
	// check if anny error occurs duruing initialisation
	// show fyne-window
	ctrl, err := views.NewCharacterController(&appfolder, log, APPNAME)
	if err != nil {
		log("error creating new character controller", err)
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
		l := container.NewStack(views.NewW_Arrangement(ctrl.Elechan, ctrl.Model))
		ctrl.Window.SetContent((l))
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
