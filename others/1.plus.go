// Package others
/**
 * @Author: xingshanghe
 * @Description:
 * @File:  1.plus
 * @Version: 0.1.0
 * @Date: 2021/12/9 10:46 上午
 */
package main

/**
不用 "+"号实现加法
所涉及基础知识储备：

	位运算：
	&(与): 	两位都为1时，结果为1
	|(或): 	两位都为0时，结果为0
	^(异或):	两位都为0时，结果为1；另 a^a=0,a^0=a

	计算机中 0和1相加，几种情况（二进制情况）:
	0 + 0 = 00
	0 + 1 = 01
	1 + 0 = 01
	1 + 1 = 10

	不考虑进位相加: 	a^b
	进位值: 			(a&b)<<1
*/

// add1 实现 一位数 只能实现0和1之间的加法
func add1(a, b int) int {
	carry := (a & b) << 1  // 进位的值
	unCarry := a ^ b       // 不考虑进位，相加的值
	return carry ^ unCarry // 或者 carry | unCarry
}

// add
// @Description: 加法，递归实现
// @Date: 2021-12-09 13:15:43
// @param a
// @param b
// @return int
//
func add(a, b int) int {
	if a == 0 || b == 0 {
		return a ^ b
	}
	carry := (a & b) << 1 // 进位的值
	unCarry := a ^ b      // 不考虑进位，相加的值
	return add(unCarry, carry)
}

func addNoRecurrence(a, b int) int {
	for {
		if b == 0 {
			break
		}
		unCarry := a ^ b
		carry := (a & b) << 1

		a = unCarry
		b = carry
	}
	return a
}
