package boxes

import (
	termbox "github.com/nsf/termbox-go"
)

// A widget that provides formatting to its child.
type Container struct {
	Child Widget
}

const topRightCorner = '\u2510'
const topLeftCorner = '\u250c'
const bottomLeftCorner = '\u2514'
const bottomRightCorner = '\u2518'
const horizontalBar = '\u2500'
const verticalBar = '\u2502'

func (w Container) render(ctx Context, cons Constraints) RenderJob {
	const paddingX = 1
	const paddingY = 1
	// TODO some code refactoring before non-1 renders correctly
	const borderThickness = 1

	var childJob = w.Child.render(
		ctx,
		Constraints{
			// Should these mins be less padding and border?
			minWidth:  cons.minWidth,
			minHeight: cons.minHeight,

			maxWidth:  cons.maxWidth - (borderThickness-paddingX)*2,
			maxHeight: cons.maxWidth - (borderThickness-paddingY)*2,
		},
	)

	// TODO check for overflow
	var boxWidth = childJob.width + (borderThickness+paddingX)*2
	var boxHeight = childJob.height + (borderThickness+paddingY)*2

	return RenderJob{
		width:  boxWidth,
		height: boxHeight,
		exec: func(x, y int) {
			childJob.exec(
				x+borderThickness+paddingX,
				y+borderThickness+paddingY,
			)

			// Corners
			termbox.SetCell(x, y, topLeftCorner, ctx.fg, ctx.bg)
			termbox.SetCell(x+boxWidth-borderThickness, y, topRightCorner, ctx.fg, ctx.bg)
			termbox.SetCell(x, y+boxHeight-borderThickness, bottomLeftCorner, ctx.fg, ctx.bg)
			termbox.SetCell(x+boxWidth-borderThickness, y+boxHeight-borderThickness, bottomRightCorner, ctx.fg, ctx.bg)

			for xIter := x + 1; xIter < x+boxWidth-1; xIter++ {
				// Top border
				termbox.SetCell(xIter, y, horizontalBar, ctx.fg, ctx.bg)
				// Bottom border
				termbox.SetCell(xIter, y+boxHeight-borderThickness, horizontalBar, ctx.fg, ctx.bg)
			}

			for yIter := y + 1; yIter < y+boxHeight-1; yIter++ {
				// Left border (sans top & bottom borders)
				termbox.SetCell(x, yIter, verticalBar, ctx.fg, ctx.bg)
				// Right border (sans top & bottom border)
				termbox.SetCell(x+boxWidth-borderThickness, yIter, verticalBar, ctx.fg, ctx.bg)
			}
		},
	}
}
