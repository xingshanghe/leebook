package main

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{7, 6, 4, 3, 1},
	}
	outputs := []int{7, 4, 0}
	for i, input := range inputs {
		n := maxProfit(input)
		if outputs[i] != n {
			t.Error(input)
		}
	}

}
