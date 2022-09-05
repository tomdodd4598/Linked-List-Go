package main

import (
	"fmt"
)

type Item struct {
	value string
	next  *Item
}

func (item *Item) PrintGetNext() *Item {
	fmt.Print(item.value)
	suffix := ", "
	if item.next == nil {
		suffix = "\n"
	}
	fmt.Print(suffix)
	return item.next
}

type ItemIterator struct {
	item *Item
}

func (iter *ItemIterator) HasNext() bool {
	return iter.item != nil
}

func (iter *ItemIterator) Next() *Item {
	next := iter.item
	iter.item = iter.item.next
	return next
}

func ItemFold(fSome func(*Item, *Item, string) string, fLast func(*Item, string) string, fEmpty func(string) string, accumulator string, item *Item) string {
	if item != nil {
		next := item.next
		if next != nil {
			return ItemFold(fSome, fLast, fEmpty, fSome(item, next, accumulator), next)
		} else {
			return fLast(item, accumulator)
		}
	} else {
		return fEmpty(accumulator)
	}
}

func ItemFoldback(fSome func(*Item, *Item, string) string, fLast func(*Item) string, fEmpty func() string, generator func(string) string, item *Item) string {
	if item != nil {
		next := item.next
		if next != nil {
			newGenerator := func(innerVal string) string {
				return generator(fSome(item, next, innerVal))
			}
			return ItemFoldback(fSome, fLast, fEmpty, newGenerator, next)
		} else {
			return generator(fLast(item))
		}
	} else {
		return generator(fEmpty())
	}
}
