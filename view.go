package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func startView(app *tview.Application, filesMap map[string]string) bool {

	table := tview.
		NewTable()

	table.Clear()
	table.SetSelectable(true, false)
	row := 0

	inputField := configureInputField(table, row, filesMap, app)

	configureTable(table, app, inputField)

	for path := range filesMap {
		table.SetCell(row, 0, tview.NewTableCell(path).SetTextColor(tcell.ColorWhiteSmoke).SetSelectable(true))
		row++
	}

	form := tview.NewForm().
		AddTextView(
			"fileV",
			"v1.0.0",
			50,
			50,
			true,
			false,
		)

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 3, 3, false).
		AddItem(inputField, 3, 1, false).
		AddItem(table, 0, 100, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

	return false
}

func configureTable(table *tview.Table, app *tview.Application, inputField *tview.InputField) {

	table.SetSelectionChangedFunc(func(row, column int) {

		for i := 0; i < table.GetRowCount(); i++ {
			cell := table.GetCell(i, 0)
			if i == row {
				cell.SetTextColor(tcell.ColorBlue).SetAttributes(tcell.AttrBold)
				cell.SetBackgroundColor(tcell.ColorWhiteSmoke)
			} else {
				cell.SetTextColor(tcell.ColorWhiteSmoke).SetAttributes(tcell.AttrNone)
				cell.SetBackgroundColor(tcell.ColorBlack)
			}
		}
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// if user press enter key or / slash key
		if event.Key() == tcell.KeyEnter {

			path := table.GetCell(table.GetSelection()).Text

			app.Suspend(func() {
				openWithVim(path)
			})

		}

		if event.Key() == tcell.KeyEsc || event.Rune() == '/' {
			app.SetFocus(inputField)
			inputField.SetText("")
		}

		return event

	})

	table.SetBorder(true).SetBorderColor(tcell.ColorWhiteSmoke)
}

func configureInputField(table *tview.Table, row int, filesMap map[string]string, app *tview.Application) *tview.InputField {

	inputField := tview.
		NewInputField().
		SetChangedFunc(func(text string) {

			if !strings.HasPrefix(text, ":") {

				table.Clear()
				row = 0

				for path := range filesMap {

					if matchesFilter(path, text, filesMap[path]) {

						table.SetCell(
							row,
							0,
							tview.
								NewTableCell(path).
								SetTextColor(tcell.ColorWhiteSmoke).
								SetSelectable(true),
						)

						row++
					}
				}
			}

		})

	inputField.SetDoneFunc(func(key tcell.Key) {

		if key == tcell.KeyEnter {

			switch inputField.GetText() {

			case ":q":
				app.Stop()
				return
			case ":quit":
				app.Stop()
				return
			default:
			}

			app.SetFocus(table)

			table.Select(0, 0)

		}
	})

	inputField.
		SetBorder(true).
		SetBorderColor(tcell.ColorWhiteSmoke)

	return inputField
}
