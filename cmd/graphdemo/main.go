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

	n = graph.NewGraphNode(nil)
	c := widget.NewVBox(
		widget.NewLabel("Fancy node!"),
		widget.NewButton("Up", func() {
			n.LogicalPosition = n.LogicalPosition.Add(fyne.Position{X: 0, Y: -10})
			n.Refresh()
		}),
		widget.NewButton("Down", func() {
			n.LogicalPosition = n.LogicalPosition.Add(fyne.Position{X: 0, Y: 10})
			n.Refresh()
		}),
		widget.NewHBox(
			widget.NewButton("Left", func() {
				n.LogicalPosition = n.LogicalPosition.Add(fyne.Position{X: -10, Y: 0})
				n.Refresh()
			}),
			widget.NewButton("Right", func() {
				n.LogicalPosition = n.LogicalPosition.Add(fyne.Position{X: 10, Y: 0})
				n.Refresh()
			}),
		),
	)
	n.InnerObject = c
	n.LogicalPosition = fyne.Position{300, 300}
	g.Nodes["node2"] = n

	w.SetContent(g)

	w.ShowAndRun()
}
