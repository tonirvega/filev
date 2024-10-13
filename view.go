package main

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	flex      *tview.Flex
	filesMap  map[string]string
	filesSize map[string]int64 = make(map[string]int64)
)

func startView(app *tview.Application, filesMap map[string]string) bool {

	table := tview.
		NewTable()

	table.Clear()
	table.SetSelectable(
		true,
		false,
	)
	table.SetBorder(true)
	table.SetTitle(": files :")

	row := 0

	inputField := configureInputField(table, row, filesMap, app)

	configureTable(table, app, inputField)

	for path := range filesMap {
		table.SetCell(
			row,
			0,
			tview.
				NewTableCell(path).
				SetTextColor(tcell.ColorWhiteSmoke).
				SetSelectable(true).
				SetAlign(tview.AlignLeft).
				SetExpansion(1),
		)
		row++
	}

	form := tview.NewForm().
		AddTextView(
			"",
			strings.Join(logoInfo, "\n"),
			100,
			6,
			true,
			false,
		)

	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorYellowGreen)
	flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(form, 10, 10, false).
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
				cell.SetTextColor(tcell.ColorBlue).
					SetAttributes(tcell.AttrBold).
					SetBackgroundColor(tcell.ColorWhiteSmoke).
					SetExpansion(1).
					SetAlign(tview.AlignLeft)
			} else {
				cell.SetTextColor(tcell.ColorWhiteSmoke).
					SetAttributes(tcell.AttrNone).
					SetBackgroundColor(tcell.ColorBlack).
					SetExpansion(1).
					SetAlign(tview.AlignLeft)

			}
		}
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Rune() == 'e' {

			path := table.GetCell(table.GetSelection()).Text

			app.Suspend(func() {

				openWithVim(path)

				body, err := os.ReadFile(path)

				if err != nil {

					panic(err)
				}

				filesMap[path] = string(body)

			})

		}

		if event.Rune() == '/' || event.Rune() == ':' {
			flex.AddItem(inputField, 3, 1, false)
			app.SetFocus(inputField)
			inputField.SetText("")
		}

		if event.Key() == tcell.KeyESC || event.Key() == tcell.KeyEscape {
			flex.AddItem(inputField, 3, 1, false)
			app.SetFocus(inputField)
		}

		return event

	})

	table.SetFocusFunc(func() {
		table.Select(0, 0)
		table.SetBorderColor(tcell.ColorBlue)
	})

	table.SetBlurFunc(func() {
		table.SetBorderColor(tcell.ColorWhiteSmoke)
	})
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

	inputField.SetFieldStyle(
		tcell.StyleDefault.Background(
			tcell.ColorBlack,
		),
	)

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
			// hide input field
			flex.RemoveItem(inputField)
			app.SetFocus(table)

		} else if key == tcell.KeyESC || key == tcell.KeyEscape {
			inputField.SetText("")
			flex.RemoveItem(inputField)
			app.SetFocus(table)
		}

	})

	inputField.
		SetBorder(true).
		SetBorderColor(tcell.ColorWhiteSmoke)

	inputField.SetLabel("> ")

	return inputField
}
