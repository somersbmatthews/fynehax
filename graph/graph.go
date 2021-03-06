package graph

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type graphRenderer struct {
	graph *GraphWidget
}

type GraphWidget struct {
	widget.BaseWidget

	Offset fyne.Position

	// DesiredSize specifies the size which the graph widget should take
	// up, defaults to 800 x 600
	DesiredSize fyne.Size

	Nodes map[string]*GraphNode
	Edges map[string]*GraphEdge
}

func (r *graphRenderer) MinSize() fyne.Size {
	return r.graph.DesiredSize
}

func (r *graphRenderer) Layout(size fyne.Size) {
}

func (r *graphRenderer) ApplyTheme(size fyne.Size) {
}

func (r *graphRenderer) Refresh() {
	for _, e := range r.graph.Edges {
		e.Refresh()
	}
	for _, n := range r.graph.Nodes {
		n.Refresh()
	}
}

func (r *graphRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *graphRenderer) Destroy() {
}

func (r *graphRenderer) Objects() []fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, 0)
	for _, n := range r.graph.Nodes {
		obj = append(obj, n)
	}
	for _, e := range r.graph.Edges {
		obj = append(obj, e)
	}

	return obj
}

func (g *GraphWidget) CreateRenderer() fyne.WidgetRenderer {
	r := graphRenderer{
		graph: g,
	}

	return &r
}

func (g *GraphWidget) Cursor() desktop.Cursor {
	return desktop.DefaultCursor
}

func (g *GraphWidget) DragEnd() {
	g.Refresh()
}

func (g *GraphWidget) Dragged(event *fyne.DragEvent) {
	delta := fyne.Position{X: event.DraggedX, Y: event.DraggedY}
	for _, n := range g.Nodes {
		n.Displace(delta)
	}
	g.Refresh()
}

func (g *GraphWidget) MouseIn(event *desktop.MouseEvent) {
}

func (g *GraphWidget) MouseOut() {
}

func (g *GraphWidget) MouseMoved(event *desktop.MouseEvent) {
}

func NewGraph() *GraphWidget {
	g := &GraphWidget{
		DesiredSize: fyne.Size{Width: 800, Height: 600},
		Offset:      fyne.Position{0, 0},
		Nodes:       map[string]*GraphNode{},
		Edges:       map[string]*GraphEdge{},
	}

	g.ExtendBaseWidget(g)

	return g
}

func (g *GraphWidget) GetEdges(n *GraphNode) []*GraphEdge {
	edges := []*GraphEdge{}

	for _, e := range g.Edges {
		if e.Origin == n {
			edges = append(edges, e)
		} else if e.Target == n {
			edges = append(edges, e)
		}
	}

	return edges
}
