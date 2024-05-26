package boxes

import (
	"fmt"
	"io"
	//"strings"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// Map from one type to a widget.
func Map[I any, O any](elements []I, f func(I) O) []O {
	var mappedElements = make([]O, len(elements))

	for i, element := range elements {
		DebugLog(fmt.Sprintf("in Map, %d - %v", i, element))
		mappedElements[i] = f(element)
	}

	return mappedElements
}

func DebugPrint(msg string) {
	var width, height = termbox.Size()
	var length = runewidth.StringWidth(msg)
	var x = width - length - 1
	for _, c := range msg {
		termbox.SetCell(x, height-2, c, termbox.ColorBlue, termbox.ColorLightRed)
		x += runewidth.RuneWidth(c)
	}
}

//var debugger io.Writer = &strings.Builder{}
var debugger io.Writer = nil

func DebugLog(msg string) {
	if debugger == nil {
		return
	}

	_, err := debugger.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	_, err = debugger.Write([]byte{10})
	if err != nil {
		panic(err)
	}
}
