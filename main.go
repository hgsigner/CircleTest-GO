package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	processed := process_list(slice, 0)
	place := find_safer_place(processed, 0)

	fmt.Println(place)

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
