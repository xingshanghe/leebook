package main

func singleNumber(nums []int) int {
	var num int
	for _, v := range nums {
		num ^= v
	}
	return num
}
