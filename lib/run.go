package boxes

import (
	termbox "github.com/nsf/termbox-go"
)

func Run(w Widget) {
	err := termbox.Init()

	if err != nil {
		panic(err)
	}

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

	w.render(ctx, cons).exec(0, 0)
	termbox.Flush()

	_ = termbox.PollEvent()

	termbox.Close()
}
