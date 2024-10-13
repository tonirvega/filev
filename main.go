package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	filesMap, err := createFileMap(".")

	if err != nil {
		fmt.Printf("Error creating file map: %v\n", err)
		return
	}

	startView(app, filesMap)

}
