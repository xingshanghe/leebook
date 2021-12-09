// Package main
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  5.intersect
 * @Version: 0.1.0
 * @Date: 2021/12/9 9:19 上午
 */
package main

/*
示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func intersect(nums1, nums2 []int) []int {
	max := nums1
	min := nums2
	if len(nums2) > len(nums1) {
		max = nums2
		min = nums1
	}
	intersectMap := make(map[int]int) //计数map
	for _, v := range max {
		if _, ok := intersectMap[v]; ok {
			intersectMap[v] += 1
		} else {
			intersectMap[v] = 1
		}
	}
	result := make([]int, 0)
	// result := []int{}
	for _, v := range min {
		if vk, ok := intersectMap[v]; ok && vk > 0 {
			intersectMap[v] -= 1
			result = append(result, v)
		} else {
			intersectMap[v] = -1
		}
	}
	//var result []int
	// result := make([]int, 0)
	// result := []int{}

	return result
}
