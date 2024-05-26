package boxes

// An alignment widget.
type Center struct {
	Child Widget
}

func (w Center) Render(ctx Context, cons Constraints) RenderJob {
	var width = cons.maxWidth
	var childJob = w.Child.Render(ctx, cons)
	var leftPad = (width - childJob.Width) / 2
	return RenderJob{
		Width:  width,
		Height: childJob.Height,
		Exec: func(x, y int) RenderBox {
			var childBox = childJob.Exec(x+leftPad, y)
			return RenderBox{
				Left:   x,
				Right:  x + childJob.Width - 1,
				Top:    y,
				Bottom: y + childJob.Height - 1,
				Children: []RenderBox{childBox},
			}
		},
	}
}

// An alignment widget.
type Row struct {
	Children []Widget
}

func (w Row) Render(ctx Context, cons Constraints) RenderJob {
	DebugLog("Laying out a row")
	var cumulativeWidth = 0
	var maxHeight = 0

	var jobs = Map(
		w.Children,
		func(child Widget) RenderJob {
			var job = child.Render(
				ctx,
				Constraints{
					maxWidth: cons.maxWidth - cumulativeWidth,
				},
			)
			cumulativeWidth += job.Width
			if job.Height > maxHeight {
				maxHeight = job.Height
			}
			return job
		},
	)

	if cumulativeWidth > cons.maxWidth {
		panic("whoops!")
	}

	if maxHeight > cons.maxHeight {
		panic("whoops!")
	}

	return RenderJob{
		Width:  cumulativeWidth,
		Height: maxHeight,
		Exec: func(x, y int) RenderBox {
			DebugLog("Rendering a row")
			var left = x
			var childBoxes = Map(jobs, func (job RenderJob) RenderBox {
				var box = job.Exec(left, y)
				left += job.Width
				return box
			})

			return RenderBox{
				Left:   x,
				Right:  x + cumulativeWidth - 1,
				Top:    y,
				Bottom: y + maxHeight - 1,
				Children: childBoxes,
			}
		},
	}
}
