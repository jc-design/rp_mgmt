package views

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*NumericEntry)(nil)

type NumericEntry struct {
	widget.Entry
}

func NewNumericEntry() *NumericEntry {
	entry := &NumericEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *NumericEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

func (e *NumericEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.Atoi(content); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *NumericEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
