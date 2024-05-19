package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"os"
)

func main() {
	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var windowWidth, windowHeight = termbox.Size()
	var ctx = RenderContext{
		constraints: Constraints{
			startingX: 0,
			finalX:    windowWidth - 1,
			startingY: 0,
			finalY:    windowHeight - 1,
		},
		fg: termbox.ColorBlue,
		bg: termbox.ColorDefault,
	}

	var root = CenterWidget{
		child: ContainerWidget{
			child: TextWidget{"Hello from a Widget"},
		},
	}
	root.render(ctx).exec(0, 0)
	termbox.Flush()

	var event = termbox.PollEvent()

	termbox.Close()
	switch event.Type {
	case termbox.EventKey:
		fmt.Println("EventKey")
	case termbox.EventResize:
		fmt.Println("EventResize")
	case termbox.EventMouse:
		fmt.Println("EventMouse")
	case termbox.EventError:
		fmt.Println("EventError")
	case termbox.EventInterrupt:
		fmt.Println("EventInterrupt")
	case termbox.EventRaw:
		fmt.Println("EventRaw")
	case termbox.EventNone:
		fmt.Println("EventNone")
	}
}
