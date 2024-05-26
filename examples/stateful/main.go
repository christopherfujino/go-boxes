package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	boxes.Run(
		boxes.Row{
			Children: []boxes.Widget{
				boxes.Container{
					Child: boxes.Text{Msg: "Hello, world!"},
				},
				&ui{
					child: boxes.Container{
						Child: boxes.Text{Msg: "I'm a button"},
					},
				},
				boxes.Container{
					Child: boxes.Text{Msg: "Goodbye, world!"},
				},
			},
		},
	)
}

type ui struct {
	state int
	child boxes.Widget
}

func (w *ui) Render(ctx boxes.Context, cons boxes.Constraints) boxes.RenderJob {
	var childJob = w.child.Render(ctx, cons)

	return boxes.RenderJob{
		Width:  childJob.Width,
		Height: childJob.Height,
		Exec: func(x, y int) boxes.RenderBox {
			return boxes.RenderBox{
				Left:   x,
				Top:    y,
				Right:  x + childJob.Width - 1,
				Bottom: y + childJob.Height - 1,
				OnClick: func() {
					w.state += 1
				},
				Children: []boxes.RenderBox{childJob.Exec(x, y)},
			}
		},
	}
}
