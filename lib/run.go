package boxes

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

func Run(w Widget) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Enable mouse input
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	var windowWidth, windowHeight = termbox.Size()
	var ctx = Context{
		fg: termbox.ColorBlue,
		bg: termbox.ColorDefault,
	}
	var cons = Constraints{
		minWidth:  0,
		maxWidth:  windowWidth,
		minHeight: 0,
		maxHeight: windowHeight,
	}

	for {
		w.render(ctx, cons).exec(0, 0)
		termbox.Flush()

		var event = termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			break
		case termbox.EventMouse:
			panic(fmt.Sprintf("(%d, %d)", event.MouseX, event.MouseY))
		default:
			panic(event.Type)
		}
	}
}
