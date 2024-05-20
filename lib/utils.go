package boxes

// Map from one type to a widget.
func Map[I any, O any](elements []I, f func(I) O) []O {
	var mappedElements = make([]O, len(elements))

	for i, element := range elements {
		mappedElements[i] = f(element)
	}

	return mappedElements
}
