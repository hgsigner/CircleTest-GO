package main

func main() {

	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

}

func process_list(slice []int, start_at int) []int {

	if start_at == 0 {
		return slice
	} else {
		return append(slice[start_at:len(slice)], slice[0:start_at]...)
	}

}
