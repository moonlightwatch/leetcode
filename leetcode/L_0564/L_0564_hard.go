package L_0564

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/find-the-closest-palindrome/

// 564. 寻找最近的回文数

// 给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。

// “最近的”定义为两个整数差的绝对值最小。

// 1 <= n.length <= 18
// n 只由数字组成
// n 不含前导 0
// n 代表在 [1, 1018 - 1] 范围内的整数

// isPalindromic 判断是否回文
func isPalindromic(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

// makePalindromic 将数字翻转后拼到尾部制造回文
func makePalindromic(num int, p int) string {
	s := fmt.Sprintf("%d", num)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	runes = runes[p:]
	return fmt.Sprintf("%s%s", s, string(runes))
}

func nearestPalindromic(n string) string {
	l := len(n)
	if l < 10 {
		// 长度小于10的
		// 使用暴力求解：
		// 向着更大数字和更小数字，逐个数字判断是否回文，返回遇到的第一个回文数
		num, _ := strconv.Atoi(n)
		for i := 1; ; i++ {
			if isPalindromic(fmt.Sprintf("%d", num-i)) {
				return fmt.Sprintf("%d", num-i)
			}

			if isPalindromic(fmt.Sprintf("%d", num+i)) {
				return fmt.Sprintf("%d", num+i)
			}
		}
	}

	// 长度大于10

	// 取数字的前半部分，如奇数长度，则多取一位，记为数字 halfNum
	// 对上述获得的数字 halfNum，取 halfNum-1 和 halfNum+1
	// 对三个数字翻转后，拼在其后，获得新的三个回文数 p、p1、p2
	// 计算三个新回文数与原数字之差，记为 _p、 _p1、 _p2
	// 通过判断，取得目标数字，判断如下：
	//         1. 不与原数字相同的
	//         2. 与原数字之差最小的
	//         3. 如果与原数字之差最小的不止一个，则选择较小的

	num, _ := strconv.Atoi(n)
	part := n[:l/2+l%2]
	halfNum, _ := strconv.Atoi(part)

	p1, _ := strconv.Atoi(makePalindromic(halfNum+1, l%2))
	p, _ := strconv.Atoi(makePalindromic(halfNum, l%2))
	p2, _ := strconv.Atoi(makePalindromic(halfNum-1, l%2))

	_p1 := math.Abs(float64(p1 - num))
	_p := math.Abs(float64(p - num))
	_p2 := math.Abs(float64(p2 - num))

	if p == num {
		if _p1 < _p2 {
			return fmt.Sprintf("%d", p1)
		} else if _p1 > _p2 {
			return fmt.Sprintf("%d", p2)
		}
		if p1 < p2 {
			return fmt.Sprintf("%d", p1)
		} else {
			return fmt.Sprintf("%d", p2)
		}
	} else if (_p < _p1 && _p < _p2) || (_p < _p1 && _p == _p2 && p < p2) || (_p < _p2 && _p == _p1 && p < p1) {
		return fmt.Sprintf("%d", p)
	} else if (_p1 < _p && _p1 < _p2) || (_p1 < _p && _p1 == _p2 && p1 < p2) || (_p1 < _p2 && _p1 == _p && p1 < p) {
		return fmt.Sprintf("%d", p1)
	} else if (_p2 < _p1 && _p2 < _p) || (_p2 < _p1 && _p2 == _p && p2 < p) || (_p2 < _p && _p2 == _p1 && p2 < p1) {
		return fmt.Sprintf("%d", p2)
	}
	return ""
}

func Test564() {
	fmt.Println(nearestPalindromic("1805170081")) // 1805115081
	fmt.Println(nearestPalindromic("2147483647")) // 2147447412
	fmt.Println(nearestPalindromic("393"))        // 383
	fmt.Println(nearestPalindromic("909"))        // 919
	fmt.Println(nearestPalindromic("303"))        // 313
	fmt.Println(nearestPalindromic("3333"))       // 3223
	fmt.Println(nearestPalindromic("333"))        // 323
	fmt.Println(nearestPalindromic("101"))        // 99
	fmt.Println(nearestPalindromic("100"))        // 99
	fmt.Println(nearestPalindromic("321"))        // 323
	fmt.Println(nearestPalindromic("11"))         // 9
	fmt.Println(nearestPalindromic("123"))        // 121
	fmt.Println(nearestPalindromic("1"))          // 0

	fmt.Println(nearestPalindromic("99999999"))
	fmt.Println(nearestPalindromic("999999999"))
	fmt.Println(nearestPalindromic("8070450983"))
	fmt.Println(nearestPalindromic("80704505783"))
	fmt.Println(nearestPalindromic("807045057983"))
	fmt.Println(nearestPalindromic("8070450579283"))
	fmt.Println(nearestPalindromic("80704505792883"))
	fmt.Println(nearestPalindromic("807045053224883"))
	fmt.Println(nearestPalindromic("8070450532247983"))
	fmt.Println(nearestPalindromic("80704505322479928"))
	fmt.Println(nearestPalindromic("807045053224792883"))
}
