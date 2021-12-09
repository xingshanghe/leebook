package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	inputs := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	}
	outputs := [][]int{
		{1, 2},
		{0, 1, 2, 3, 4},
	}
	for i, input := range inputs {
		n, out := removeDuplicates(input)
		if len(outputs[i]) != n || !equal(out, outputs[i]) {
			t.Error(input)
		}
	}
}

// equal reflect.DeepEqual效率不高
func equal(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
