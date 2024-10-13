package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func startView(app *tview.Application, filesMap map[string]string) bool {

	table := tview.
		NewTable()

	table.Clear()
	table.SetSelectable(true, false)
	row := 0

	table.SetSelectionChangedFunc(func(row, column int) {

		for i := 0; i < table.GetRowCount(); i++ {
			cell := table.GetCell(i, 0)
			if i == row {
				cell.SetTextColor(tcell.ColorBlue).SetAttributes(tcell.AttrBold)
			} else {
				cell.SetTextColor(tcell.ColorWhiteSmoke).SetAttributes(tcell.AttrNone)
			}
		}
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEnter {

			path := table.GetCell(table.GetSelection()).Text

			app.Suspend(func() {

				openWithVim(path)

			})

			return event
		}

		return event
	})

	for path := range filesMap {
		table.SetCell(row, 0, tview.NewTableCell(path).SetTextColor(tcell.ColorWhiteSmoke).SetSelectable(true))
		row++
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
	return false
}
