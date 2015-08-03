package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests_process_list = []struct {
	in  []int
	pos int
	out []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		0,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		3,
		[]int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3},
	},
}

var tests_safer_place_list = []struct {
	in  []int
	pos int
	out int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		0,
		5,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		4,
		9,
	},
}

func Test_Process_List_Func(t *testing.T) {
	a := assert.New(t)

	for _, item := range tests_process_list {
		a.Equal(item.out, process_list(item.in, item.pos))
	}
}

func Test_Find_Safer_Place(t *testing.T) {
	a := assert.New(t)

	for _, item := range tests_safer_place_list {
		proc := process_list(item.in, item.pos)
		a.Equal(item.out, find_safer_place(proc, 0))
	}
}
