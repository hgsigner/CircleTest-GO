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

func Test_Process_List_Func(t *testing.T) {
	a := assert.New(t)

	for _, item := range tests_process_list {
		a.Equal(item.out, process_list(item.in, item.pos))
	}
}
