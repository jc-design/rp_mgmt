package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*dicelayout)(nil)

type dicelayout struct {
	split, lbl, entry, btn, desc, errmsg fyne.CanvasObject
}

func (dl *dicelayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	dicevaluewidth := float32(100)
	labelwidth := float32(100)
	row1height := dl.entry.MinSize().Height
	row2height := float32(0)

	if dl.errmsg.Visible() {
		row2height = dl.errmsg.MinSize().Height
	}
	totalheight := row1height + row2height + 2*innerPad
	return fyne.Size{
		Width:  2*labelwidth + dicevaluewidth + row1height + 3*innerPad,
		Height: totalheight,
	}
}

func (dl *dicelayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {

	innerPad := theme.Size(theme.SizeNameInnerPadding)
	stroke := dl.split.(*canvas.Line).StrokeWidth

	dicevaluewidth := float32(100)
	labelwidth := float32(100)
	row1height := dl.entry.MinSize().Height

	totalheight := row1height + 2*innerPad
	row2height := float32(0)

	if dl.errmsg.Visible() {
		row2height = dl.errmsg.MinSize().Height
		totalheight = totalheight + row2height + innerPad
	}

	dl.split.Move(fyne.Position{
		X: 0,
		Y: totalheight,
	})
	dl.split.Resize(fyne.Size{
		Width:  size.Width,
		Height: 0,
	})

	dl.lbl.Move(fyne.Position{
		X: 0,
		Y: (totalheight - dl.lbl.MinSize().Height - stroke) / 2,
	})
	dl.lbl.Resize(fyne.Size{
		Width:  labelwidth,
		Height: dl.lbl.MinSize().Height,
	})

	dl.entry.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: innerPad - stroke,
	})
	dl.entry.Resize(fyne.Size{
		Width:  dicevaluewidth,
		Height: row1height,
	})

	dl.btn.Move(fyne.Position{
		X: labelwidth + dicevaluewidth + 2*innerPad,
		Y: innerPad - stroke,
	})
	dl.btn.Resize(fyne.Size{
		Width:  row1height,
		Height: row1height,
	})

	dl.desc.Move(fyne.Position{
		X: labelwidth + dicevaluewidth + row1height + 3*innerPad,
		Y: innerPad - stroke,
	})
	dl.desc.Resize(fyne.Size{
		Width:  size.Width - labelwidth - dicevaluewidth - 3*innerPad,
		Height: row1height,
	})

	dl.errmsg.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: row1height + 2*innerPad,
	})
	dl.errmsg.Resize(fyne.Size{
		Width:  size.Width - labelwidth - 1*innerPad,
		Height: row2height,
	})
}
