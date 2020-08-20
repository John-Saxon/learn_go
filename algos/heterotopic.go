package main

import (
	"fmt"
)

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 字母异位词指字母相同，但排列不同的字符串。
// 同时考查候选人对题目的理解以及沟通能力
// 考查边界

func main() {
	s1 := "aacdf"
	s2 := "cdafa"
	s3 := "cdffa"
	fmt.Println(hetero(s1, s2))
	fmt.Println(hetero(s1, s3))
	fmt.Println(hetero(s3, s2))
	fmt.Println(h2(s1, s2))
	fmt.Println(h2(s1, s3))
	fmt.Println(h2(s3, s2))
	fmt.Println(h3(s1, s2))
	fmt.Println(h3(s1, s3))
	fmt.Println(h3(s3, s2))
}

func hetero(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, c := range s1 {
		if v, ok := m1[c]; ok {
			m1[c] = v + 1
		} else {
			m1[c] = 1
		}
	}
	for _, c := range s2 {
		if v, ok := m2[c]; ok {
			m2[c] = v + 1
		} else {
			m2[c] = 1
		}
	}
	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			if v != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func h2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var s0 byte = 0
	for i := 0; i < len(s); i++ {
		s0 = s0 ^ s[i]
	}
	var t0 byte = 0
	for i := 0; i < len(t); i++ {
		t0 = t0 ^ t[i]
	}
	return s0 == t0
}

func h3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var s0 byte = 0
	for i := 0; i < len(s); i++ {
		s0 = s0 ^ s[i]
	}
	for i := 0; i < len(t); i++ {
		s0 = s0 ^ t[i]
	}
	return s0 == 0
}
