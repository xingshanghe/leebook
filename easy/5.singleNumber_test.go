package main

import "testing"

func TestSingleNumber(t *testing.T) {
	inputs := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
	}
	out := []int{1, 4}
	for k, input := range inputs {
		n := singleNumber(input)
		if out[k] != n {
			t.Error(input)
		}
	}
}
