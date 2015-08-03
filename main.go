package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type intslice []int

var mylist intslice

var Usage = func() {
	fmt.Fprintf(os.Stderr, `
NAME:
  circle_array - Find the safest place on a list of integers
USAGE:
  circle_array [--options]
VERSION:
  %s
COMMANDS:
  help, h Shows a list of commands or help
EXAMPLES:
  circle_array -i "1" -i "2" -i "3" -i "4" -i "5" -position "0"
OPTIONS:
	`, Version)

	flag.PrintDefaults()
}

var Version = "0.0.1"

func (i *intslice) String() string {
	return fmt.Sprintf("%d", *i)
}

func (i *intslice) Set(value string) error {

	val, err := strconv.Atoi(value)
	if err != nil {
		*i = append(*i, -1)
	} else {
		*i = append(*i, val)
	}

	return nil

}

func process_list(slice []int, start_at int) []int {

	if start_at == 0 {
		return slice
	} else {
		return append(slice[start_at:len(slice)], slice[0:start_at]...)
	}

}

func find_safer_place(slice []int, turn_start_cut_at int) int {

	if len(slice) == 1 {
		return slice[0]
	}

	current_last_item := slice[len(slice)-1]
	new_slice := make([]int, 0)

	for i, item := range slice {
		if turn_start_cut_at == 0 {
			if i%2 == 0 {
				new_slice = append(new_slice, item)
			}
		} else {
			if i%2 != 0 {
				new_slice = append(new_slice, item)
			}
		}
	}

	new_slice_last_item := new_slice[len(new_slice)-1]

	if current_last_item == new_slice_last_item {
		return find_safer_place(new_slice, 1)
	} else {
		return find_safer_place(new_slice, 0)
	}

}

func run(slice []int, pos int, out io.Writer) {

	var output string

	if pos > len(slice)-1 {
		output = fmt.Sprintf("Your list doesn't cointain %d element(s)\n", pos+1)
	} else {
		processed := process_list(slice, pos)
		place := find_safer_place(processed, 0)

		output = fmt.Sprintf("The safest place of the list starting from (%d) is: %d\n", pos, place)
	}

	fmt.Fprintf(out, output)
}

func main() {

	flag.Var(&mylist, "i", "List of items")
	position := flag.Int("position", 0, "Set the position starting by 0")
	flag.Usage = Usage
	flag.Parse()

	run(mylist, *position, os.Stdout)

}
