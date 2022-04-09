package main

func initialModel(l []Copyed) clipboardHistory {
	return clipboardHistory{
		list:  l,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}