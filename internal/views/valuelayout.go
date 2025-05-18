package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*valuelayout)(nil)

type valuelayout struct {
	split, lbl, element, errmsg fyne.CanvasObject
}

func (lt *valuelayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	labelwidth := float32(100)
	row1height := lt.element.MinSize().Height
	row2height := float32(0)

	if lt.errmsg.Visible() {
		row2height = lt.errmsg.MinSize().Height
	}
	totalheight := row1height + row2height + 2*innerPad
	return fyne.Size{
		Width:  2*labelwidth + 1*innerPad,
		Height: totalheight,
	}
}

func (lt *valuelayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {

	innerPad := theme.Size(theme.SizeNameInnerPadding)
	stroke := lt.split.(*canvas.Line).StrokeWidth

	labelwidth := float32(100)
	row1height := lt.element.MinSize().Height

	totalheight := row1height + 2*innerPad
	row2height := float32(0)

	if lt.errmsg.Visible() {
		row2height = lt.errmsg.MinSize().Height
		totalheight = totalheight + row2height + innerPad
	}

	lt.split.Move(fyne.Position{
		X: 0,
		Y: totalheight,
	})
	lt.split.Resize(fyne.Size{
		Width:  size.Width,
		Height: 0,
	})

	lt.lbl.Move(fyne.Position{
		X: 0,
		Y: (totalheight - lt.lbl.MinSize().Height - stroke) / 2,
	})
	lt.lbl.Resize(fyne.Size{
		Width:  labelwidth,
		Height: lt.lbl.MinSize().Height,
	})

	lt.element.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: innerPad - stroke,
	})
	lt.element.Resize(fyne.Size{
		Width:  size.Width - labelwidth - 1*innerPad,
		Height: row1height,
	})

	lt.errmsg.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: row1height + 2*innerPad,
	})
	lt.errmsg.Resize(fyne.Size{
		Width:  size.Width - labelwidth - 1*innerPad,
		Height: row2height,
	})
}
