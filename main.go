package main

import (
	"go-learn/homework"
	"go-learn/utils"
	"strconv"
)

func main() {

	utils.Log("只出现一次的数字")
	utils.Expect(
		"[4, 1, 2, 1, 2]",
		strconv.Itoa(homework.SingleNumber([]int{4, 1, 2, 1, 2})),
	)

	utils.Log("回文数")
	utils.Expect(
		"121",
		strconv.FormatBool(homework.IsPalindrome(121)),
	)

	utils.Log("有效的括号")
	utils.Expect(
		"IsValid: ()[]{[()]}",
		strconv.FormatBool(homework.IsValid("()[]{[()]}")),
	)
	utils.Expect(
		"IsValidV2: ()[]{[()]}",
		strconv.FormatBool(homework.IsValidV2("()[]{[()]}")),
	)

	utils.Log("最长公共前缀")
	utils.Expect(
		`"flower", "flow", "flight"`,
		homework.LongestCommonPrefix([]string{"flower", "flow", "flight"}),
	)

	utils.Log("加一 V1")
	result1 := homework.PlusOne([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result1)
	result2 := homework.PlusOne([]int{9, 9})
	utils.Expect("[9, 9]", result2)

	utils.Log("加一 V2")
	result21 := homework.PlusOneV2([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result21)
	result22 := homework.PlusOneV2([]int{9, 9})
	utils.Expect("[9, 9]", result22)

	utils.Log("加一 V3")
	result31 := homework.PlusOneV3([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result31)
	result32 := homework.PlusOneV3([]int{9, 9})
	utils.Expect("[9, 9]", result32)

	utils.Log("删除有序数组中的重复项")
	dup1 := []int{1, 1, 2}
	utils.Expect(
		"[1,1,2]",
		homework.RemoveDuplicates(dup1),
	)
	utils.Log(dup1)
	dup2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	utils.Expect(
		"[0,0,1,1,1,2,2,3,3,4]",
		homework.RemoveDuplicates(dup2),
	)
	utils.Log(dup2)

	utils.Log("合并区间")
	utils.Expect(
		"[[1,3],[2,6],[8,10],[15,18]]",
		homework.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
	)
	utils.Expect(
		"[[1,4],[4,5]]",
		homework.Merge([][]int{{1, 4}, {4, 5}}),
	)
	utils.Expect(
		"[[4,7],[1,4]]",
		homework.Merge([][]int{{4, 7}, {1, 4}}),
	)

	utils.Log("两数之和")
	utils.Expect(
		"[2,7,11,15], target= 9",
		homework.TwoSum([]int{2, 7, 11, 15}, 9),
	)

}
