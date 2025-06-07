package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
)

var _ fyne.Widget = (*IntvalueItem)(nil)

type IntvalueItem struct {
	widget.BaseWidget

	layoutcont *fyne.Container

	splitelem fyne.CanvasObject
	label     *canvas.Text
	entry     *NumericEntry
	errorLbl  *canvas.Text

	elechan   chan *models.Character
	charmodel *CharacterModel
	ident     string

	data *models.Element

	savevalue func(string)
}

func NewIntvalueItem(chn chan *models.Character, model *CharacterModel, ident string) *IntvalueItem {
	iv := &IntvalueItem{}
	iv.elechan = chn
	iv.charmodel = model
	iv.ident = ident

	th := iv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	iv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	iv.label = canvas.NewText("", th.Color(theme.ColorNameForeground, v))
	iv.entry = NewNumericEntry()
	iv.entry.OnChanged = func(s string) {
		if iv.data != nil {
			iv.data.SetValue(s)
			iv.elechan <- iv.charmodel.SelectedCharacter
		}
	}

	t := canvas.NewText("", th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	iv.errorLbl = t

	iv.ExtendBaseWidget(iv)
	return iv
}

func (iv *IntvalueItem) CreateRenderer() fyne.WidgetRenderer {
	iv.ExtendBaseWidget(iv)

	iv.layoutcont = container.New(
		&valuelayout{
			split:   iv.splitelem,
			lbl:     iv.label,
			element: iv.entry,
			errmsg:  iv.errorLbl,
		},
		iv.splitelem,
		iv.label,
		iv.entry,
		iv.errorLbl,
	)

	renderer := widget.NewSimpleRenderer(iv.layoutcont)
	iv.Refresh()
	return renderer
}

func (iv *IntvalueItem) Refresh() {
	iv.ExtendBaseWidget(iv)

	if iv.charmodel.SelectedCharacter != nil {
		iv.data = iv.charmodel.SelectedCharacter.GetElement(iv.ident)
		iv.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				iv.Refresh()
			})
		}
		iv.label.Text = iv.data.Fieldtype.Label
		iv.entry.Text = iv.data.GetValueInfo(models.Value)
		if iv.data.GetValidation() {
			iv.errorLbl.Hide()
		} else {
			iv.errorLbl.Show()
		}
		iv.errorLbl.Text = iv.data.ErrorMsg
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	iv.layoutcont.Refresh()
}
