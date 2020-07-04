// Package arrowhead implements an arrowhead canvas object.
package arrowhead

import (
	"image/color"
	"math"

	"git.sr.ht/~charles/fynehax/geometry/r2"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

const (
	defaultTheta       float32 = 10
	defaultStrokeWidth float32 = 2
	defaultLength      int     = 15
)

// Arrowhead defines a canvas object which renders an arrow pointing in
// a particular direction.
//
//
//             Left
//               \
//                \  Length
//           Theta \
//     Base ------- + Tip
//                 /
//                /
//               /
//             Right
//
type Arrowhead struct {
	// Base is used to define the "base" of the arrow, which thus defines
	// the direction which the arrow faces.
	Base fyne.Position

	// Tip is the point at which the tip of the arrow will be placed.
	Tip fyne.Position

	// StrokeWidth is the width of the arrowhead lines
	StrokeWidth float32

	// StrokeColor is the color of the arrowhead
	StrokeColor color.Color

	// Theta is the angle between the two "tails" that intersect at the
	// tip.
	Theta float32

	// Length is the length of the two "tails" that intersect at the tip.
	Length int

	central *canvas.Line
	left    *canvas.Line
	right   *canvas.Line
	visible bool
}

func MakeArrowhead(base, tip fyne.Position) *Arrowhead {
	return &Arrowhead{
		Base:        base,
		Tip:         tip,
		StrokeWidth: defaultStrokeWidth,
		StrokeColor: theme.TextColor(),
		Theta:       defaultTheta,
		Length:      defaultLength,
		central:     canvas.NewLine(theme.TextColor()),
		left:        canvas.NewLine(theme.TextColor()),
		right:       canvas.NewLine(theme.TextColor()),
		visible:     true,
	}
}

func (a *Arrowhead) Refresh() {
	a.central.StrokeWidth = a.StrokeWidth
	a.left.StrokeWidth = a.StrokeWidth
	a.right.StrokeWidth = a.StrokeWidth

	a.central.StrokeColor = a.StrokeColor
	a.left.StrokeColor = a.StrokeColor
	a.right.StrokeColor = a.StrokeColor

	a.central.Position1 = a.Tip
	a.central.Position2 = a.Base

	a.left.Position1 = a.Tip
	a.left.Position2 = a.LeftPoint()

	a.right.Position1 = a.Tip
	a.right.Position2 = a.RightPoint()

	if a.visible {
		a.central.Show()
		a.left.Show()
		a.right.Show()
	} else {
		a.central.Hide()
		a.left.Hide()
		a.right.Hide()
	}

	canvas.Refresh(a.central)
	canvas.Refresh(a.left)
	canvas.Refresh(a.right)

}

func (a *Arrowhead) LeftPoint() fyne.Position {
	return a.Tip.Add(fyne.Position{
		X: int(float64(a.Length) * math.Cos(float64(a.Theta))),
		Y: int(float64(a.Length) * math.Sin(float64(a.Theta))),
	})
}

func (a *Arrowhead) RightPoint() fyne.Position {
	return a.Tip.Add(fyne.Position{
		X: int(float64(a.Length) * math.Cos(float64(a.Theta))),
		Y: int(float64(a.Length) * -1.0 * math.Sin(float64(a.Theta))),
	})
}

func (a *Arrowhead) Size() fyne.Size {
	lp := a.LeftPoint()
	rp := a.RightPoint()
	points := []r2.Vec2{
		r2.Vec2{X: float64(a.Tip.X), Y: float64(a.Tip.Y)},
		r2.Vec2{X: float64(a.Base.X), Y: float64(a.Base.Y)},
		r2.Vec2{X: float64(lp.X), Y: float64(lp.Y)},
		r2.Vec2{X: float64(rp.X), Y: float64(rp.Y)},
	}

	bounding := r2.BoundingBox(points)
	return fyne.Size{
		Width:  int(bounding.Width()),
		Height: int(bounding.Height()),
	}
}

func (a *Arrowhead) Resize(s fyne.Size) {
	l := r2.V2(float64(s.Width), float64(s.Height))
	a.Length = int(l.Length())

	tip := r2.V2(float64(a.Tip.X), float64(a.Tip.Y))
	base := r2.V2(float64(a.Base.X), float64(a.Base.Y))
	v := tip.Add(base.Scale(-1))
	v = v.ScaleToLength(l.Length())
	base = tip.Add(v)

	a.Base = fyne.Position{X: int(base.X), Y: int(base.Y)}
}

func (a *Arrowhead) Move(p fyne.Position) {
	a.Tip = p

	tip := r2.V2(float64(a.Tip.X), float64(a.Tip.Y))
	base := r2.V2(float64(a.Base.X), float64(a.Base.Y))
	v := tip.Add(base.Scale(-1))
	base = tip.Add(v)

	a.Base = fyne.Position{X: int(base.X), Y: int(base.Y)}
}

func (a *Arrowhead) MinSize() fyne.Size {
	return a.Size()
}

func (a *Arrowhead) Visible() bool {
	return a.visible
}

func (a *Arrowhead) Show() {
	a.visible = true
}

func (a *Arrowhead) Hide() {
	a.visible = false
}

func (a *Arrowhead) Position() fyne.Position {
	return a.Tip
}

// temporary hack, don't do this
func (a *Arrowhead) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{
		a.central,
		a.left,
		a.right,
	}
}
