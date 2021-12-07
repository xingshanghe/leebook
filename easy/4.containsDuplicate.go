package main

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	var result bool
	var lastValue int
	for k, v := range nums {
		if k == 0 {
			lastValue = v
		} else {
			if v != lastValue {
				lastValue = v
			} else {
				result = true
			}
		}
	}
	//bucket := make(map[int]int)
	//for _,v:= range nums{
	//
	//	if _,ok:= bucket[v];ok{
	//		result = true
	//		break
	//	}else{
	//		bucket[v] = 1
	//	}
	//}
	return result
}
