package main

import (
	"fmt"
	"image/color"
	"strconv"

	"git.sr.ht/~charles/fynehax/viewport"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()
	w := app.NewWindow("Viewport Demo")

	w.SetMaster()

	stepSize := 1.0

	stepSizeEntry := widget.NewEntry()
	stepSizeEntry.OnChanged = func(text string) {
		var err error
		stepSize, err = strconv.ParseFloat(text, 64)
		if err != nil {
			stepSizeEntry.SetText(fmt.Sprintf("%f", stepSize))
		}
	}
	stepSizeEntry.SetText("1.0")

	vp := viewport.NewViewportWidget(800, 600)

	vp.Objects = append(vp.Objects, viewport.ViewportLine{
		X1:          20,
		Y1:          20,
		X2:          200,
		Y2:          400,
		StrokeColor: color.RGBA{255, 255, 255, 255},
		StrokeWidth: 1,
	})

	w.SetContent(widget.NewHSplitContainer(
		vp,
		widget.NewVBox(
			stepSizeEntry,
			widget.NewButton("Pan Left", func() {
				fmt.Printf("vp.XOffset %v", vp.XOffset)
				vp.XOffset += vp.Zoom * stepSize
				fmt.Printf(" -> %v\n", vp.XOffset)
			}),
			widget.NewButton("Pan Right", func() {
				fmt.Printf("vp.XOffset %v", vp.XOffset)
				vp.XOffset -= vp.Zoom * stepSize
				fmt.Printf(" -> %v\n", vp.XOffset)
			}),
			widget.NewButton("Pan Up", func() {
				fmt.Printf("vp.YOffset %v", vp.YOffset)
				vp.YOffset += vp.Zoom * stepSize
				fmt.Printf(" -> %v\n", vp.YOffset)
			}),
			widget.NewButton("Pan Down", func() {
				fmt.Printf("vp.YOffset %v", vp.YOffset)
				vp.YOffset -= vp.Zoom * stepSize
				fmt.Printf(" -> %v\n", vp.YOffset)
			}),
			widget.NewButton("Zoom In", func() {
				fmt.Printf("vp.Zoom %v", vp.Zoom)
				vp.Zoom *= 1.1
				fmt.Printf(" -> %v\n", vp.Zoom)
			}),
			widget.NewButton("Zoom Out", func() {
				fmt.Printf("vp.Zoom %v", vp.Zoom)
				vp.Zoom *= 0.9
				fmt.Printf(" -> %v\n", vp.Zoom)
			}),
		),
	))

	w.ShowAndRun()

}
