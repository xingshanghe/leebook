package main

import (
	"testing"
)

/**
示例 1:

输入: [1,2,3,1]
输出: true

示例 2:

输入: [1,2,3,4]
输出: false

示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/
func TestContainsDuplicate(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	outputs := []bool{true, false, true}
	for i, input := range inputs {
		n := containsDuplicate(input)
		if outputs[i] != n {
			t.Error(input)
		}
	}
}
