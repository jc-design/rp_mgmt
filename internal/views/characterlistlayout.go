package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*characterlistlayout)(nil)

type characterlistlayout struct {
	btn, list, content fyne.CanvasObject
}

func (cl *characterlistlayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	labelwidth := float32(100)

	return fyne.Size{
		Width:  labelwidth,
		Height: cl.btn.MinSize().Height * 2,
	}
}

func (cl *characterlistlayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	colwidth := float32(200)

	cl.btn.Move(fyne.Position{
		X: 0,
		Y: 0,
	})
	cl.btn.Resize(fyne.Size{
		Width:  colwidth,
		Height: cl.btn.MinSize().Height,
	})

	cl.list.Move(fyne.Position{
		X: 0,
		Y: cl.btn.MinSize().Height + innerPad,
	})
	cl.list.Resize(fyne.Size{
		Width:  colwidth,
		Height: size.Height - cl.btn.MinSize().Height - innerPad,
	})

	cl.content.Move(fyne.Position{
		X: colwidth + innerPad,
		Y: 0,
	})
	cl.content.Resize(fyne.Size{
		Width:  size.Width - innerPad - colwidth,
		Height: size.Height,
	})
}
