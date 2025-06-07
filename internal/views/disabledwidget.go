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

var _ fyne.Widget = (*DisabledItem)(nil)

type DisabledItem struct {
	widget.BaseWidget

	layoutcont *fyne.Container

	splitelem fyne.CanvasObject
	label     *canvas.Text
	value     *canvas.Text
	errorLbl  *canvas.Text

	elechan   chan *models.Character
	charmodel *CharacterModel
	ident     string

	data *models.Element
}

func NewDisabledItem(chn chan *models.Character, model *CharacterModel, ident string) *DisabledItem {
	di := &DisabledItem{}
	di.elechan = chn
	di.charmodel = model
	di.ident = ident

	th := di.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	di.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	di.label = canvas.NewText("", th.Color(theme.ColorNameForeground, v))
	di.value = canvas.NewText("", th.Color(theme.ColorNameForeground, v))
	di.errorLbl = canvas.NewText("", th.Color(theme.ColorNameForeground, v))

	di.ExtendBaseWidget(di)
	return di
}

func (di *DisabledItem) CreateRenderer() fyne.WidgetRenderer {
	di.ExtendBaseWidget(di)

	di.layoutcont = container.New(
		&valuelayout{
			split:   di.splitelem,
			lbl:     di.label,
			element: di.value,
			errmsg:  di.errorLbl,
		},
		di.splitelem,
		di.label,
		di.value,
		di.errorLbl,
	)

	renderer := widget.NewSimpleRenderer(di.layoutcont)
	di.Refresh()
	return renderer
}

func (di *DisabledItem) Refresh() {
	di.ExtendBaseWidget(di)
	if di.charmodel.SelectedCharacter != nil {
		di.data = di.charmodel.SelectedCharacter.GetElement(di.ident)
		di.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				di.Refresh()
			})
		}
		di.label.Text = di.data.Fieldtype.Label
		di.value.Text = di.data.GetValueInfo(models.Value)
		di.errorLbl.Hide()
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	di.layoutcont.Refresh()
}
