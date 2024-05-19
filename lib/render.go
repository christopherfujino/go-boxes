package main

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

type RenderContext struct {
	constraints Constraints

	fg termbox.Attribute
	bg termbox.Attribute
}

// With Flutter, this is min/max width/height
type Constraints struct {
	startingX int
	finalX    int

	startingY int
	finalY    int
}

type RenderJob struct {
	width  int
	height int
	exec   func(int, int)
}

type CenterWidget struct {
	child Widget
}

func (w CenterWidget) render(ctx RenderContext) RenderJob {
	var width = ctx.constraints.finalX - ctx.constraints.startingX
	var childJob = w.child.render(ctx)
	var leftPad = (width - childJob.width) / 2
	return RenderJob{
		width:  width,
		height: childJob.height,
		exec: func(x, y int) {
			childJob.exec(x+leftPad, y)
		},
	}
}

// TODO consider constraints
type Widget interface {
	render(RenderContext) RenderJob
}

type TextWidget struct {
	msg string
}

type ContainerWidget struct {
	child Widget
}

const topRightCorner = '\u2510'
const topLeftCorner = '\u250c'
const bottomLeftCorner = '\u2514'
const bottomRightCorner = '\u2518'
const horizontalBar = '\u2500'
const verticalBar = '\u2502'

func (w ContainerWidget) render(ctx RenderContext) RenderJob {
	const paddingX = 1
	const paddingY = 1
	// TODO some code refactoring before non-1 renders correctly
	const borderThickness = 1

	var childJob = w.child.render(RenderContext{
		constraints: Constraints{
			startingX: ctx.constraints.startingX + borderThickness + paddingX,
			finalX:    ctx.constraints.finalX,
			startingY: ctx.constraints.startingY + borderThickness + paddingY,
			finalY:    ctx.constraints.finalY,
		},
		fg: ctx.fg,
		bg: ctx.bg,
	})

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

func (w TextWidget) render(ctx RenderContext) RenderJob {
	var width int = 0
	// TODO layout
	for _, c := range w.msg {
		width += runewidth.RuneWidth(c)
	}

	return RenderJob{
		width:  width,
		height: 1, // TODO
		exec: func(x, y int) {
			for _, c := range w.msg {
				termbox.SetCell(x, y, c, ctx.fg, ctx.bg)
				x += runewidth.RuneWidth(c)
			}
		},
	}
}
