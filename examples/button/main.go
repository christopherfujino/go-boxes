package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	boxes.Run(
		boxes.Row{
			Children: boxes.Map(
				[]string{
					"Hello, world!",
					"Hmm...",
					"Goodbye, world!",
				},
				func(s string) boxes.Widget {
					return boxes.Container{Child: boxes.Text{Msg: s}}
				},
			),
		},
	)
}
