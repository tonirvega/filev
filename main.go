package main

import (
	"fmt"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	fMap, err := createFileMap(".", []string{".git", "node_modules", "vendor"})

	if err != nil {
		fmt.Printf("Error creating file map: %v\n", err)
		return
	}

	filesMap = fMap

	startView(app, filesMap)

}
