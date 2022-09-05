package main

import (
	"fmt"
)

func InsertItem(start *Item, val string, insertBefore func(string, *Item) bool) *Item {
	fmt.Printf("Creating item: %s\n", val)
	current := start
	var previous *Item = nil

	for current != nil && !insertBefore(val, current) {
		previous = current
		current = current.next
	}
	item := &Item{val, current}

	if previous == nil {
		start = item
	} else {
		previous.next = item
	}

	return start
}

func RemoveItem(start *Item, val string, valueEquals func(*Item, string) bool) *Item {
	current := start
	var previous *Item = nil

	for current != nil && !valueEquals(current, val) {
		previous = current
		current = current.next
	}

	if current == nil {
		fmt.Printf("Item %s does not exist!\n", val)
	} else {
		if previous == nil {
			start = current.next
		} else {
			previous.next = current.next
		}
		fmt.Printf("Removed item: %s\n", val)
	}

	return start
}

func RemoveAll(start *Item) *Item {
	return nil
}

func PrintLoop(start *Item) {
	for start != nil {
		start = start.PrintGetNext()
	}
}

func PrintIterator(start *Item) {
	if start != nil {
		for iter := (ItemIterator{start}); iter.HasNext(); {
			iter.Next().PrintGetNext()
		}
	}
}

func PrintRecursive(start *Item) {
	if start != nil {
		PrintRecursive(start.PrintGetNext())
	}
}

func PrintFold(start *Item) {
	fSome := func(current *Item, _ *Item, accumulator string) string {
		return fmt.Sprintf("%s%s, ", accumulator, current.value)
	}

	fLast := func(current *Item, accumulator string) string {
		return fmt.Sprintf("%s%s\n", accumulator, current.value)
	}

	fEmpty := func(accumulator string) string {
		return accumulator
	}

	fmt.Print(ItemFold(fSome, fLast, fEmpty, "", start))
}

func PrintFoldback(start *Item) {
	fSome := func(current *Item, _ *Item, innerVal string) string {
		return fmt.Sprintf("%s, %s", current.value, innerVal)
	}

	fLast := func(current *Item) string {
		return fmt.Sprintf("%s\n", current.value)
	}

	fEmpty := func() string {
		return ""
	}

	id := func(x string) string {
		return x
	}

	fmt.Print(ItemFoldback(fSome, fLast, fEmpty, id, start))
}
