package boxes

import (
	"fmt"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

const debug = true

func Run(w Widget) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

	var ctx = Context{
		fg: termbox.ColorBlue,
		bg: termbox.ColorDefault,
	}

	defer (func() {
		if debugger == nil {
			return
		}

		debugger := debugger.(*strings.Builder)
		fmt.Println(debugger.String())
	})()

	defer termbox.Close()

	// Enable mouse input
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	var windowWidth, windowHeight = termbox.Size()

	var cons = Constraints{
		minWidth:  0,
		maxWidth:  windowWidth,
		minHeight: 0,
		maxHeight: windowHeight,
	}

	for {
		var renderBox = w.Render(ctx, cons).Exec(0, 0)
		termbox.Flush()

		var event = termbox.PollEvent()
		switch event.Type {
		case termbox.EventKey:
			DebugPrint("Got Key")
			break
		case termbox.EventMouse:
			var maybe = renderBox.findHitBox(event.MouseX, event.MouseY)
			if maybe != nil && maybe.OnClick != nil {
				maybe.OnClick()
				DebugPrint("Clicking!")
			} else {
				//panic(fmt.Sprintf("Click at (%d, %d) missed hitbox in\n%v"))
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
