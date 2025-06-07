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

var _ fyne.Widget = (*W_Arrangement)(nil)

type W_Arrangement struct {
	widget.BaseWidget

	// basic windowelements
	listelement *Characterlist
	content     *fyne.Container

	// elements for character creation/update
	toolbar    *Controlbar
	properties fyne.CanvasObject

	// data elements
	elechan        chan *models.Character
	charmodel      *CharacterModel
	contentchecked bool
}

func refresh_onselected(f func(i int), c fyne.CanvasObject) func(i int) {
	return func(i int) {
		if f != nil {
			f(i)
		}
		if c != nil {
			c.Refresh()
		}
	}
}

func NewW_Arrangement(chn chan *models.Character, model *CharacterModel) *W_Arrangement {
	w := &W_Arrangement{
		elechan:   chn,
		charmodel: model,
	}

	w.listelement = NewCharacterlist(chn, model)
	w.content = container.NewStack()

	// initialize elements
	w.toolbar = NewControlbar(model)
	w.properties = nil

	// check count of []Character, once the OnAdded is called
	// when count == 1 set contentchecked = false
	// because we have visual change from no to one character
	w.listelement.OnAdded = func(f func(i int)) func(i int) {
		return func(i int) {
			if f != nil {
				f(i)
			}
			if len(w.charmodel.Characters) == 1 {
				w.contentchecked = false
				w.Refresh()
			}
		}
	}(w.listelement.OnAdded)

	// refresh toolbar, when OnSelected ist called
	w.listelement.OnSelected = refresh_onselected(w.listelement.OnSelected, w.toolbar)

	// refresh listelement, when btndelete.OnTapped is called
	// when count == 0 set contentchecked = false
	// because we have visual change from one to no character
	w.toolbar.btndelete.OnTapped = func(f func()) func() {
		return func() {
			if f != nil {
				f()
			}
			w.charmodel.RemoveCharacter(w.listelement.selectedid)
			if len(w.charmodel.Characters) == 0 {
				w.contentchecked = false
				w.Refresh()
			} else {
				w.listelement.characterlist.Select(len(w.charmodel.Characters) - 1)
				w.listelement.characterlist.OnSelected(len(w.charmodel.Characters) - 1)
			}
			w.listelement.Refresh()
		}
	}(w.toolbar.btndelete.OnTapped)

	// refresh listitem, when name.OnChanged is called
	w.toolbar.name.OnChanged = func(f func(s string)) func(s string) {
		return func(s string) {
			if f != nil {
				f(s)
			}
			w.listelement.characterlist.RefreshItem(w.listelement.selectedid)
		}
	}(w.toolbar.name.OnChanged)
	// refresh listitem, when exp.OnChanged is called
	w.toolbar.exp.OnChanged = func(f func(s string)) func(s string) {
		return func(s string) {
			if f != nil {
				f(s)
			}
			w.listelement.characterlist.RefreshItem(w.listelement.selectedid)
		}
	}(w.toolbar.exp.OnChanged)

	w.ExtendBaseWidget(w)
	return w
}

func (w *W_Arrangement) CreateRenderer() fyne.WidgetRenderer {

	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)
	vline := &canvas.Line{
		StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
		StrokeWidth: 2,
	}
	layout := container.New(
		&w_arrangementlayout{
			listelement: w.listelement,
			content:     w.content,
			vline:       vline,
		},
		w.listelement,
		w.content,
		vline,
	)
	renderer := widget.NewSimpleRenderer(layout)
	w.Refresh()
	return renderer
}

func (w *W_Arrangement) Refresh() {
	w.ExtendBaseWidget(w)

	th := w.Theme()
	v := fyne.CurrentApp().Settings().ThemeVariant()

	primecol := th.Color(theme.ColorNamePrimary, v).(color.NRGBA)

	if !w.contentchecked {
		if len(w.charmodel.Characters) == 0 {
			w.content.Objects = []fyne.CanvasObject{&canvas.Text{
				Text:      "Kein Charackter ausgewählt",
				Color:     theme.Color(theme.ColorNameForeground),
				Alignment: fyne.TextAlignCenter,
				TextSize:  24,
			}}
			w.contentchecked = true
		} else {
			hline := &canvas.Line{
				StrokeColor: color.NRGBA{R: primecol.R, G: primecol.G, B: primecol.B, A: 31},
				StrokeWidth: 2,
			}
			w.properties = w.getApptab()
			w.content.Objects = []fyne.CanvasObject{container.New(&propertieslayout{toolbar: w.toolbar, content: w.properties, hline: hline},
				w.toolbar,
				w.properties,
				hline,
			)}
			w.contentchecked = true
		}
	}

	w.listelement.Refresh()
	w.content.Refresh()

	w.toolbar.Refresh()
	if w.properties != nil {
		w.properties.Refresh()
	}

	canvas.Refresh(w)
}

func (w *W_Arrangement) getApptab() fyne.CanvasObject {
	items := []*container.TabItem{}
	for _, group := range w.charmodel.SelectedCharacter.PropsGrouped {
		vbox := container.NewVBox()
		for _, e := range group.Elements {
			var c fyne.CanvasObject
			if w.charmodel.SelectedCharacter.Status&e.Editable == 1 {
				switch e.Value.(type) {
				case *models.Intvalue:
					c = NewIntvalueItem(w.elechan, w.charmodel, e.Fieldtype.Identify())
					vbox.Add(c)
					break
				case *models.Stringvalue:
					c = NewStringvalueItem(w.elechan, w.charmodel, e.Fieldtype.Identify())
					vbox.Add(c)
					break
				case *models.Typevalue:
					c = NewTypevalueItem(w.elechan, w.charmodel, e.Fieldtype.Identify())
					vbox.Add(c)
					break
				case *models.Dice:
					c = NewDiceItem(w.elechan, w.charmodel, e.Fieldtype.Identify())
					vbox.Add(c)
					break
				default:
				}
			} else {
				c = NewDisabledItem(w.elechan, w.charmodel, e.Fieldtype.Identify())
				vbox.Add(c)
			}
			w.listelement.OnSelected = refresh_onselected(w.listelement.OnSelected, c)
		}
		items = append(items, container.NewTabItem(group.Group.Label, container.NewVScroll(vbox)))
	}
	tabs := container.NewAppTabs(items...)
	return tabs
}
