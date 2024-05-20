package boxes

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"os"
)

func Run(w Widget) {
	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

	var root = w
	root.render(ctx, cons).exec(0, 0)
	termbox.Flush()

	_ = termbox.PollEvent()

	termbox.Close()
}
