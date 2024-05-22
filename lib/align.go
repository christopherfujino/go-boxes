package boxes

// An alignment widget.
type Center struct {
	Child Widget
}

func (w Center) render(ctx Context, cons Constraints) RenderJob {
	var width = cons.maxWidth
	var childJob = w.Child.render(ctx, cons)
	var leftPad = (width - childJob.width) / 2
	return RenderJob{
		width:  width,
		height: childJob.height,
		exec: func(x, y int) RenderBox {
			childJob.exec(x+leftPad, y)
			return RenderBox{
				Left:   x,
				Right:  x + childJob.width - 1,
				Top:    y,
				Bottom: y + childJob.height - 1,
			}
		},
	}
}

// An alignment widget.
type Row struct {
	Children []Widget
}

func (w Row) render(ctx Context, cons Constraints) RenderJob {
	var cumulativeWidth = 0
	var maxHeight = 0

	var jobs = Map(
		w.Children,
		func(child Widget) RenderJob {
			var job = child.render(
				ctx,
				Constraints{
					maxWidth: cons.maxWidth - cumulativeWidth,
				},
			)
			cumulativeWidth += job.width
			if job.height > maxHeight {
				maxHeight = job.height
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
		width:  cumulativeWidth,
		height: maxHeight,
		exec: func(x, y int) RenderBox {
			var left = x
			for _, job := range jobs {
				job.exec(left, y)
				left += job.width
			}

			return RenderBox{
				Left:   x,
				Right:  x + cumulativeWidth - 1,
				Top:    y,
				Bottom: y + maxHeight - 1,
			}
		},
	}
}
