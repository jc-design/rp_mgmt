package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*propertieslayout)(nil)

type propertieslayout struct {
	toolbar, content, hline fyne.CanvasObject
}

func (l *propertieslayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(l.toolbar.MinSize().Width, l.toolbar.MinSize().Height*3)
}

func (l *propertieslayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	thickness := l.hline.(*canvas.Line).StrokeWidth

	l.toolbar.Move(fyne.NewPos(innerPad, 0))
	l.toolbar.Resize(fyne.NewSize(size.Width, l.toolbar.MinSize().Height))

	l.hline.Move(fyne.NewPos(0, l.toolbar.MinSize().Height+innerPad))
	l.hline.Resize(fyne.NewSize(size.Width, 0))

	l.content.Move(fyne.NewPos(innerPad, l.toolbar.MinSize().Height+innerPad+thickness+innerPad))
	l.content.Resize(fyne.NewSize(size.Width-innerPad, size.Height-l.toolbar.MinSize().Height-innerPad-thickness-innerPad))
}
