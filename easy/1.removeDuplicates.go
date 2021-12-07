package main

/*
要求: 原地修改

示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

提示：

0 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按升序排列

	作者：力扣 (LeetCode)
	链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// removeDuplicates 双指针去除重复
func removeDuplicates(in []int) (length int, out []int) {
	inLength := len(in)

	left := 0
	right := 1

	/*
		for ; right < inLength; right++ {
			if in[left] != in[right] {
				left++
				in[left] = in[right]

			}
		}
	*/
	// 循环数组更容易理解
	for range in {
		if right < inLength { // 注意超限panic
			if in[left] != in[right] {
				left++
				in[left] = in[right]
			}
			right++
		}
	}

	return left + 1, in[:left+1]
}
