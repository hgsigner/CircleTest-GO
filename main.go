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
	processed := process_list(slice, pos)
	place := find_safer_place(processed, 0)

	output := fmt.Sprintf("The safest place of the list starting from (%d) is: %d\n", pos, place)

	fmt.Fprintf(out, output)
}

func main() {

	flag.Var(&mylist, "i", "List of items")
	position := flag.Int("position", 0, "Set the position")
	flag.Parse()

	run(mylist, *position, os.Stdout)

}
