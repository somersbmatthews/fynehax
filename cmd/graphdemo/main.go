package main

import (
	"fmt"

	"git.sr.ht/~charles/fynehax/graph"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	fmt.Println("vim-go")

	app := app.New()
	w := app.NewWindow("Viewport Demo")

	w.SetMaster()

	l := widget.NewLabel("teeexxttt")
	n := graph.NewGraphNode(l)

	w.SetContent(n)

	w.ShowAndRun()
}
