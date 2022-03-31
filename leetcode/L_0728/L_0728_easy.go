package L_0728

import "fmt"

// https://leetcode-cn.com/problems/self-dividing-numbers/

// 728. 自除数

// 自除数 是指可以被它包含的每一位数整除的数。

//     例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。

// 自除数 不允许包含 0 。

// 给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ { // 逐个数字遍历
		ok := true                                  // 表示当前 i 是不是 自除数，默认是，接下来进行计算
		for temp := i; temp > 0; temp = temp / 10 { // 提取每一位数（反复除以10 ， 商与10取余，即为每一位数）
			if temp%10 == 0 || i%(temp%10) != 0 { // 包含 0 的，或者 自己除后不为 0 的，不是自除数
				ok = false
				break
			}
		}
		if ok { // 如果是 自除数 ，则加入结果集
			result = append(result, i)
		}
	}
	return result
}

func Test728() {
	fmt.Println(selfDividingNumbers(1, 22))  // 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22
	fmt.Println(selfDividingNumbers(47, 85)) // 48,55,66,77
}
