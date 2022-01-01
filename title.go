package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type BoxSize int

const (
	Medium BoxSize = iota
	Small
	Large
)

//New widget behavior
type MyTitle struct {
	widget.BaseWidget

	text     string
	TextSize int
	Area     BoxSize
	// OnTapped func() `json:"-"`
}

func NewTitle(text string) *MyTitle {
	title := &MyTitle{
		text: text,
		// OnTapped: tapped,
	}
	title.ExtendBaseWidget(title)
	return title

}

//New widget rendering
type titleRenderer struct {
	widget *MyTitle

	label *canvas.Text
	bg    *canvas.Rectangle
	area  BoxSize
}

func (t *MyTitle) CreateRenderer() fyne.WidgetRenderer {
	return myWidget(t.text, t.TextSize, t.Area)
}

func myWidget(text string, textsize int, area BoxSize) *titleRenderer {
	label := canvas.NewText(text, theme.ForegroundColor())
	label.TextSize = float32(textsize)

	colour := &color.NRGBA{255, 255, 255, 250}
	label.Color = colour

	return &titleRenderer{
		bg:    canvas.NewRectangle(theme.BackgroundColor()),
		label: label,
		area:  area,
	}
}

func (t *titleRenderer) Layout(size fyne.Size) {
	//resizing of the box is handled in MinSize()
	si := fyne.MeasureText(t.label.Text, t.label.TextSize, t.label.TextStyle)

	// center text
	t.label.Move(fyne.Position{X: (size.Width - si.Width) / 2, Y: (size.Height - si.Height) / 2})

	t.bg.Resize(size)
}

func (t *titleRenderer) MinSize() (size fyne.Size) {
	x := fyne.MeasureText(t.label.Text, t.label.TextSize, t.label.TextStyle).Height

	switch t.area {
	case Small:
		//x is good
	case Medium:
		x += x
	case Large:
		x = x * 3
	}
	return fyne.Size{Width: size.Width, Height: x}
}

//Important: sequence of drawing!
func (t *titleRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.bg, t.label}
}

func (t *titleRenderer) Refresh() {
}

func (t *titleRenderer) Destroy() {

}
