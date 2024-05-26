package boxes

import (
	"fmt"
	"io"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

type Context struct {
	fg       termbox.Attribute
	bg       termbox.Attribute
	debugger io.Writer
}

type Constraints struct {
	minWidth int
	maxWidth int

	minHeight int
	maxHeight int
}

type RenderBox struct {
	Left   int
	Right  int
	Top    int
	Bottom int
	// Should this take in the click point?
	OnClick  func()
	Children []RenderBox
}

func (b *RenderBox) findHitBox(x, y int) *RenderBox {
	if x < b.Left || x > b.Right || y > b.Bottom || y < b.Top {
		DebugLog(fmt.Sprintf("hit (%d, %d) not within RenderBox", x, y))
		return nil
	}
	if b.OnClick == nil {
		DebugLog(fmt.Sprintf("hit (%d, %d) within RenderBox that does not have an OnClick", x, y))
	} else {
		DebugLog("hit within RenderBox that has an OnClick")
	}

	var children = b.Children
	if children == nil {
		if b.OnClick == nil {
			return nil
		}
		return b
	}

	for _, child := range children {
		var maybe = child.findHitBox(x, y)
		if maybe != nil && maybe.OnClick != nil {
			// TODO could multiple children contain the same point?
			DebugLog("child RenderBox contains hit and is Clickable")
			return maybe
		}
	}
	if b.OnClick == nil {
		return nil
	}
	DebugLog("found hit target")
	return b
}

type RenderJob struct {
	Width  int
	Height int
	Exec   func(int, int) RenderBox
}

type Widget interface {
	Render(Context, Constraints) RenderJob
}

type Clickable struct {
	Child   Widget
	OnClick func()
}

func (w Clickable) Render(ctx Context, cons Constraints) RenderJob {
	DebugLog("Laying out a clickable...")
	var childJob = w.Child.Render(ctx, cons)

	return RenderJob{
		Width:  childJob.Width,
		Height: childJob.Height,
		Exec: func(x, y int) RenderBox {
			DebugLog("rendering a clickable...")
			var childBox = childJob.Exec(x, y)
			// width = 3
			// 123
			// ***
			// * *
			// ***
			return RenderBox{
				Left:     x,
				Top:      y,
				Right:    x + childJob.Width - 1,
				Bottom:   y + childJob.Height - 1,
				OnClick:  w.OnClick,
				Children: []RenderBox{childBox},
			}
		},
	}
}

// A text widget.
type Text struct {
	Msg string
}

func (w Text) Render(ctx Context, cons Constraints) RenderJob {
	DebugLog("Laying out a Text...")
	// TODO
	const height = 1
	var width int = 0
	// TODO layout
	for _, c := range w.Msg {
		width += runewidth.RuneWidth(c)
	}

	return RenderJob{
		Width:  width,
		Height: height,
		Exec: func(x, y int) RenderBox {
			DebugLog("Rendering a Text...")
			var renderX = x
			for _, c := range w.Msg {
				termbox.SetCell(renderX, y, c, ctx.fg, ctx.bg)
				renderX += runewidth.RuneWidth(c)
			}

			return RenderBox{
				Left:   x,
				Right:  x + width - 1,
				Top:    y,
				Bottom: y + height - 1,
			}
		},
	}
}
