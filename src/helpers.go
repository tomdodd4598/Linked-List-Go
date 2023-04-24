package main

import (
	"fmt"
)

func InsertItem[T any](start **Item[T], val T, insertBefore func(*T, *Item[T]) bool) {
	fmt.Printf("Creating item: %v\n", val)

	for *start != nil && !insertBefore(&val, *start) {
		start = &(*start).next
	}

	*start = &Item[T]{val, *start}
}

func RemoveItem[T any](start **Item[T], val T, valueEquals func(*Item[T], *T) bool) {
	for *start != nil && !valueEquals(*start, &val) {
		start = &(*start).next
	}

	if *start == nil {
		fmt.Printf("Item %v does not exist!\n", val)
	} else {
		*start = (*start).next
		fmt.Printf("Removed item: %v\n", val)
	}
}

func RemoveAll[T any](start **Item[T]) {
	*start = nil
}

func PrintLoop[T any](start *Item[T]) {
	for start != nil {
		start = start.PrintGetNext()
	}
}

func PrintIterator[T any](start *Item[T]) {
	if start != nil {
		for iter := (ItemIterator[T]{start}); iter.HasNext(); {
			iter.Next().PrintGetNext()
		}
	}
}

func PrintRecursive[T any](start *Item[T]) {
	if start != nil {
		PrintRecursive(start.PrintGetNext())
	}
}

func PrintFold[T any](start *Item[T]) {
	fSome := func(current *Item[T], _ *Item[T], accumulator string) string {
		return fmt.Sprintf("%s%v, ", accumulator, current.value)
	}

	fLast := func(current *Item[T], accumulator string) string {
		return fmt.Sprintf("%s%v\n", accumulator, current.value)
	}

	fEmpty := func(accumulator string) string {
		return accumulator
	}

	fmt.Print(ItemFold(fSome, fLast, fEmpty, "", start))
}

func PrintFoldback[T any](start *Item[T]) {
	fSome := func(current *Item[T], _ *Item[T], innerVal string) string {
		return fmt.Sprintf("%v, %s", current.value, innerVal)
	}

	fLast := func(current *Item[T]) string {
		return fmt.Sprintf("%v\n", current.value)
	}

	fEmpty := func() string {
		return ""
	}

	id := func(x string) string {
		return x
	}

	fmt.Print(ItemFoldback(fSome, fLast, fEmpty, id, start))
}
