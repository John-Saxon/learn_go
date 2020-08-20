package main

import (
	"fmt"
)

// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。考查边界。示例：
//      输入: [2,3,1,1,4]
// 输出: true
// 解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1   跳 3 步到达最后一个位置。

// 输入: [3,2,1,0,4]
// 输出: false
// 解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。

var debug bool

// func init() {
// 	debug = false
// }

func main() {
	debug = false
	testJump()
}

func testJump() {
	s1 := []int{2, 3, 1, 1, 4}
	s2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(s1))
	fmt.Println("--------------------------")
	fmt.Println(canJump(s2))
	fmt.Println("--------------------------")
	s3 := []int{0, 1}
	fmt.Println(canJump(s3))
	fmt.Println("--------------------------")
	s4 := []int{5, 0, 0, 0, 0}
	fmt.Println(canJump(s4))
}

func canJump(s []int) bool {
	if len(s) <= 1 {
		return true
	}
	lastIndex := len(s) - 1
	if debug {
		fmt.Println("last = ", lastIndex)
	}
	can := true
	for i := lastIndex - 1; i >= 0; i-- {
		if debug {
			fmt.Println("i = ", i)
			fmt.Println("s[i] = ", s[i])
			fmt.Println("last = ", lastIndex)
		}
		if s[i]+i >= lastIndex {
			lastIndex = i
			can = true
		} else {
			can = false
		}
		if debug {
			fmt.Println("can = ", can)
		}
	}
	return can
}
