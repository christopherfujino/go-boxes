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

	var root = w
	root.render(ctx).exec(0, 0)
	termbox.Flush()

  _ = termbox.PollEvent()

	termbox.Close()
}
