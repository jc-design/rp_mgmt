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

var _ fyne.Widget = (*StringvalueItem)(nil)

type StringvalueItem struct {
	widget.BaseWidget

	layoutcont *fyne.Container

	splitelem fyne.CanvasObject
	label     *canvas.Text
	entry     *widget.Entry
	errorLbl  *canvas.Text

	elechan   chan *models.Character
	charmodel *CharacterModel
	ident     string

	data *models.Element

	savevalue func(string)
}

func NewStringvalueItem(chn chan *models.Character, model *CharacterModel, ident string) *StringvalueItem {
	sv := &StringvalueItem{}
	sv.elechan = chn
	sv.charmodel = model
	sv.ident = ident

	th := sv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	sv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	sv.label = canvas.NewText("", th.Color(theme.ColorNameForeground, v))
	sv.entry = widget.NewEntry()
	sv.entry.OnChanged = func(s string) {
		if sv.data != nil {
			sv.data.SetValue(s)
			sv.elechan <- sv.charmodel.SelectedCharacter
		}
	}

	t := canvas.NewText("", th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	sv.errorLbl = t

	sv.ExtendBaseWidget(sv)
	return sv
}

func (sv *StringvalueItem) CreateRenderer() fyne.WidgetRenderer {
	sv.ExtendBaseWidget(sv)

	sv.layoutcont = container.New(
		&valuelayout{
			split:   sv.splitelem,
			lbl:     sv.label,
			element: sv.entry,
			errmsg:  sv.errorLbl,
		},
		sv.splitelem,
		sv.label,
		sv.entry,
		sv.errorLbl,
	)

	renderer := widget.NewSimpleRenderer(sv.layoutcont)
	sv.Refresh()
	return renderer
}

func (sv *StringvalueItem) Refresh() {
	sv.ExtendBaseWidget(sv)
	if sv.charmodel.SelectedCharacter != nil {
		sv.data = sv.charmodel.SelectedCharacter.GetElement(sv.ident)
		sv.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				sv.Refresh()
			})
		}
		sv.label.Text = sv.data.Fieldtype.Label
		sv.entry.Text = sv.data.GetValueInfo(models.Value)
		if sv.data.GetValidation() {
			sv.errorLbl.Hide()
		} else {
			sv.errorLbl.Show()
		}
		sv.errorLbl.Text = sv.data.ErrorMsg
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	sv.layoutcont.Refresh()
}
