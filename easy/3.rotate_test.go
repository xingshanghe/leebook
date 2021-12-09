// Package main
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  3.rotate
 * @Version: 0.1.0
 * @Date: 2021/12/7 11:55 下午
 */
package main

import "testing"

// rotate
// @Description:  投机取巧1
// @Date: 2021-12-08 00:57:00
// @param nums
// @param k
//
func TestRotate(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
	}
	for _, input := range inputs {
		rotate(input, 3)
	}
}
