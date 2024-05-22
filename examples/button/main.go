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
				boxes.Clickable{
					Child: boxes.Container{
						Child: boxes.Text{Msg: "I'm a button"},
					},
					OnClick: func() {
						panic("yay")
					},
				},
				boxes.Container{
					Child: boxes.Text{Msg: "Goodbye, world!"},
				},
			},
		},
	)
}
