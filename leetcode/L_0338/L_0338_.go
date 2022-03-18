package L_0338

import "fmt"

// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

/*

思路：

遍历一下，对每个数字求其比特位1的数量。

求比特位1的数量，可以逐位和 1 进行与操作，此位是 1 则返回 1，0 则返回 0，累加结果即可。

由于 golang 中 int 的长度是 32 （详见 buildin.go 中 int 的定义），则逐位 与1 并累加的操作，可以不使用循环，以增加效率。

*/

func countBits(num int) []int {
	result := []int{0}
	for i := 1; i <= num; i++ {
		x := 0
		x += i & 1
		x += i >> 1 & 1
		x += i >> 2 & 1
		x += i >> 3 & 1
		x += i >> 4 & 1
		x += i >> 5 & 1
		x += i >> 6 & 1
		x += i >> 7 & 1
		x += i >> 8 & 1
		x += i >> 9 & 1
		x += i >> 10 & 1
		x += i >> 11 & 1
		x += i >> 12 & 1
		x += i >> 13 & 1
		x += i >> 14 & 1
		x += i >> 15 & 1
		x += i >> 16 & 1
		x += i >> 17 & 1
		x += i >> 18 & 1
		x += i >> 19 & 1
		x += i >> 20 & 1
		x += i >> 21 & 1
		x += i >> 22 & 1
		x += i >> 23 & 1
		x += i >> 24 & 1
		x += i >> 25 & 1
		x += i >> 26 & 1
		x += i >> 27 & 1
		x += i >> 28 & 1
		x += i >> 29 & 1
		x += i >> 30 & 1
		x += i >> 31 & 1
		x += i >> 32 & 1
		result = append(result, x)
	}
	return result
}

// 执行用时：4 ms, 在所有 Go 提交中击败了93.22% 的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了19.48% 的用户

// 这样写，大差不差，但是固定次数的循环，徒增消耗
// func countBits(num int) []int {
// 	result := []int{0}
// 	for i := 1; i <= num; i++ {
// 		x := 0
// 		for t := 0; t <= 32; t++ {
// 			x += i >> t & 1
// 		}
// 		result = append(result, x)
// 	}
// 	return result
// }

// 执行用时：8 ms, 在所有 Go 提交中击败了21.03% 的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了19.48% 的用户

func Test338() {
	fmt.Printf("%v\n", countBits(2)) // 0,1,1
	fmt.Printf("%v\n", countBits(5)) // 0,1,1,2,1,2
}
