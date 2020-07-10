package main

import (
	"fmt"
	"image/color"
	"time"

	"git.sr.ht/~charles/fynehax/graph"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var forceticks int = 0

var globalgraph *graph.GraphWidget

func forceanim() {

	// XXX: very naughty -- accesses shared memory in potentially unsafe
	// ways, this almost certainly has race conditions... don't do this!

	for {
		if forceticks > 0 {
			globalgraph.StepForceLayout(300)
			globalgraph.Refresh()
			forceticks--
			fmt.Printf("forceticks=%d\n", forceticks)
		}

		time.Sleep(time.Millisecond * (1000 / 30))
	}
}

func main() {
	app := app.New()
	w := app.NewWindow("Graph Demo")

	w.SetMaster()

	g := graph.NewGraph()

	go forceanim()

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

	n = graph.NewGraphNode(g, widget.NewButton("force layout step", func() {
		g.StepForceLayout(300)
		g.Refresh()
	}))
	n.Move(fyne.Position{400, 200})
	g.Nodes["node4"] = n
	n4 := n

	n = graph.NewGraphNode(g, widget.NewButton("auto layout", func() {
		forceticks += 100
		g.Refresh()
	}))
	n.Move(fyne.Position{400, 500})
	g.Nodes["node5"] = n
	n5 := n

	globalgraph = g

	g.Edges["edge0"] = graph.NewGraphEdge(g, n1, n2)
	g.Edges["edge1"] = graph.NewGraphEdge(g, n3, n2)
	g.Edges["edge1"].EdgeColor = color.RGBA{255, 64, 64, 255}
	g.Edges["edge1"].Directed = true
	g.Edges["edge2"] = graph.NewGraphEdge(g, n1, n4)
	g.Edges["edge3"] = graph.NewGraphEdge(g, n3, n4)
	g.Edges["edge4"] = graph.NewGraphEdge(g, n5, n4)

	w.SetContent(g)

	w.ShowAndRun()
}
