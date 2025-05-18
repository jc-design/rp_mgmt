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

	splitelem  fyne.CanvasObject
	layoutcont *fyne.Container
	label      *canvas.Text
	entry      *NumericEntry
	errorLbl   *canvas.Text

	controller *CharacterController
	data       *models.Element

	savevalue func(string)
}

func NewStringvalueItem(ctrl *CharacterController, e *models.Element) *StringvalueItem {
	sv := &StringvalueItem{}
	sv.controller = ctrl
	sv.data = e

	sv.savevalue = func(s string) {
		sv.data.SetValue(s)
		sv.controller.Model.ApplyCreationRules()
		sv.Refresh()
	}

	sv.ExtendBaseWidget(sv)

	return sv
}

func (sv *StringvalueItem) CreateRenderer() fyne.WidgetRenderer {
	sv.ExtendBaseWidget(sv)

	th := sv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	sv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	sv.label = canvas.NewText(sv.data.Fieldtype.Label, th.Color(theme.ColorNameForeground, v))
	sv.entry = NewNumericEntry()
	sv.entry.OnChanged = sv.savevalue

	t := canvas.NewText(sv.data.ErrorMsg, th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	sv.errorLbl = t

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

	sv.entry.Text = sv.data.Value.GetInfo(value)
	sv.errorLbl.Text = sv.data.ErrorMsg
	if sv.data.ErrorMsg == "" {
		sv.errorLbl.Hide()
	} else {
		sv.errorLbl.Show()
	}
	sv.layoutcont.Refresh()
	canvas.Refresh(sv)
}

func (sv *StringvalueItem) MinSize() fyne.Size {
	return sv.BaseWidget.MinSize()
}
