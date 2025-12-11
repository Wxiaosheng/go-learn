package homework

import (
	"go-learn/utils"
	"sync"
	"time"
)

func ExectTask02() {
	utils.Log("指针")
	num := 2
	AddTen(&num)
	utils.Expect("数字 2 修改后为", num)

	utils.Log("切片乘以 2")
	nums := []int{2, 4, 6}
	MultiplyByTwo(&nums)
	utils.Expect("切片乘以 2 后为", nums)

	utils.Log("测试 go 协程")
	TestGo()

	utils.Log("测试任务调度器")
	result := TaskScheduler([]func(){t1, t2, t3})
	utils.Expect("任务执行时间：", result)
}

/******✅指针*******/
/*
1. 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func AddTen(num *int) {
	*num += 10
}

/*
2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func MultiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

/******✅Goroutine******/
/*
	题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	考察点 ： go 关键字的使用、协程的并发执行。
*/
func TestGo() {
	var wg = sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			utils.Expect("第一个协程：", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			utils.Expect("第二个协程：", i)
		}
	}()

	wg.Wait()
}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func TaskScheduler(tasks []func()) []time.Duration {

	wg := sync.WaitGroup{}

	results := make([]time.Duration, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		go func(i int, task func()) {
			defer wg.Done()

			start := time.Now()
			task()
			elapsed := time.Since(start)
			results[i] = elapsed
		}(i, task)
	}

	wg.Wait()
	return results
}
func t1() {
	time.Sleep(1 * time.Second) // 模拟耗时操作
}
func t2() {
	time.Sleep(2 * time.Second) // 模拟耗时操作
}
func t3() {
	time.Sleep(3 * time.Second) // 模拟耗时操作
}
