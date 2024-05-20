package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	boxes.Run(
		boxes.Row{
			Children: []boxes.Widget{
				boxes.Text{Msg: "Hello, world!"},
				boxes.Clickable{
					Child: boxes.Text{Msg: "Hmm..."},
					OnClick: func() {
						panic("yay")
					},
				},
				boxes.Text{Msg: "Goodbye, world!"},
			},
		},
	)
}
