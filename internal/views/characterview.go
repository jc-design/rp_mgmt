package views

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/jc-design/rp_mgmt/internal/models"
)

const (
	sidebar float32 = 220
)

type baselayout struct {
	grad1, grad2, top, left, content fyne.CanvasObject
}

func newbaselayout(grad1, grad2, top, left, content fyne.CanvasObject) fyne.Layout {
	return &baselayout{
		grad1:   grad1,
		grad2:   grad2,
		top:     top,
		left:    left,
		content: content,
	}
}

func (l *baselayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	l.grad1.Resize(size)
	l.grad2.Resize(size)
	l.top.Resize(fyne.NewSize(size.Width, l.top.MinSize().Height))
	l.left.Move(fyne.NewPos(0, l.top.MinSize().Height))
	l.left.Resize(fyne.NewSize(sidebar, size.Height-l.top.MinSize().Height))
	l.content.Move(fyne.NewPos(sidebar, l.top.MinSize().Height))
	l.content.Resize(fyne.NewSize(size.Width-sidebar, size.Height-l.top.MinSize().Height))
}

func (l *baselayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}

func (ctrl *CharacterController) CreateToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			ctrl.Model.NewCharacter()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {}),
		widget.NewToolbarAction(theme.FolderNewIcon(), func() {}),
	)
	return toolbar
}

func (ctrl *CharacterController) CreateCharacterList() *widget.List {
	list := widget.NewList(
		func() int {
			return len(ctrl.Model.Characters)
		},
		func() fyne.CanvasObject {
			name := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop := canvas.NewText("template", theme.Color(theme.ColorNameForeground))
			prop.TextSize = name.TextSize * 0.8

			return container.NewVBox(name, prop)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			vbox := o.(*fyne.Container)
			vbox.Objects[0].(*canvas.Text).Text = ctrl.Model.Characters[i].Name
			vbox.Objects[1].(*canvas.Text).Text = fmt.Sprintf("R:%v - K:%v - G:%v",
				ctrl.Model.Characters[i].GetValueInfo("baseproperty|race", models.Value),
				ctrl.Model.Characters[i].GetValueInfo("baseproperty|class", models.Value),
				ctrl.Model.Characters[i].Level,
			)
		})
	return list
}

func (ctrl *CharacterController) CreateContent() fyne.CanvasObject {
	elements := make([]fyne.CanvasObject, 0)
	for _, e := range ctrl.Model.SelectedCharacter.Properties {
		elements = append(elements, createElementConatainer(e))
	}
	return container.NewVScroll(container.NewVBox(elements...))
}

func (ctrl *CharacterController) CreateBaseLayout(top, left, content fyne.CanvasObject) *fyne.Container {
	grad1 := canvas.NewLinearGradient(
		color.NRGBA{R: 0x6e, G: 0xc8, B: 0xfa, A: 0x77},
		color.NRGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xff},
		315)
	grad2 := canvas.NewLinearGradient(
		color.NRGBA{R: 0xff, G: 0xff, B: 0xbf, A: 0x77},
		color.Transparent,
		225)
	return container.New(newbaselayout(grad1, grad2, top, left, content), grad1, grad2, top, left, content)
}

func createElementConatainer(e *models.Element) *fyne.Container {

	switch e.Value.(type) {
	case *models.Intvalue:
		return container.NewHBox(
			canvas.NewText(e.Fieldtype.Label, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Id, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Type, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Value.GetInfo(models.Value), theme.Color(theme.ColorNameForeground)),
		)
	case *models.Stringvalue:
		return container.NewHBox(
			canvas.NewText(e.Fieldtype.Label, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Id, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Type, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Value.GetInfo(models.Value), theme.Color(theme.ColorNameForeground)),
		)

	case *models.Dice:
		return container.NewHBox(
			canvas.NewText(e.Fieldtype.Label, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Id, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Type, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Value.GetInfo(models.Value), theme.Color(theme.ColorNameForeground)),
		)

	case *models.Typevalue:
		return container.NewHBox(
			canvas.NewText(e.Fieldtype.Label, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Id, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Fieldtype.Type, theme.Color(theme.ColorNameForeground)),
			canvas.NewText(e.Value.GetInfo(models.Value), theme.Color(theme.ColorNameForeground)),
		)

	default:
		return container.NewHBox(
			canvas.NewText("---", theme.Color(theme.ColorNameForeground)),
			canvas.NewText("---", theme.Color(theme.ColorNameForeground)),
			canvas.NewText("---", theme.Color(theme.ColorNameForeground)),
			canvas.NewText("---", theme.Color(theme.ColorNameForeground)),
		)
	}
}
