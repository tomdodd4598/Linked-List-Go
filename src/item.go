package main

import (
	"fmt"
)

type Item[T any] struct {
	value T
	next  *Item[T]
}

func (item *Item[T]) PrintGetNext() *Item[T] {
	var suffix string
	if item.next == nil {
		suffix = "\n"
	} else {
		suffix = ", "
	}
	fmt.Printf("%v%s", item.value, suffix)
	return item.next
}

type ItemIterator[T any] struct {
	item *Item[T]
}

func (iter *ItemIterator[T]) HasNext() bool {
	return iter.item != nil
}

func (iter *ItemIterator[T]) Next() *Item[T] {
	next := iter.item
	iter.item = iter.item.next
	return next
}

func ItemFold[T any, A any, R any](fSome func(*Item[T], *Item[T], A) A, fLast func(*Item[T], A) R, fEmpty func(A) R, accumulator A, item *Item[T]) R {
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

func ItemFoldback[T any, A any, R any](fSome func(*Item[T], *Item[T], A) A, fLast func(*Item[T]) A, fEmpty func() A, generator func(A) R, item *Item[T]) R {
	if item != nil {
		next := item.next
		if next != nil {
			newGenerator := func(innerVal A) R {
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
