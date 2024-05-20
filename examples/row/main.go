package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	var texts = (func(msgs []string) []boxes.Widget {
		var widgets = make([]boxes.Widget, len(msgs))

		for i, text := range msgs {
			widgets[i] = boxes.Container{
				Child: boxes.Text{Msg: text},
			}
		}

		return widgets
	})([]string{
		"Hello, world!",
		"Hmm...",
		"Goodbye, world!",
	})

	boxes.Run(
		boxes.Row{
			Children: texts,
		},
	)
}
