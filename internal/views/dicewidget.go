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

const ()

var _ fyne.Widget = (*DiceItem)(nil)

type DiceItem struct {
	widget.BaseWidget

	splitelem      fyne.CanvasObject
	layoutcont     *fyne.Container
	label          *canvas.Text
	entry          *NumericEntry
	button         *widget.Button
	descriptionLbl *canvas.Text
	errorLbl       *canvas.Text

	controller *CharacterController
	data       *models.Element

	rolldice  func()
	savevalue func(string)
}

func NewDiceItem(ctrl *CharacterController, e *models.Element) *DiceItem {
	di := &DiceItem{}
	di.controller = ctrl
	di.data = e

	di.rolldice = func() {
		di.data.Execute()
		di.entry.SetText(di.data.Value.GetInfo(models.Value))
		di.controller.Model.ApplyCreationRules()
		di.controller.refreshbindings(di.data.Fieldtype.Identify())
		di.Refresh()
	}

	di.savevalue = func(s string) {
		di.data.SetValue(s)
		di.controller.Model.ApplyCreationRules()
		di.controller.refreshbindings(di.data.Fieldtype.Identify())
		di.Refresh()
	}

	di.ExtendBaseWidget(di)

	ctrl.addbindings(e.Fieldtype.Identify(), di)
	return di
}

func (di *DiceItem) CreateRenderer() fyne.WidgetRenderer {
	di.ExtendBaseWidget(di)

	th := di.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)
	di.splitelem = &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	di.label = canvas.NewText(di.data.Fieldtype.Label, th.Color(theme.ColorNameForeground, v))

	di.entry = NewNumericEntry()
	di.entry.OnChanged = di.savevalue

	di.button = &widget.Button{
		Icon:     resourceDicedarkSvg,
		OnTapped: di.rolldice,
	}
	di.descriptionLbl = canvas.NewText(di.data.Value.GetInfo(models.Description), th.Color(theme.ColorNameForeground, v))

	t := canvas.NewText(di.data.ErrorMsg, th.Color(theme.ColorNameError, v))
	t.TextSize = t.TextSize * 0.8
	di.errorLbl = t
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

	di.entry.Text = di.data.Value.GetInfo(models.Value)
	di.descriptionLbl.Text = di.data.Value.GetInfo(models.Description)
	di.errorLbl.Text = di.data.ErrorMsg
	if di.data.ErrorMsg == "" {
		di.errorLbl.Hide()
	} else {
		di.errorLbl.Show()
	}
	di.layoutcont.Refresh()
	// canvas.Refresh(di)
}

func (di *DiceItem) MinSize() fyne.Size {
	return di.BaseWidget.MinSize()
}
