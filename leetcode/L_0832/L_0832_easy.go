package L_0832

import "fmt"

// https://leetcode-cn.com/problems/flipping-an-image/

// 832. 翻转图像

// 给定一个 n x n 的二进制矩阵 image ，先 水平 翻转图像，然后 反转 图像并返回 结果 。

// 水平翻转图片就是将图片的每一行都进行翻转，即逆序。

//     例如，水平翻转 [1,1,0] 的结果是 [0,1,1]。

// 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。

//     例如，反转 [0,1,1] 的结果是 [1,0,0]。

/*

思路：

水平翻转的同时各项与1异或，可以达到县水平翻转在逐个进行 0-1 反转

*/

func flipAndInvertImage(A [][]int) [][]int {
	l := len(A)
	for i, j := 0, l-1; i <= j; i, j = i+1, j-1 {
		for k := 0; k < l; k++ {
			A[k][i], A[k][j] = A[k][j]^1, A[k][i]^1
		}
	}
	return A
}

// 执行用时：4 ms, 在所有 Go 提交中击败了82.14% 的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了61.90% 的用户

func Test832() {
	fmt.Printf("%v\n", flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Printf("%v\n", flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}
