// Package main
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  1.plus_test
 * @Version: 0.1.0
 * @Date: 2021/12/9 1:24 下午
 */
package main

import "testing"

func TestAdd(t *testing.T) {
	inputs := [][2]int{
		{1, 2},
		{3, 0},
		{123, 321},
		{99, 1},
		{-17, -3},
	}
	outputs := []int{
		3, 3, 444, 100, -20,
	}
	for k, input := range inputs {
		sum := add(input[0], input[1])
		if outputs[k] != sum {
			t.Error(input)
		}
	}
}

func TestAddNoRecurrence(t *testing.T) {
	inputs := [][2]int{
		{1, 2},
		{3, 0},
		{123, 321},
		{99, 1},
		{-17, -3},
	}
	outputs := []int{
		3, 3, 444, 100, -20,
	}
	for k, input := range inputs {
		sum := addNoRecurrence(input[0], input[1])
		if outputs[k] != sum {
			t.Error(input)
		}
	}
}
