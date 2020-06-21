package viewport

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type viewportRenderer struct {
	viewport *ViewportWidget
}

func (r *viewportRenderer) MinSize() fyne.Size {
	return fyne.Size{r.viewport.Width, r.viewport.Height}
}

func (r *viewportRenderer) Layout(size fyne.Size) {
}

func (r *viewportRenderer) ApplyTheme(size fyne.Size) {
}

func (r *viewportRenderer) Refresh() {
}

func (r *viewportRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *viewportRenderer) Destroy() {
}

func (r *viewportRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0)
	for _, viewportObj := range r.viewport.Objects {
		objects = append(objects, viewportObj.CanvasObjects(r.viewport)...)
	}
	return objects
}

type ViewportWidget struct {
	widget.BaseWidget
	Width   int
	Height  int
	Zoom    float64
	XOffset float64
	YOffset float64
	Objects []ViewportObject
}

func (w *ViewportWidget) Tapped(ev *fyne.PointEvent) {
}

func (w *ViewportWidget) TappedSecondary(ev *fyne.PointEvent) {
}

func (w *ViewportWidget) CreateRenderer() fyne.WidgetRenderer {
	r := viewportRenderer{
		viewport: w,
	}

	r.Refresh()
	return &r
}

func NewViewportWidget(width, height int) *ViewportWidget {
	vp := &ViewportWidget{
		Width:   width,
		Height:  height,
		Zoom:    1.0,
		XOffset: 0,
		YOffset: 0,
		Objects: make([]ViewportObject, 0),
	}

	vp.ExtendBaseWidget(vp)
	return vp
}

type ViewportObject interface {
	CanvasObjects(viewport *ViewportWidget) []fyne.CanvasObject
}

type ViewportLine struct {
	obj         *canvas.Line
	X1          float64
	Y1          float64
	X2          float64
	Y2          float64
	StrokeColor color.Color
	StrokeWidth float64
}

func setLineEndpoints(l *canvas.Line, X1, Y1, X2, Y2 float64) {
	l.Move(fyne.NewPos(int(X1), int(Y1)))
	l.Resize(fyne.NewSize(int(X2)-int(X1), int(Y2)-int(Y1)))
}

func (l ViewportLine) CanvasObjects(viewport *ViewportWidget) []fyne.CanvasObject {
	if l.obj == nil {
		l.obj = canvas.NewLine(l.StrokeColor)
	}

	setLineEndpoints(l.obj,
		(l.X1+viewport.XOffset)*viewport.Zoom,
		(l.Y1+viewport.YOffset)*viewport.Zoom,
		(l.X2+viewport.XOffset)*viewport.Zoom,
		(l.Y2+viewport.YOffset)*viewport.Zoom,
	)
	l.obj.StrokeWidth = float32(l.StrokeWidth)
	l.obj.Hidden = false

	return []fyne.CanvasObject{l.obj}
}
