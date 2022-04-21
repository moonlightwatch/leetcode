package L_0396

import "fmt"

// https://leetcode-cn.com/problems/rotate-function/

// 396. 旋转函数

// 给定一个长度为 n 的整数数组 numss 。

// 假设 arrk 是数组 numss 顺时针旋转 k 个位置后的数组，我们定义 numss 的 旋转函数  F 为：

//     F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]

// 返回 F(0), F(1), ..., F(n-1)中的最大值 。

// 生成的测试用例让答案符合 32 位 整数。

// 看数据规模，暴力求解不可取。
// 所以，找规律：

// 设 : n = 6
// F(0) = 0*nums[0] + 1*nums[1] + 2*nums[2] + 3*nums[3] + 4*nums[4] + 5*nums[5]
// F(1) = 1*nums[0] + 2*nums[1] + 3*nums[2] + 4*nums[3] + 5*nums[4] + 0*nums[5] = F(0) + sum(nums) - (n)*nums[n-1]
// F(2) = 2*nums[0] + 3*nums[1] + 4*nums[2] + 5*nums[3] + 0*nums[4] + 1*nums[5] = F(1) + sum(nums) - (n)*nums[n-2]
// F(3) = 3*nums[0] + 4*nums[1] + 5*nums[2] + 0*nums[3] + 1*nums[4] + 2*nums[5] = F(2) + sum(nums) - (n)*nums[n-3]
// F(4) = 4*nums[0] + 5*nums[1] + 0*nums[2] + 1*nums[3] + 2*nums[4] + 3*nums[5] = F(3) + sum(nums) - (n)*nums[n-4]
// F(5) = 5*nums[0] + 0*nums[1] + 1*nums[2] + 2*nums[3] + 3*nums[4] + 4*nums[5] = F(4) + sum(nums) - (n)*nums[n-5]

// 获得递推式： F(k) + sum(nums) - (n)*nums[n-k]

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0  // 记录 sum(nums)
	temp := 0 // 记录 F(k) 的值
	for i := 0; i < n; i++ {
		sum += nums[i]      // 累加 nums
		temp += i * nums[i] // 计算 F(0)
	}
	max := temp
	for k := 1; k < n; k++ {
		temp = temp + sum - n*nums[n-k]
		if temp > max {
			max = temp
		}
	}

	return max
}

func Test396() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6})) // 26
	fmt.Println(maxRotateFunction([]int{100}))        // 0
}
