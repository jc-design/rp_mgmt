package views

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
)

var _ fyne.Widget = (*Controlbar)(nil)

type Controlbar struct {
	widget.BaseWidget

	layoutcont *fyne.Container
	gfx        *canvas.Image
	labelname  *canvas.Text
	name       *widget.Entry
	labelexp   *canvas.Text
	exp        *NumericEntry
	labellvl   *canvas.Text
	lvl        *canvas.Text
	btnsave    *widget.Button
	btndelete  *widget.Button

	cm *CharacterModel
}

func NewControlbar(c *CharacterModel) *Controlbar {
	cb := &Controlbar{}
	cb.cm = c

	th := cb.Theme()

	cb.gfx = &canvas.Image{
		Resource: theme.QuestionIcon(),
		FillMode: canvas.ImageFillContain,
	}

	cb.labelname = &canvas.Text{
		Text:     "Name",
		TextSize: theme.TextSize(),
	}
	cb.labelexp = &canvas.Text{
		Text:     "EP",
		TextSize: theme.TextSize(),
	}
	cb.labellvl = &canvas.Text{
		Text:     "Grad",
		TextSize: theme.TextSize(),
	}

	cb.name = &widget.Entry{
		PlaceHolder: "Charactername",
		OnChanged: func(s string) {
			cb.cm.SelectedCharacter.Name = s
		},
	}
	cb.exp = NewNumericEntry()
	cb.exp.PlaceHolder = "EP"
	cb.exp.OnChanged = func(s string) {
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			cb.cm.SelectedCharacter.Exp = int(i)
		}
	}

	cb.lvl = &canvas.Text{
		Text:     "",
		TextSize: theme.TextSize(),
	}

	cb.btnsave = &widget.Button{
		Icon:       th.Icon(theme.IconNameDocumentSave),
		Text:       "Speichern",
		Importance: widget.HighImportance,
	}
	cb.btndelete = &widget.Button{
		Icon:       th.Icon(theme.IconNameDelete),
		Text:       "Entfernen",
		Importance: widget.LowImportance,
	}

	cb.ExtendBaseWidget(cb)
	return cb
}

func (cb *Controlbar) CreateRenderer() fyne.WidgetRenderer {
	cb.ExtendBaseWidget(cb)

	cb.layoutcont = container.New(
		&controlbarlayout{
			gfx:       cb.gfx,
			labelname: cb.labelname,
			name:      cb.name,
			labelexp:  cb.labelexp,
			exp:       cb.exp,
			labellvl:  cb.labellvl,
			lvl:       cb.lvl,
			save:      cb.btnsave,
			del:       cb.btndelete,
		},
		cb.gfx,
		cb.labelname,
		cb.name,
		cb.labelexp,
		cb.exp,
		cb.labellvl,
		cb.lvl,
		cb.btnsave,
		cb.btndelete,
	)
	renderer := widget.NewSimpleRenderer(cb.layoutcont)
	cb.Refresh()
	return renderer
}

func (cb *Controlbar) Refresh() {
	if len(cb.cm.Characters) > 0 {
		if cb.cm.SelectedCharacter.IsAllValid() {
			cb.btnsave.Enable()
		} else {
			cb.btnsave.Disable()
		}

		cb.name.Text = cb.cm.SelectedCharacter.Name
		cb.exp.Text = fmt.Sprintf("%d", cb.cm.SelectedCharacter.Exp)
		if cb.cm.SelectedCharacter.Status&models.Activationmode(models.Levelup) == 0 {
			cb.exp.Disable()
		} else {
			cb.exp.Enable()
		}
		cb.lvl.Text = fmt.Sprintf("%d", cb.cm.SelectedCharacter.Level)

		if cb.layoutcont != nil {
			cb.layoutcont.Refresh()
		}
	}
	canvas.Refresh(cb)
}
