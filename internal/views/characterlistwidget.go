package views

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
)

var _ fyne.Widget = (*Characterlist)(nil)

type Characterlist struct {
	widget.BaseWidget

	layoutcont    *fyne.Container
	addbtn        *widget.Button
	characterlist *widget.List
	contentcont   *fyne.Container

	selectedid      int
	nocontentpassed bool
	contentpassed   bool

	controller   *CharacterController
	activechars  []*models.Character
	deletedchars []*models.Character

	addcharacter func()
}

func NewCharacterlist(ctrl *CharacterController) *Characterlist {
	cl := &Characterlist{}
	cl.controller = ctrl

	cl.addcharacter = func() {
		ctrl.Model.NewCharacter()
		ctrl.Model.SelectedCharacter = ctrl.Model.Characters[len(ctrl.Model.Characters)-1]
		ctrl.Model.ApplyCreationRules()
		cl.characterlist.Select(len(ctrl.Model.Characters) - 1)
		cl.layoutcont.Refresh()
	}

	cl.ExtendBaseWidget(cl)
	return cl
}

func (cl *Characterlist) CreateRenderer() fyne.WidgetRenderer {
	cl.ExtendBaseWidget(cl)

	th := cl.Theme()

	cl.addbtn = &widget.Button{
		Icon:       th.Icon(theme.IconNameContentAdd),
		Text:       "Neuer Character",
		Importance: widget.HighImportance,
		OnTapped:   cl.addcharacter,
	}

	cl.characterlist = &widget.List{
		Length: func() int {
			return len(cl.controller.Model.Characters)
		},
		CreateItem: func() fyne.CanvasObject {
			name := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop.TextSize = name.TextSize * 0.8

			return container.NewVBox(name, prop)
		},
		UpdateItem: func(i widget.ListItemID, o fyne.CanvasObject) {
			vbox := o.(*fyne.Container)
			vbox.Objects[0].(*canvas.Text).Text = cl.controller.Model.Characters[i].Name
			vbox.Objects[1].(*canvas.Text).Text = fmt.Sprintf("Grad: %v - EP: %v",
				cl.controller.Model.Characters[i].Level,
				cl.controller.Model.Characters[i].Exp,
			)
		},
		OnSelected: func(id widget.ListItemID) {
			cl.selectedid = id
			cl.controller.Model.SelectedCharacter = cl.controller.Model.Characters[id]
			cl.Refresh()
		},
	}
	cl.controller.addbindings("selectedchar", cl.characterlist)

	cl.contentcont = container.NewStack()
	cl.layoutcont = container.New(
		&characterlistlayout{
			btn:     cl.addbtn,
			list:    cl.characterlist,
			content: cl.contentcont,
		},
		cl.addbtn,
		cl.characterlist,
		cl.contentcont,
	)
	renderer := widget.NewSimpleRenderer(cl.layoutcont)
	cl.Refresh()
	return renderer
}

func (cl *Characterlist) Refresh() {
	cl.ExtendBaseWidget(cl)

	if len(cl.controller.Model.Characters) == 0 {
		if !cl.nocontentpassed {
			cl.contentcont.Objects = []fyne.CanvasObject{&canvas.Text{
				Text:      "Kein Charackter ausgewählt",
				Color:     theme.Color(theme.ColorNameForeground),
				Alignment: fyne.TextAlignCenter,
				TextSize:  24,
			}}
			cl.nocontentpassed = true
		}
	} else {
		if !cl.contentpassed {
			cl.contentcont.Objects = []fyne.CanvasObject{&canvas.Text{
				Text:     cl.controller.Model.Characters[cl.selectedid].Name,
				Color:    theme.Color(theme.ColorNameForeground),
				TextSize: 24,
			}}
			cl.contentpassed = true
		}
	}
	cl.layoutcont.Refresh()
}
