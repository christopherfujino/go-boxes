package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	boxes.Run(
		boxes.Center{
			Child: boxes.Container{
				Child: boxes.Text{
					Msg: "Hello from a Widget",
				},
			},
		},
	)
}
