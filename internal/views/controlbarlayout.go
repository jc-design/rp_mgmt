package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Layout = (*controlbarlayout)(nil)

type controlbarlayout struct {
	gfx, labelname, name, labelexp, exp, labellvl, lvl, save, del fyne.CanvasObject
}

func (cl *controlbarlayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	innerPad := theme.Size(theme.SizeNameInnerPadding)

	return fyne.Size{
		Width:  640,
		Height: cl.name.MinSize().Height*3 + 2*innerPad,
	}
}

func (cl *controlbarlayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	innerPad := theme.Size(theme.SizeNameInnerPadding)
	colwidth := float32(100)
	btnwidth := float32(100)

	cl.gfx.Move(fyne.Position{
		X: 0,
		Y: 0,
	})
	cl.gfx.Resize(fyne.Size{
		Width:  colwidth,
		Height: cl.name.MinSize().Height*3 + 2*innerPad,
	})

	cl.labelname.Move(fyne.Position{
		X: colwidth + 1*innerPad,
		Y: 0,
	})
	cl.labelname.Resize(fyne.Size{
		Width:  colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.labelexp.Move(fyne.Position{
		X: colwidth + 1*innerPad,
		Y: cl.name.MinSize().Height*1 + 1*innerPad,
	})
	cl.labelexp.Resize(fyne.Size{
		Width:  colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.labellvl.Move(fyne.Position{
		X: colwidth + 1*innerPad,
		Y: cl.name.MinSize().Height*2 + 2*innerPad,
	})
	cl.labellvl.Resize(fyne.Size{
		Width:  colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.name.Move(fyne.Position{
		X: 2*colwidth + 2*innerPad,
		Y: 0,
	})
	cl.name.Resize(fyne.Size{
		Width:  2 * colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.exp.Move(fyne.Position{
		X: 2*colwidth + 2*innerPad,
		Y: cl.name.MinSize().Height*1 + 1*innerPad,
	})
	cl.exp.Resize(fyne.Size{
		Width:  2 * colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.lvl.Move(fyne.Position{
		X: 2*colwidth + 2*innerPad,
		Y: cl.name.MinSize().Height*2 + 2*innerPad,
	})
	cl.lvl.Resize(fyne.Size{
		Width:  2 * colwidth,
		Height: cl.name.MinSize().Height,
	})

	cl.save.Move(fyne.Position{
		X: size.Width - btnwidth - innerPad,
		Y: (cl.name.MinSize().Height*1 + 1*innerPad) / 2,
	})
	cl.save.Resize(fyne.Size{
		Width:  btnwidth,
		Height: cl.save.MinSize().Height,
	})

	cl.del.Move(fyne.Position{
		X: size.Width - btnwidth - innerPad,
		Y: (cl.name.MinSize().Height*1+1*innerPad)/2 + cl.name.MinSize().Height + 1*innerPad,
	})
	cl.del.Resize(fyne.Size{
		Width:  btnwidth,
		Height: cl.save.MinSize().Height,
	})

}
