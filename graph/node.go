package graph

import (
	"image/color"

	"git.sr.ht/~charles/fynehax/geometry/r2"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	// default inner size
	defaultWidth  int = 50
	defaultHeight int = 50

	// default padding around the inner object in a node
	defaultPadding int = 10
)

type graphNodeRenderer struct {
	node   *GraphNode
	handle *canvas.Line
	box    *canvas.Rectangle
}

// GraphNode represents a node in the graph widget. It contains an inner
// widget, and also draws a border, and a "handle" that can be used to drag it
// around.
type GraphNode struct {
	widget.BaseWidget

	Graph *GraphWidget

	// InnerSize stores size that the inner object should have, may not
	// be respected if not large enough for the object.
	InnerSize fyne.Size

	// InnerObject is the canvas object that should be drawn inside of
	// the graph node.
	InnerObject fyne.CanvasObject

	// Padding is the distance between the inner object's drawing area
	// and the box.
	Padding int

	// BoxStrokeWidth is the stroke width of the box which delineates the
	// node. Defaults to 1.
	BoxStrokeWidth float32

	// BoxFill is the fill color of the node, the inner object will be
	// drawn on top of this. Defaults to the theme.BackgroundColor().
	BoxFillColor color.Color

	// BoxStrokeColor is the stroke color of the node rectangle. Defaults
	// to theme.TextColor().
	BoxStrokeColor color.Color

	// HandleColor is the color of node handle.
	HandleColor color.Color

	// HandleStrokeWidth is the stroke width of the node handle, defaults
	// to 3.
	HandleStroke float32
}

func (r *graphNodeRenderer) MinSize() fyne.Size {
	// space for the inner widget, plus padding on all sides.
	inner := r.node.effectiveInnerSize()
	return fyne.Size{
		Width:  inner.Width + 2*r.node.Padding,
		Height: inner.Height + 2*r.node.Padding,
	}
}

func (r *graphNodeRenderer) Layout(size fyne.Size) {
	r.node.Resize(r.MinSize())

	r.node.InnerObject.Move(r.node.innerPos())
	r.node.InnerObject.Resize(r.node.effectiveInnerSize())

	r.box.Resize(r.MinSize())

	canvas.Refresh(r.node.InnerObject)
}

func (r *graphNodeRenderer) ApplyTheme(size fyne.Size) {
}

func (r *graphNodeRenderer) Refresh() {
	// move and size the inner object appropriately
	r.node.InnerObject.Move(r.node.innerPos())
	r.node.InnerObject.Resize(r.node.effectiveInnerSize())

	// move the box and update it's colors
	r.box.StrokeWidth = r.node.BoxStrokeWidth
	r.box.FillColor = r.node.BoxFillColor
	r.box.StrokeColor = r.node.BoxStrokeColor
	r.box.Resize(r.MinSize())

	// calculate the handle positions
	r.handle.Position1 = fyne.Position{
		X: r.node.Padding,
		Y: r.node.Padding / 2,
	}

	r.handle.Position2 = fyne.Position{
		X: r.node.effectiveInnerSize().Width + r.node.Padding,
		Y: r.node.Padding / 2,
	}

	r.handle.StrokeWidth = r.node.HandleStroke
	r.handle.StrokeColor = r.node.HandleColor

	for _, e := range r.node.Graph.GetEdges(r.node) {
		e.Refresh()
	}

	canvas.Refresh(r.box)
	canvas.Refresh(r.handle)
	canvas.Refresh(r.node.InnerObject)
}

func (r *graphNodeRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *graphNodeRenderer) Destroy() {
}

func (r *graphNodeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		r.box,
		r.handle,
		r.node.InnerObject,
	}
}

func (n *GraphNode) CreateRenderer() fyne.WidgetRenderer {
	r := graphNodeRenderer{
		node:   n,
		handle: canvas.NewLine(n.HandleColor),
		box:    canvas.NewRectangle(n.BoxStrokeColor),
	}

	r.handle.StrokeWidth = n.HandleStroke
	r.box.StrokeWidth = n.BoxStrokeWidth
	r.box.FillColor = n.BoxFillColor

	(&r).Refresh()

	return &r
}

func NewGraphNode(g *GraphWidget, obj fyne.CanvasObject) *GraphNode {
	w := &GraphNode{
		Graph:          g,
		InnerSize:      fyne.Size{Width: defaultWidth, Height: defaultHeight},
		InnerObject:    obj,
		Padding:        defaultPadding,
		BoxStrokeWidth: 1,
		BoxFillColor:   theme.BackgroundColor(),
		BoxStrokeColor: theme.TextColor(),
		HandleColor:    theme.TextColor(),
		HandleStroke:   3,
	}

	w.ExtendBaseWidget(w)

	return w
}

func (n *GraphNode) innerPos() fyne.Position {
	return fyne.Position{
		X: n.Padding,
		Y: n.Padding,
	}
}

func (n *GraphNode) effectiveInnerSize() fyne.Size {
	return n.InnerSize.Max(n.InnerObject.MinSize())
}

func (n *GraphNode) Cursor() desktop.Cursor {
	return desktop.DefaultCursor
}

func (n *GraphNode) DragEnd() {
	n.Refresh()
}

func (n *GraphNode) Dragged(event *fyne.DragEvent) {
	delta := fyne.Position{X: event.DraggedX, Y: event.DraggedY}
	n.Displace(delta)
	n.Refresh()
}

func (n *GraphNode) MouseIn(event *desktop.MouseEvent) {
	n.HandleColor = theme.FocusColor()
	n.Refresh()
}

func (n *GraphNode) MouseOut() {
	n.HandleColor = theme.TextColor()
	n.Refresh()
}

func (n *GraphNode) MouseMoved(event *desktop.MouseEvent) {
}

func (n *GraphNode) Displace(delta fyne.Position) {
	n.Move(n.Position().Add(delta))
}

func (n *GraphNode) R2Position() r2.Vec2 {
	return r2.V2(float64(n.Position().X), float64(n.Position().Y))
}

func (n *GraphNode) R2Box() r2.Box {
	inner := n.effectiveInnerSize()
	s := r2.V2(
		float64(inner.Width+2*n.Padding),
		float64(inner.Height+2*n.Padding),
	)

	return r2.MakeBox(n.R2Position(), s)
}

func (n *GraphNode) R2Center() r2.Vec2 {
	return n.R2Box().Center()
}

func (n *GraphNode) Center() fyne.Position {
	return fyne.Position{int(n.R2Center().X), int(n.R2Center().Y)}
}
