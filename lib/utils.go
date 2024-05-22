package boxes

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// Map from one type to a widget.
func Map[I any, O any](elements []I, f func(I) O) []O {
	var mappedElements = make([]O, len(elements))

	for i, element := range elements {
		mappedElements[i] = f(element)
	}

	return mappedElements
}

func findBox(box RenderBox, x, y int) *RenderBox {
	if x < box.Left || x > box.Right || y > box.Bottom || y < box.Top {
		//panic(fmt.Sprintf(
		//	"[L: %d, R: %d, T: %d, B: %d] (%d, %d)",
		//	box.Left,
		//	box.Right,
		//	box.Top,
		//	box.Bottom,
		//	x,
		//	y,
		//))
		return nil
	}

	var children = box.Children
	if children == nil {
		if box.OnClick == nil {
			return nil
		}
		return &box
	}

	for _, child := range children {
		var maybe = findBox(child, x, y)
		if maybe != nil && maybe.OnClick != nil {
			// TODO could multiple children contain the same point?
			return maybe
		}
	}
	if box.OnClick == nil {
		return nil
	}
	return &box
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
