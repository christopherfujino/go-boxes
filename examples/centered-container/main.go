package main

import (
	"github.com/christopherfujino/go-boxes"
)

func main() {
	boxes.Run(
		boxes.CenterWidget{
			Child: boxes.ContainerWidget{
				Child: boxes.TextWidget{
					Msg: "Hello from a Widget",
				},
			},
		},
	)
}
