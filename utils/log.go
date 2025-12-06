package utils

import "fmt"

func Log(v any) {
	fmt.Println()
	str, ok := v.(string)
	if ok {
		fmt.Println("------" + str + "------")
	} else {
		fmt.Println(v)
	}
}

func Expect(v string, r any) {
	fmt.Print(v + ", 实际结果为：")
	fmt.Println(r)
}
