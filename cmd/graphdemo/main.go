package main

import (
	"fmt"

	"git.sr.ht/~charles/fynehax/graph"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	fmt.Println("vim-go")

	app := app.New()
	w := app.NewWindow("Viewport Demo")

	w.SetMaster()

	g := graph.NewGraph()

	l := widget.NewLabel("teeexxttt")
	n := graph.NewGraphNode(l)
	g.Nodes["node0"] = n

	b := widget.NewButton("button", func() { fmt.Printf("tapped!\n") })
	n = graph.NewGraphNode(b)
	n.LogicalPosition = fyne.Position{200, 200}
	g.Nodes["node1"] = n

	w.SetContent(g)

	w.ShowAndRun()
}
