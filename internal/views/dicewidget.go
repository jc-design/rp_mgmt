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

var _ fyne.Widget = (*DiceItem)(nil)

type DiceItem struct {
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

func NewDiceItem(chn chan *models.Character, model *CharacterModel, ident string) *DiceItem {
	di := &DiceItem{}
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

	di.entry = NewNumericEntry()
	di.entry.OnChanged = func(s string) {
		if di.data != nil {
			di.data.SetValue(s)
			di.elechan <- di.charmodel.SelectedCharacter
		}
	}

	di.button = &widget.Button{
		Icon: resourceDicedarkSvg,
		OnTapped: func() {
			if di.data != nil {
				di.data.Execute()
				v := di.data.GetValueInfo(models.Value)
				di.entry.SetText(v)
				di.Refresh()
				di.elechan <- di.charmodel.SelectedCharacter
			}
		},
	}
	di.descriptionLbl = canvas.NewText("", th.Color(theme.ColorNameForeground, v))

	t := canvas.NewText("", th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	di.errorLbl = t

	di.ExtendBaseWidget(di)
	return di
}

func (di *DiceItem) CreateRenderer() fyne.WidgetRenderer {
	di.ExtendBaseWidget(di)

	di.layoutcont = container.New(
		&dicelayout{
			split:  di.splitelem,
			lbl:    di.label,
			entry:  di.entry,
			btn:    di.button,
			desc:   di.descriptionLbl,
			errmsg: di.errorLbl,
		},
		di.splitelem,
		di.label,
		di.entry,
		di.button,
		di.descriptionLbl,
		di.errorLbl,
	)
	renderer := widget.NewSimpleRenderer(di.layoutcont)
	di.Refresh()
	return renderer
}

func (di *DiceItem) Refresh() {
	di.ExtendBaseWidget(di)
	if di.charmodel.SelectedCharacter != nil {
		di.data = di.charmodel.SelectedCharacter.GetElement(di.ident)
		// maybe this isn't even necessary
		// check if refresh() in charactercontroller.go|NewCharactewrController - go func() is sufficant!
		di.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				di.Refresh()
			})
		}
		di.label.Text = di.data.Fieldtype.Label
		di.entry.Text = di.data.GetValueInfo(models.Value)
		di.descriptionLbl.Text = di.data.GetValueInfo(models.Description)
		if di.data.GetValidation() {
			di.errorLbl.Hide()
		} else {
			di.errorLbl.Show()
		}
		di.errorLbl.Text = di.data.ErrorMsg
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	di.layoutcont.Refresh()
}
