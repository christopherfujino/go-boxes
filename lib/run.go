package boxes

import (
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
		var renderBox = w.render(ctx, cons).exec(0, 0)
		termbox.Flush()

		var event = termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			DebugPrint("Got Key")
			break
		case termbox.EventMouse:
			var maybe = findBox(renderBox, event.MouseY, event.MouseY)
			if maybe != nil && maybe.OnClick != nil {
				maybe.OnClick()
			} else {
				if maybe == nil {
					DebugPrint("Got a nil back from findBox()")
				} else {
					DebugPrint("Got a RenderBox, but it's OnClick was nil")
				}
			}
		default:
			DebugPrint("Unknown event")
			panic(event.Type)
		}
	}
}
