package views

import (
	"fmt"
	"image/color"

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

	selectedid int

	elechan   chan *models.Character
	charmodel *CharacterModel

	OnAdded    func(i int)
	OnSelected func(i int)
}

func NewCharacterlist(c chan *models.Character, m *CharacterModel) *Characterlist {
	cl := &Characterlist{}
	cl.elechan = c
	cl.charmodel = m

	th := cl.Theme()

	cl.addbtn = &widget.Button{
		Icon:       th.Icon(theme.IconNameContentAdd),
		Text:       "Neuer Character",
		Importance: widget.HighImportance,
		OnTapped: func() {
			cl.charmodel.NewCharacter()
			// cl.charmodel.SelectedCharacter = cl.charmodel.Characters[len(cl.charmodel.Characters)-1]
			cl.characterlist.Select(len(cl.charmodel.Characters) - 1)
			c <- m.SelectedCharacter
			if cl.OnAdded != nil {
				cl.OnAdded(cl.characterlist.Length() - 1)
			}
			cl.Refresh()
		},
	}

	cl.characterlist = &widget.List{
		Length: func() int {
			return len(cl.charmodel.Characters)
		},
		CreateItem: func() fyne.CanvasObject {
			name := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop.TextSize = name.TextSize * 0.8
			return container.NewVBox(name, prop)
		},
		UpdateItem: func(i widget.ListItemID, o fyne.CanvasObject) {
			vbox := o.(*fyne.Container)
			vbox.Objects[0].(*canvas.Text).Text = cl.charmodel.Characters[i].Name
			vbox.Objects[1].(*canvas.Text).Text = fmt.Sprintf("Grad: %v - EP: %v",
				cl.charmodel.Characters[i].Level,
				cl.charmodel.Characters[i].Exp,
			)
		},
		OnSelected: func(id widget.ListItemID) {
			cl.selectedid = id
			cl.charmodel.SelectedCharacter = m.Characters[id]
			if cl.OnSelected != nil {
				cl.OnSelected(id)
			}
			cl.Refresh()
		},
	}

	cl.ExtendBaseWidget(cl)
	return cl
}

func (cl *Characterlist) CreateRenderer() fyne.WidgetRenderer {
	cl.ExtendBaseWidget(cl)
	x := canvas.NewRectangle(color.Black)
	x.Move(fyne.NewPos(100, 100))

	cl.layoutcont = container.New(
		&characterlistlayout{
			btn:  cl.addbtn,
			list: cl.characterlist,
		},
		cl.addbtn,
		cl.characterlist,
	)
	renderer := widget.NewSimpleRenderer(cl.layoutcont)
	cl.Refresh()
	return renderer
}

func (cl *Characterlist) Refresh() {
	cl.ExtendBaseWidget(cl)
	// if len(cl.controller.Model.Characters) == 0 {
	// 	if !cl.nocontentpassed {
	// 		cl.contentcont.Objects = []fyne.CanvasObject{&canvas.Text{
	// 			Text:      "Kein Charackter ausgewählt",
	// 			Color:     theme.Color(theme.ColorNameForeground),
	// 			Alignment: fyne.TextAlignCenter,
	// 			TextSize:  24,
	// 		}}
	// 		cl.nocontentpassed = true
	// 		cl.contentpassed = false
	// 	}
	// } else {
	// 	if !cl.contentpassed {
	// 		cb := NewControlbar(cl.controller)
	// 		cb.name.OnChanged = func(s string) {
	// 			cl.controller.Model.SelectedCharacter.Name = s
	// 			cl.characterlist.RefreshItem(cl.selectedid)
	// 		}
	// 		cb.exp.OnChanged = func(s string) {
	// 			if val, err := strconv.ParseInt(s, 10, 32); err == nil {
	// 				cl.controller.Model.SelectedCharacter.Exp = int(val)
	// 				cl.characterlist.RefreshItem(cl.selectedid)
	// 			}
	// 		}
	// 		cb.btnsave.OnTapped = func() {
	// 			cl.controller.Model.SaveCharacters()
	// 		}

	// 		cb.btndelete.OnTapped = func() {
	// 			err := cb.controller.Model.RemoveCharacter(cl.selectedid)
	// 			if err == nil {
	// 				cl.controller.Log("error removing charachter @CharacterlistWidget|btndelete.OnTapped", err)
	// 			}
	// 			if cl.characterlist.Length() > 0 {
	// 				cl.characterlist.Select(0)
	// 			}
	// 			cl.Refresh()
	// 		}
	// 		cl.contentcont.Objects = []fyne.CanvasObject{
	// 			container.NewVBox(cb,
	// 				NewDiceItem(cl.controller, cl.controller.Model.SelectedCharacter.GetElement("baseproperty|st")),
	// 				NewDiceItem(cl.controller, cl.controller.Model.SelectedCharacter.GetElement("baseproperty|gw")),
	// 			)}
	// 		cl.contentpassed = true
	// 		cl.nocontentpassed = false
	// 	}
	// }
	cl.characterlist.Refresh()
	canvas.Refresh(cl)
}

func (cl *Characterlist) MinSize() fyne.Size {
	if cl.layoutcont == nil {
		return fyne.Size{
			Width:  1280,
			Height: 720,
		}
	} else {
		return cl.layoutcont.MinSize()
	}
}
