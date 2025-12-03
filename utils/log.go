package utils

import "fmt"

func Log(v string) {
	fmt.Println()
	fmt.Println("------" + v + "------")
}

func Expect(v string, r string) {
	fmt.Println(v+", 实际结果为：", r)
}
