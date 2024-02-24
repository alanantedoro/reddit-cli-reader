package tui

import "github.com/rivo/tview"

type TUI struct {
	App        *tview.Application
	Grid       *tview.Grid
	FooterText *tview.TextView
	TextInput  *tview.InputField
}
