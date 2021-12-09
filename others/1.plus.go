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
*/

// add1 实现 一位数 只能实现0和1之间的加法
func add1(a, b int) int {
	carry := (a & b) << 1  // 进位的值
	unCarry := a ^ b       // 不考虑进位，相加的值
	return carry ^ unCarry // 或者 carry | unCarry
}

// add
// @Description:
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
	return add(carry, unCarry)
}
