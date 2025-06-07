package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*characterlistlayout)(nil)

type characterlistlayout struct {
	btn, list fyne.CanvasObject
}

func (cl *characterlistlayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	size := cl.btn.MinSize().Add(cl.list.MinSize())
	return size
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
}
