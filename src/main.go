package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"
)

var ValidRegex = regexp.MustCompile("^(0|-?[1-9][0-9]*|[A-Za-z][0-9A-Z_a-z]*)$")
var NumberRegex = regexp.MustCompile("^-?[0-9]+$")

func isValidString(str string) bool {
	return ValidRegex.MatchString(str)
}

func isNumberString(str string) bool {
	return NumberRegex.MatchString(str)
}

func insertBefore(val *string, item *Item[string]) bool {
	if isNumberString(*val) && isNumberString(item.value) {
		valInt, _ := new(big.Int).SetString(*val, 10)
		othInt, _ := new(big.Int).SetString(item.value, 10)
		return valInt.Cmp(othInt) < 1
	} else {
		return *val < item.value
	}
}

func valueEquals(item *Item[string], val *string) bool {
	return item.value == *val
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var start *Item[string] = nil

	begin := true
	var input string

	for true {
		if !begin {
			fmt.Println()
		} else {
			begin = false
		}

		fmt.Println("Awaiting input...")
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")

		if len(input) == 0 {
			fmt.Println("\nProgram terminated!")
			RemoveAll(&start)
			return
		} else if input[0] == '~' {
			if len(input) == 1 {
				fmt.Println("\nDeleting list...")
				RemoveAll(&start)
			} else {
				input = input[1:]
				if isValidString(input) {
					fmt.Println("\nRemoving item...")
					RemoveItem(&start, input, valueEquals)
				} else {
					fmt.Println("\nCould not parse input!")
				}
			}
		} else if input == "l" {
			fmt.Println("\nLoop print...")
			PrintLoop(start)
		} else if input == "i" {
			fmt.Println("\nIterator print...")
			PrintIterator(start)
		} else if input == "a" {
			fmt.Println("\nArray print not implemented!")
		} else if input == "r" {
			fmt.Println("\nRecursive print...")
			PrintRecursive(start)
		} else if input == "f" {
			fmt.Println("\nFold print...")
			PrintFold(start)
		} else if input == "b" {
			fmt.Println("\nFoldback print...")
			PrintFoldback(start)
		} else if isValidString(input) {
			fmt.Println("\nInserting item...")
			InsertItem(&start, input, insertBefore)
		} else {
			fmt.Println("\nCould not parse input!")
		}
	}
}
