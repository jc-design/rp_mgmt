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

var _ fyne.Widget = (*TypevalueItem)(nil)

type TypevalueItem struct {
	widget.BaseWidget

	layoutcont *fyne.Container

	splitelem fyne.CanvasObject
	label     *canvas.Text
	selectwg  *widget.Select
	errorLbl  *canvas.Text

	elechan   chan *models.Character
	charmodel *CharacterModel
	ident     string

	data *models.Element
}

func NewTypevalueItem(chn chan *models.Character, model *CharacterModel, ident string) *TypevalueItem {
	tv := &TypevalueItem{}

	tv.elechan = chn
	tv.charmodel = model
	tv.ident = ident

	th := tv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	tv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}

	tv.label = canvas.NewText("", th.Color(theme.ColorNameForeground, v))
	tv.selectwg = widget.NewSelect(nil, func(s string) {
		if tv.data != nil {
			tv.data.SetValue(s)
			tv.elechan <- tv.charmodel.SelectedCharacter
		}
	})

	t := canvas.NewText("", th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	tv.errorLbl = t

	tv.ExtendBaseWidget(tv)
	return tv
}

func (tv *TypevalueItem) CreateRenderer() fyne.WidgetRenderer {
	tv.ExtendBaseWidget(tv)

	tv.layoutcont = container.New(
		&valuelayout{
			split:   tv.splitelem,
			lbl:     tv.label,
			element: tv.selectwg,
			errmsg:  tv.errorLbl,
		},
		tv.splitelem,
		tv.label,
		tv.selectwg,
		tv.errorLbl,
	)

	renderer := widget.NewSimpleRenderer(tv.layoutcont)
	tv.Refresh()
	return renderer
}

func (tv *TypevalueItem) Refresh() {
	tv.ExtendBaseWidget(tv)
	if tv.charmodel.SelectedCharacter != nil {
		tv.data = tv.charmodel.SelectedCharacter.GetElement(tv.ident)
		tv.data.OnValidated = func(b bool) {
			fyne.Do(func() {
				tv.Refresh()
			})
		}

		tv.label.Text = tv.data.Fieldtype.Label
		tv.selectwg.Options = tv.getoptions()
		if tv.data.GetValidation() {
			tv.errorLbl.Hide()
		} else {
			tv.errorLbl.Show()
		}
		tv.errorLbl.Text = tv.data.ErrorMsg
	}

	// call Refresh() on layout, so alle CanvasObjects are correctly redrawn
	tv.layoutcont.Refresh()
}

func (tv *TypevalueItem) getoptions() []string {
	switch ass := tv.data.Value.(type) {
	case *models.Typevalue:
		// exists := false
		index := 0
		s := make([]string, 0)
		for i, v := range ass.Validvalues {
			if tv.selectwg != nil && tv.data.GetValueInfo(models.Value) == v.Label {
				index = i
			}
			s = append(s, v.Label)
		}
		if len(s) > 0 {
			tv.selectwg.Selected = s[index]
		} else {
			tv.selectwg.Selected = ""
		}
		return s
	default:
		tv.selectwg.Selected = ""
		return []string{}
	}
}
