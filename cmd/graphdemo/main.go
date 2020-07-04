package main

import (
	"fmt"
	"image/color"

	"git.sr.ht/~charles/fynehax/graph"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Graph Demo")

	w.SetMaster()

	g := graph.NewGraph()

	l := widget.NewLabel("teeexxttt")
	n := graph.NewGraphNode(g, l)
	g.Nodes["node0"] = n
	n1 := n

	b := widget.NewButton("button", func() { fmt.Printf("tapped!\n") })
	n = graph.NewGraphNode(g, b)
	n.Move(fyne.Position{200, 200})
	g.Nodes["node1"] = n
	n2 := n

	n = graph.NewGraphNode(g, nil)
	c := widget.NewVBox(
		widget.NewLabel("Fancy node!"),
		widget.NewButton("Up", func() {
			n.Displace(fyne.Position{X: 0, Y: -10})
			n.Refresh()
		}),
		widget.NewButton("Down", func() {
			n.Displace(fyne.Position{X: 0, Y: 10})
			n.Refresh()
		}),
		widget.NewHBox(
			widget.NewButton("Left", func() {
				n.Displace(fyne.Position{X: -10, Y: 0})
				n.Refresh()
			}),
			widget.NewButton("Right", func() {
				n.Displace(fyne.Position{X: 10, Y: 0})
				n.Refresh()
			}),
		),
	)
	n.InnerObject = c
	n.Move(fyne.Position{300, 300})
	g.Nodes["node2"] = n
	n3 := n

	g.Edges["edge0"] = graph.NewGraphEdge(g, n1, n2)
	g.Edges["edge1"] = graph.NewGraphEdge(g, n3, n2)
	g.Edges["edge1"].EdgeColor = color.RGBA{255, 64, 64, 255}
	g.Edges["edge1"].Directed = true

	w.SetContent(g)

	w.ShowAndRun()
}
