package boxes

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

type Context struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

type Constraints struct {
	minWidth int
	maxWidth int

	minHeight int
	maxHeight int
}

type RenderBox struct {
	Left int
	Right int
	Top int
	Bottom int
	OnClick func()
}

type RenderJob struct {
	width  int
	height int
	exec   func(int, int) RenderBox
}

type Widget interface {
	render(Context, Constraints) RenderJob
}

type Clickable struct {
	Child Widget
	OnClick func()
}

func (w Clickable) render(ctx Context, cons Constraints) RenderJob {
	var childJob = w.Child.render(ctx, cons)

	return RenderJob{
		width: childJob.width,
		height: childJob.height,
		exec: func(x, y int) RenderBox {
			childJob.exec(x, y)
			return RenderBox{
				Left: x,
				Top: y,
				Right: x + childJob.width, // TODO is this right?
				Bottom: y + childJob.height,
				OnClick: w.OnClick,
			}
		},
	}
}

// A text widget.
type Text struct {
	Msg string
}

func (w Text) render(ctx Context, cons Constraints) RenderJob {
	var width int = 0
	// TODO layout
	for _, c := range w.Msg {
		width += runewidth.RuneWidth(c)
	}

	return RenderJob{
		width:  width,
		height: 1, // TODO
		exec: func(x, y int) {
			for _, c := range w.Msg {
				termbox.SetCell(x, y, c, ctx.fg, ctx.bg)
				x += runewidth.RuneWidth(c)
			}
		},
	}
}
