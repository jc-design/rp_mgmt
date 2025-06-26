package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*skilllayout)(nil)

type skilllayout struct {
	split, lbl, entry, btn, desc, errmsg fyne.CanvasObject
}

func (sl *skilllayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	dicevaluewidth := float32(100)
	labelwidth := float32(100)
	row1height := sl.entry.MinSize().Height
	row2height := float32(0)

	if sl.errmsg.Visible() {
		row2height = sl.errmsg.MinSize().Height
	}
	totalheight := row1height + row2height + 2*innerPad
	return fyne.Size{
		Width:  2*labelwidth + dicevaluewidth + row1height + 3*innerPad,
		Height: totalheight,
	}
}

func (sl *skilllayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {

	innerPad := theme.Size(theme.SizeNameInnerPadding)
	stroke := sl.split.(*canvas.Line).StrokeWidth

	dicevaluewidth := float32(100)
	labelwidth := float32(200) + innerPad
	row1height := sl.entry.MinSize().Height

	totalheight := row1height + 2*innerPad
	row2height := float32(0)

	if sl.errmsg.Visible() {
		row2height = sl.errmsg.MinSize().Height
		totalheight = totalheight + row2height + innerPad
	}

	sl.split.Move(fyne.Position{
		X: 0,
		Y: totalheight,
	})
	sl.split.Resize(fyne.Size{
		Width:  size.Width,
		Height: 0,
	})

	sl.lbl.Move(fyne.Position{
		X: 0,
		Y: (totalheight - sl.lbl.MinSize().Height - stroke) / 2,
	})
	sl.lbl.Resize(fyne.Size{
		Width:  labelwidth,
		Height: sl.lbl.MinSize().Height,
	})

	sl.entry.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: innerPad - stroke,
	})
	sl.entry.Resize(fyne.Size{
		Width:  dicevaluewidth,
		Height: row1height,
	})

	sl.btn.Move(fyne.Position{
		X: labelwidth + dicevaluewidth + 2*innerPad,
		Y: innerPad - stroke,
	})
	sl.btn.Resize(fyne.Size{
		Width:  row1height,
		Height: row1height,
	})

	sl.desc.Move(fyne.Position{
		X: labelwidth + dicevaluewidth + row1height + 3*innerPad,
		Y: innerPad - stroke,
	})
	sl.desc.Resize(fyne.Size{
		Width:  size.Width - labelwidth - dicevaluewidth - 3*innerPad,
		Height: row1height,
	})

	sl.errmsg.Move(fyne.Position{
		X: labelwidth + innerPad,
		Y: row1height + 2*innerPad,
	})
	sl.errmsg.Resize(fyne.Size{
		Width:  size.Width - labelwidth - 1*innerPad,
		Height: row2height,
	})
}
