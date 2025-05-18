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

	splitelem  fyne.CanvasObject
	layoutcont *fyne.Container
	label      *canvas.Text
	entry      *NumericEntry
	errorLbl   *canvas.Text

	controller *CharacterController
	data       *models.Element

	savevalue func(string)
}

func NewIntvalueItem(ctrl *CharacterController, e *models.Element) *IntvalueItem {
	iv := &IntvalueItem{}
	iv.controller = ctrl
	iv.data = e

	iv.savevalue = func(s string) {
		iv.data.SetValue(s)
		iv.controller.Model.ApplyCreationRules()
		iv.Refresh()
	}

	iv.ExtendBaseWidget(iv)

	return iv
}

func (iv *IntvalueItem) CreateRenderer() fyne.WidgetRenderer {
	iv.ExtendBaseWidget(iv)

	th := iv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	iv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	iv.label = canvas.NewText(iv.data.Fieldtype.Label, th.Color(theme.ColorNameForeground, v))
	iv.entry = NewNumericEntry()
	iv.entry.OnChanged = iv.savevalue

	t := canvas.NewText(iv.data.ErrorMsg, th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	iv.errorLbl = t

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

	iv.entry.Text = iv.data.Value.GetInfo(value)
	iv.errorLbl.Text = iv.data.ErrorMsg
	if iv.data.ErrorMsg == "" {
		iv.errorLbl.Hide()
	} else {
		iv.errorLbl.Show()
	}
	iv.layoutcont.Refresh()
	canvas.Refresh(iv)
}

func (iv *IntvalueItem) MinSize() fyne.Size {
	return iv.BaseWidget.MinSize()
}
