package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App        *tview.Application
	Grid       *tview.Grid
	FooterText *tview.TextView
	TextInput  *tview.InputField
}

func InitTUI() *TUI {
	t := TUI{}

	t.App = tview.NewApplication()

	t.TextInput = tview.NewInputField()
	t.FooterText = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText("footer").SetTextColor(tcell.ColorGray)

	t.TextInput.SetTitle("titulo input").SetBorder(true)

	//handlers???

	t.Grid = tview.NewGrid().
		SetRows(0, 1).
		SetColumns(30, 0).
		SetBorders(false).
		AddItem(t.TextInput, 0, 1, 1, 1, 0, 0, false).
		AddItem(t.FooterText, 1, 0, 1, 2, 0, 0, false)

	return &t
}

func (tui *TUI) Start() error {
	return tui.App.SetRoot(tui.Grid, true).EnableMouse(true).Run()
}
