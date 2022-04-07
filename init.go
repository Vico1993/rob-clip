package main

import "time"

func initialModel() clipboardHistory {
	copy := Copyed{
		word: GetValue(),
		date: time.Now(),
	}

	return clipboardHistory{
		list:  []Copyed{copy},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}