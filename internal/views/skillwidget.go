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

var _ fyne.Widget = (*SkillItem)(nil)

type SkillItem struct {
	widget.BaseWidget

	layoutcont *fyne.Container

	splitelem      fyne.CanvasObject
	label          *canvas.Text
	entry          *NumericEntry
	button         *widget.Button
	descriptionLbl *canvas.Text
	errorLbl       *canvas.Text

	elechan   chan *models.Character
	charmodel *CharacterModel
	ident     string
	data      *models.Element
}

func NewSkillItem(chn chan *models.Character, model *CharacterModel, ident string) *SkillItem {
	si := &SkillItem{}
	si.elechan = chn
	si.charmodel = model
	si.ident = ident

	th := si.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)
	si.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	si.label = canvas.NewText("", th.Color(theme.ColorNameForeground, v))

	si.entry = NewNumericEntry()
	si.entry.OnChanged = func(s string) {
		if si.data != nil {
			si.data.SetValue(s)
			si.elechan <- si.charmodel.SelectedCharacter
		}
	}

	si.button = &widget.Button{
		Icon: resourceDicedarkSvg,
		OnTapped: func() {
			if si.data != nil {
				si.data.Execute()
				v := si.data.GetValueInfo(models.Value)
				si.entry.SetText(v)
				si.Refresh()
				si.elechan <- si.charmodel.SelectedCharacter
			}
		},
	}
	si.descriptionLbl = canvas.NewText("", th.Color(theme.ColorNameForeground, v))

	t := canvas.NewText("", th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	si.errorLbl = t

	si.ExtendBaseWidget(si)
	return si
}

func (si *SkillItem) CreateRenderer() fyne.WidgetRenderer {
	si.ExtendBaseWidget(si)

	si.layoutcont = container.New(
		&skilllayout{
			split:  si.splitelem,
			lbl:    si.label,
			entry:  si.entry,
			btn:    si.button,
			desc:   si.descriptionLbl,
			errmsg: si.errorLbl,
		},
		si.splitelem,
		si.label,
		si.entry,
		si.button,
		si.descriptionLbl,
		si.errorLbl,
	)
	renderer := widget.NewSimpleRenderer(si.layoutcont)
	si.Refresh()
	return renderer
}

func (si *SkillItem) Refresh() {
	si.ExtendBaseWidget(si)
	if si.charmodel.SelectedCharacter != nil {
		si.data = si.charmodel.SelectedCharacter.GetElement(si.ident)
		// maybe this isn't even necessary
		// check if refresh() in charactercontroller.go|NewCharactewrController - go func() is sufficant!
		si.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				si.Refresh()
			})
		}
		si.label.Text = si.data.Fieldtype.Label
		si.entry.Text = si.data.GetValueInfo(models.Value)
		si.descriptionLbl.Text = si.data.GetValueInfo(models.Description)
		fmt.Println(si.descriptionLbl.Text)
		if si.data.GetValidation() {
			si.errorLbl.Hide()
		} else {
			si.errorLbl.Show()
		}
		si.errorLbl.Text = si.data.ErrorMsg
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	si.layoutcont.Refresh()
}
