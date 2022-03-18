package L_0599

import (
	"fmt"
)

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/

// 599. 两个列表的最小索引总和

// 假设 Andy 和 Doris 想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。
// 你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设答案总是存在。

// 1 <= list1.length, list2.length <= 1000
// 1 <= list1[i].length, list2[i].length <= 30
// list1[i] 和 list2[i] 由空格 ' ' 和英文字母组成。
// list1 的所有字符串都是 唯一 的。
// list2 中的所有字符串都是 唯一 的。

func findRestaurant(list1 []string, list2 []string) []string {
	results := []string{}
	indexSum := 2001
	listMap := map[string]int{}
	for i := 0; i < len(list1); i++ {
		listMap[list1[i]] = i
	}
	for i := 0; i < len(list2); i++ {
		if index, ok := listMap[list2[i]]; ok {
			if indexSum > index+i {
				indexSum = index + i
				results = []string{list2[i]}
			} else if indexSum == index+i {
				results = append(results, list2[i])
			}
		}
	}

	return results
}

func Test599() {
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"})) // Shogun
	fmt.Println(findRestaurant(
		[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		[]string{"KFC", "Shogun", "Burger King"})) // Shogun

}
