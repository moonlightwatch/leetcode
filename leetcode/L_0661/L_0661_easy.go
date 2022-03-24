package L_0661

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/image-smoother/

// 661. 图片平滑器

// 图像平滑器 是大小为 3 x 3 的过滤器，用于对图像的每个单元格平滑处理，平滑处理后单元格的值为该单元格的平均灰度。

// 每个单元格的  平均灰度 定义为：该单元格自身及其周围的 8 个单元格的平均值，结果需向下取整。（即，需要计算蓝色平滑器中 9 个单元格的平均值）。

// 如果一个单元格周围存在单元格缺失的情况，则计算平均灰度时不考虑缺失的单元格（即，需要计算红色平滑器中 4 个单元格的平均值）。

// m == img.length
// n == img[i].length
// 1 <= m, n <= 200
// 0 <= img[i][j] <= 255

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])
	result := make([][]int, m)
	x := 0

	// 直接遍历，通过判断，处理边缘
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			x = 0
			result[i][j] += img[i][j]
			x++
			if j > 0 {
				result[i][j] += img[i][j-1]
				x++
			}
			if j < n-1 {
				result[i][j] += img[i][j+1]
				x++
			}
			if i > 0 {
				if j > 0 {
					result[i][j] += img[i-1][j-1]
					x++
				}
				result[i][j] += img[i-1][j]
				x++
				if j < n-1 {
					result[i][j] += img[i-1][j+1]
					x++
				}
			}
			if i < m-1 {
				if j > 0 {
					result[i][j] += img[i+1][j-1]
					x++
				}
				result[i][j] += img[i+1][j]
				x++
				if j < n-1 {
					result[i][j] += img[i+1][j+1]
					x++
				}
			}
			// fmt.Printf("(%d,%d) %d / %d = %f\n", i, j, result[i][j], x, math.Floor(float64(result[i][j])/float64(x)))
			result[i][j] = int(math.Floor(float64(result[i][j]) / float64(x)))
		}
	}
	return result
}

// func imageSmoother(img [][]int) [][]int {
// 	m := len(img)
// 	n := len(img[0])
// 	result := make([][]int, m)

// 	// 构建结果矩阵
// 	for i := 0; i < m; i++ {
// 		result[i] = make([]int, n)
// 	}

// 	// 处理极限情况
// 	if m == 1 && n > 1 {
// 		// 先算两头
// 		result[0][0] += int(math.Floor(float64(img[0][0]+img[0][1]) / 2))
// 		result[0][n-1] += int(math.Floor(float64(img[0][n-1]+img[0][n-2]) / 2))
// 		// 再算当间
// 		for j := 1; j < n-1; j++ {
// 			result[0][j] += int(math.Floor(float64(img[0][j-1]+img[0][j]+img[0][j+1]) / 3))
// 		}
// 		return result
// 	} else if n == 1 && m > 1 {
// 		// 先算两头
// 		result[0][0] += int(math.Floor(float64(img[0][0]+img[1][0]) / 2))
// 		result[m-1][0] += int(math.Floor(float64(img[m-1][0]+img[m-2][0]) / 2))
// 		// 再算当间
// 		for j := 1; j < m-1; j++ {
// 			result[j][0] += int(math.Floor(float64(img[j][0]+img[j-1][0]+img[j+1][0]) / 3))
// 		}
// 		return result
// 	} else if n == 1 && m == 1 {
// 		// 不用算
// 		return img
// 	}

// 	// 先计算四角
// 	result[0][0] += int(math.Floor(float64(img[0][0]+img[0][1]+img[1][0]+img[1][1]) / 4))
// 	result[0][n-1] += int(math.Floor(float64(img[0][n-1]+img[0][n-2]+img[1][n-1]+img[1][n-2]) / 4))
// 	result[m-1][0] += int(math.Floor(float64(img[m-1][0]+img[m-1][1]+img[m-2][0]+img[m-2][1]) / 4))
// 	result[m-1][n-1] += int(math.Floor(float64(img[m-1][n-1]+img[m-1][n-2]+img[m-2][n-1]+img[m-2][n-2]) / 4))

// 	// 计算四个边缘
// 	// 左右边缘
// 	for i := 1; i < m-1; i++ {
// 		result[i][0] += int(math.Floor(float64(img[i-1][0]+img[i-1][1]+img[i][0]+img[i][1]+img[i+1][0]+img[i+1][1]) / 6))
// 		result[i][n-1] += int(math.Floor(float64(img[i-1][n-1]+img[i-1][n-2]+img[i][n-1]+img[i][n-2]+img[i+1][n-1]+img[i+1][n-2]) / 6))
// 	}
// 	// 上下边缘
// 	for i := 1; i < n-1; i++ {
// 		result[0][i] += int(math.Floor(float64(img[0][i-1]+img[1][i-1]+img[0][i]+img[1][i]+img[0][i+1]+img[1][i+1]) / 6))
// 		result[m-1][i] += int(math.Floor(float64(img[m-1][i-1]+img[m-2][i-1]+img[m-1][i]+img[m-2][i]+img[m-1][i+1]+img[m-2][i+1]) / 6))
// 	}

// 	// 计算内部
// 	for i := 1; i < m-1; i++ {
// 		for j := 1; j < n-1; j++ {
// 			result[i][j] += img[i-1][j-1]
// 			result[i][j] += img[i-1][j]
// 			result[i][j] += img[i-1][j+1]
// 			result[i][j] += img[i][j-1]
// 			result[i][j] += img[i][j]
// 			result[i][j] += img[i][j+1]
// 			result[i][j] += img[i+1][j-1]
// 			result[i][j] += img[i+1][j]
// 			result[i][j] += img[i+1][j+1]
// 			result[i][j] = int(math.Floor(float64(result[i][j]) / 9))
// 		}
// 	}
// 	return result
// }

func Test661() {
	fmt.Println(imageSmoother([][]int{{7}, {9}, {6}})) // [[8],[7],[7]]
	fmt.Println(imageSmoother([][]int{{1}}))
	fmt.Println(imageSmoother([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))                  // [[0, 0, 0],[0, 0, 0], [0, 0, 0]]
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}})) //  [[137,141,137],[141,138,141],[137,141,137]]
}
