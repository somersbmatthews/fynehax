package graph

import (
	"image/color"

	"git.sr.ht/~charles/fynehax/arrowhead"
	"git.sr.ht/~charles/fynehax/geometry/r2"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type graphEdgeRenderer struct {
	edge  *GraphEdge
	line  *canvas.Line
	arrow *arrowhead.Arrowhead
}

type GraphEdge struct {
	widget.BaseWidget

	Graph *GraphWidget

	EdgeColor color.Color

	Width float32

	Origin *GraphNode
	Target *GraphNode

	Directed bool
}

func (r *graphEdgeRenderer) MinSize() fyne.Size {
	xdelta := r.edge.Origin.Position().X - r.edge.Target.Position().X
	if xdelta < 0 {
		xdelta *= -1
	}

	ydelta := r.edge.Origin.Position().Y - r.edge.Target.Position().Y
	if ydelta < 0 {
		ydelta *= -1
	}

	return fyne.Size{Width: xdelta, Height: ydelta}
}

func (r *graphEdgeRenderer) Layout(size fyne.Size) {
}

func (r *graphEdgeRenderer) ApplyTheme(size fyne.Size) {
}

func (r *graphEdgeRenderer) Refresh() {
	l := r.edge.R2Line()
	b1 := r.edge.Origin.R2Box()
	b2 := r.edge.Target.R2Box()

	p1, _ := b1.Intersect(l)
	p2, _ := b2.Intersect(l)

	r.line.Position1 = fyne.Position{
		X: int(p1.X),
		Y: int(p1.Y),
	}

	r.line.Position2 = fyne.Position{
		X: int(p2.X),
		Y: int(p2.Y),
	}

	r.line.StrokeColor = r.edge.EdgeColor
	r.line.StrokeWidth = r.edge.Width

	if r.edge.Directed {
		r.arrow.Show()
		r.arrow.Tip = r.line.Position1
		r.arrow.Base = r.line.Position2
		r.arrow.StrokeColor = r.edge.EdgeColor
		r.arrow.StrokeWidth = r.edge.Width
	} else {
		r.arrow.Hide()
	}

	canvas.Refresh(r.line)
	canvas.Refresh(r.arrow)
	r.arrow.Refresh()
}

func (r *graphEdgeRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *graphEdgeRenderer) Destroy() {
}

func (r *graphEdgeRenderer) Objects() []fyne.CanvasObject {
	// obj := []fyne.CanvasObject{
	//         r.line,
	//         // r.arrow,
	// }

	// XXX: temporary hack because otherwise I can't get my canvas object
	// to show up???
	obj := r.arrow.Objects()
	obj = append(obj, r.line)
	return obj
}

func (e *GraphEdge) CreateRenderer() fyne.WidgetRenderer {
	r := graphEdgeRenderer{
		edge:  e,
		line:  canvas.NewLine(e.EdgeColor),
		arrow: arrowhead.MakeArrowhead(fyne.Position{0, 0}, fyne.Position{0, 0}),
	}

	(&r).Refresh()

	return &r
}

func (e *GraphEdge) R2Line() r2.Line {
	return r2.MakeLineFromEndpoints(e.Origin.R2Center(), e.Target.R2Center())
}

func NewGraphEdge(g *GraphWidget, v, u *GraphNode) *GraphEdge {
	e := &GraphEdge{
		Graph:     g,
		EdgeColor: theme.TextColor(),
		Width:     2,
		Origin:    v,
		Target:    u,
		Directed:  false,
	}

	e.ExtendBaseWidget(e)

	return e
}
