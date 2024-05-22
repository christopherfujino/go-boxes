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
	// Should this take in the click point?
	OnClick func()
	Children []RenderBox
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
			// width = 3
			// 123
			// ***
			// * *
			// ***
			return RenderBox{
				Left: x,
				Top: y,
				Right: x + childJob.width - 1,
				Bottom: y + childJob.height - 1,
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
	// TODO
	const height = 1
	var width int = 0
	// TODO layout
	for _, c := range w.Msg {
		width += runewidth.RuneWidth(c)
	}

	return RenderJob{
		width:  width,
		height: height,
		exec: func(x, y int) RenderBox {
			var renderX = x
			for _, c := range w.Msg {
				termbox.SetCell(renderX, y, c, ctx.fg, ctx.bg)
				renderX += runewidth.RuneWidth(c)
			}

			return RenderBox{
				Left: x,
				Right: x + width - 1,
				Top: y,
				Bottom: y + height - 1,
			}
		},
	}
}
