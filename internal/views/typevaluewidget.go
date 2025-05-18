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

	splitelem  fyne.CanvasObject
	layoutcont *fyne.Container
	label      *canvas.Text
	selectwg   *widget.Select
	errorLbl   *canvas.Text

	controller *CharacterController
	data       *models.Element

	savevalue func(string)
}

func NewTypevalueItem(ctrl *CharacterController, e *models.Element) *TypevalueItem {
	tv := &TypevalueItem{}

	tv.controller = ctrl
	tv.data = e

	tv.savevalue = func(s string) {
		tv.data.SetValue(s)
		tv.controller.Model.ApplyCreationRules()
		tv.controller.refreshbindings(tv.data.Fieldtype.Identify())
		tv.Refresh()
	}

	tv.ExtendBaseWidget(tv)

	ctrl.addbindings(e.Fieldtype.Identify(), tv)
	return tv
}

func (tv *TypevalueItem) CreateRenderer() fyne.WidgetRenderer {
	tv.ExtendBaseWidget(tv)

	th := tv.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	tv.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}

	tv.label = canvas.NewText(tv.data.Fieldtype.Label, th.Color(theme.ColorNameForeground, v))
	tv.selectwg = widget.NewSelect(nil, tv.savevalue)

	t := canvas.NewText(tv.data.ErrorMsg, th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	tv.errorLbl = t

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

	tv.selectwg.Options = tv.getoptions()
	tv.errorLbl.Text = tv.data.ErrorMsg
	if tv.data.ErrorMsg == "" {
		tv.errorLbl.Hide()
	} else {
		tv.errorLbl.Show()
	}
	tv.layoutcont.Refresh()
}

func (tv *TypevalueItem) MinSize() fyne.Size {
	return tv.BaseWidget.MinSize()
}

func (tv *TypevalueItem) getoptions() []string {
	switch ass := tv.data.Value.(type) {
	case *models.Typevalue:
		exists := false
		s := make([]string, 0)
		for _, v := range ass.Validvalues {
			if tv.selectwg != nil && tv.selectwg.Selected == v.Label {
				exists = true
			}
			s = append(s, v.Label)
		}

		if !exists && len(s) > 0 {
			tv.selectwg.Selected = s[0]
		}
		return s
	default:
		tv.selectwg.Selected = "Keine Angabe"
		return []string{"Keine Angabe"}
	}
}
