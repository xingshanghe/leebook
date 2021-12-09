// Package main
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  6.intersect_test
 * @Version: 0.1.0
 * @Date: 2021/12/9 9:46 上午
 */
package main

import "testing"

/**
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
func TestIntersect(t *testing.T) {
	inputs := [][][]int{
		{{1, 2, 2, 1}, {2, 2}},
		{{4, 9, 5}, {9, 4, 9, 8, 4}},
		{{1, 2, 3, 3, 4, 5, 6, 7, 8}, {3, 3, 3}},
	}
	out := [][]int{
		{2, 2},
		{4, 9},
		{3, 3},
	}
	for k, input := range inputs {
		n := intersect(input[0], input[1])
		if !equal(out[k], n) {
			t.Error(input)
		}
	}
}
