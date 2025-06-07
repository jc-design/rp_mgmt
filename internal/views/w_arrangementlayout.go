package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*w_arrangementlayout)(nil)

type w_arrangementlayout struct {
	listelement, content, vline fyne.CanvasObject
}

func (l *w_arrangementlayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(l.listelement.MinSize().Width+l.content.MinSize().Width, l.listelement.MinSize().Height)
}

func (l *w_arrangementlayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	listelementwidth := float32(200)
	thickness := l.vline.(*canvas.Line).StrokeWidth

	l.listelement.Move(fyne.NewPos(0, 0))
	l.listelement.Resize(fyne.NewSize(listelementwidth, size.Height))

	l.vline.Move(fyne.NewPos(listelementwidth+innerPad, 0))
	l.vline.Resize(fyne.NewSize(0, size.Height))

	l.content.Move(fyne.NewPos(listelementwidth+innerPad+thickness, 0))
	l.content.Resize(fyne.NewSize(size.Width-listelementwidth-innerPad-thickness, size.Height))
}
