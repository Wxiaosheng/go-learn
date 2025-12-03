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

}
