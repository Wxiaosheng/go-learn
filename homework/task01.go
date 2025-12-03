package homework

import (
	"strconv"
)

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

/*
*
 */
// func plusOne(digits []int) []int {

// }
