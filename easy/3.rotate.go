// Package main
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  3.rotate
 * @Version: 0.1.0
 * @Date: 2021/12/7 11:55 下午
 */
package main

// rotate
// @Description:  投机取巧1
// @Date: 2021-12-08 00:57:00
// @param nums
// @param k
//
func rotate(nums []int, k int) {
	l := len(nums)
	n := k % l
	if l > n {
		// nums = append([]int{}, append(nums[l-n:], nums[:l-n]...)...) // 地址变化，
		// nums = append(nums[l-n:], nums[:l-n]...)
		nums = append(nums[0:0], append(nums[l-n:], nums[:l-n]...)...) // 浅拷贝
	}
}
