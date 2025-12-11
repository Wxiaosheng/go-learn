package homework

import (
	"go-learn/utils"
	"math"
	"slices"
	"strconv"
)

func ExectTask01() {
	utils.Log("只出现一次的数字")
	utils.Expect(
		"[4, 1, 2, 1, 2]",
		strconv.Itoa(SingleNumber([]int{4, 1, 2, 1, 2})),
	)

	utils.Log("回文数")
	utils.Expect(
		"121",
		strconv.FormatBool(IsPalindrome(121)),
	)

	utils.Log("有效的括号")
	utils.Expect(
		"IsValid: ()[]{[()]}",
		strconv.FormatBool(IsValid("()[]{[()]}")),
	)
	utils.Expect(
		"IsValidV2: ()[]{[()]}",
		strconv.FormatBool(IsValidV2("()[]{[()]}")),
	)

	utils.Log("最长公共前缀")
	utils.Expect(
		`"flower", "flow", "flight"`,
		LongestCommonPrefix([]string{"flower", "flow", "flight"}),
	)

	utils.Log("加一 V1")
	result1 := PlusOne([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result1)
	result2 := PlusOne([]int{9, 9})
	utils.Expect("[9, 9]", result2)

	utils.Log("加一 V2")
	result21 := PlusOneV2([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result21)
	result22 := PlusOneV2([]int{9, 9})
	utils.Expect("[9, 9]", result22)

	utils.Log("加一 V3")
	result31 := PlusOneV3([]int{1, 2, 3, 4})
	utils.Expect("[1, 2, 3, 4]", result31)
	result32 := PlusOneV3([]int{9, 9})
	utils.Expect("[9, 9]", result32)

	utils.Log("删除有序数组中的重复项")
	dup1 := []int{1, 1, 2}
	utils.Expect(
		"[1,1,2]",
		RemoveDuplicates(dup1),
	)
	utils.Log(dup1)
	dup2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	utils.Expect(
		"[0,0,1,1,1,2,2,3,3,4]",
		RemoveDuplicates(dup2),
	)
	utils.Log(dup2)

	utils.Log("合并区间")
	utils.Expect(
		"[[1,3],[2,6],[8,10],[15,18]]",
		Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}),
	)
	utils.Expect(
		"[[1,4],[4,5]]",
		Merge([][]int{{1, 4}, {4, 5}}),
	)
	utils.Expect(
		"[[4,7],[1,4]]",
		Merge([][]int{{4, 7}, {1, 4}}),
	)

	utils.Log("两数之和")
	utils.Expect(
		"[2,7,11,15], target= 9",
		TwoSum([]int{2, 7, 11, 15}, 9),
	)
}

/*
 136. 只出现一次的数字：
    给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
    可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
    例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func SingleNumber(nums []int) int {
	cache := make(map[int]int)

	for _, num := range nums {
		cache[num]++
	}

	for key, val := range cache {
		if val == 1 {
			return key
		}
	}

	return 0
}

/*
*
回文数

	考察：数字操作、条件判断
	题目：判断一个整数是否是回文数
*/
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	_str := ""

	for i := len(str) - 1; i >= 0; i-- {
		_str += string(str[i])
	}

	return str == _str
}

// 进阶：你能不将整数转为字符串来解决这个问题吗？
// func isPalindromeV2(x int) bool {

// }

/*
有效的括号

	考察：字符串处理、栈的使用
	题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		每个右括号都有一个对应的相同类型的左括号。
*/
func IsValid(s string) bool {
	stack := make([]rune, 10)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			signe := string(left) + string(char)
			if signe != "()" && signe != "[]" && signe != "{}" {
				return false
			}
		}
	}

	return stack[len(stack)-1] == 0
}

func IsValidV2(s string) bool {
	stack := make([]rune, 10)
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pairs[char] != top {
				return false
			}
		}
	}

	return stack[len(stack)-1] == 0
}

/*
*
最长公共前缀

	考察：字符串处理、循环嵌套
	题目：查找字符串数组中的最长公共前缀
*/
func DiffString(str1 string, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	minStr := str1
	maxStr := str2
	if len(str1) > len(str2) {
		maxStr = str1
		minStr = str2
	}
	for index, char := range minStr {
		if char != rune(maxStr[index]) {
			return minStr[:index]
		}
	}
	return minStr
}
func LongestCommonPrefix(strs []string) string {
	long := strs[0]
	if len(strs) == 1 {
		return long
	}

	// 从第二个字符串开始比起
	for i := 1; i < len(strs); i++ {
		// 每次比较两个字符串
		long = DiffString(long, strs[i])
		if long == "" {
			return ""
		}
	}

	return long
}

/*加一 */
func PlusOne(digits []int) []int {
	plus := 1

	// 从后向前遍历
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += plus
			plus = 0
		}
	}

	if plus == 1 {
		return append([]int{1}, digits...)
	}

	return digits

}

func PlusOneV2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			return digits
		}
	}

	return append([]int{1}, digits...)
}

func PlusOneV3(digits []int) []int {
	// 用递归改实现
	return Plus(digits, len(digits)-1)
}

func Plus(digits []int, index int) []int {
	if index < 0 {
		return append([]int{1}, digits...)
	}
	if digits[index] == 9 {
		digits[index] = 0
		return Plus(digits, index-1)
	} else {
		digits[index] += 1
		return digits
	}
}

/* 删除有序数组中的重复项 */
func RemoveDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	// 双指针 i, j
	i := 0
	for j := i; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	return MergeLoop(intervals, 1)
}

/** 合并区间*/
func MergeLoop(intervals [][]int, index int) [][]int {

	// 如果遍历到最后一个元素，则返回结果
	if index >= len(intervals) {
		return intervals
	}

	loop := false
	result := [][]int{}

	result = append(result, intervals[0])
	// 一次确定一个范围
	for i := 1; i < len(intervals); i++ {
		overlap := IsOverlap(result[0], intervals[i])
		if overlap {
			// 合并
			min := math.Min(float64(result[0][0]), float64(intervals[i][0]))
			max := math.Max(float64(result[0][1]), float64(intervals[i][1]))
			result[0] = []int{int(min), int(max)}
		} else {
			// 复制无交集的区间
			result = append(result, intervals[i])
			loop = true
		}
	}

	// 如果还有交集，则继续遍历
	if loop {
		return MergeLoop(result, index+1)
	} else {
		return result
	}
}

// 如果 a 的开始结束都不在b中，则不重叠
func IsOverlap(a []int, b []int) bool {
	left := a[0] >= b[0] && a[0] <= b[1]
	right := a[1] >= b[0] && a[1] <= b[1]
	return left || right
}

/* 两数之和 */
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		// 从 i+1查起，后后续数组中是否存在 targt - nums[i]
		idx := slices.Index(nums[i+1:], target-nums[i])

		if idx != -1 {
			return []int{i, i + 1 + idx}
		}
	}
	return []int{}
}
