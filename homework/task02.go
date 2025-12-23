package homework

import (
	"fmt"
	"go-learn/utils"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

func ExectTask02() {
	utils.Log("测试 ✅指针")
	num := 2
	AddTen(&num)
	utils.Expect("数字 2 修改后为", num)

	utils.Log("切片乘以 2")
	nums := []int{2, 4, 6}
	MultiplyByTwo(&nums)
	utils.Expect("切片乘以 2 后为", nums)

	utils.Log("测试 ✅ go 协程")
	TestGo()

	utils.Log("测试 ✅任务调度器")
	result := TaskScheduler([]func(){t1, t2, t3})
	utils.Expect("任务执行时间：", result)

	utils.Log("测试 ✅面向对象")
	r := Rectangle{
		width:  2.0,
		height: 3.0,
	}
	utils.Expect("Rectangle.Area: ", r.Area())
	utils.Expect("Rectangle.Perimeter: ", r.Perimeter())

	c := Circle{radio: 3}
	utils.Expect("Circle.Area: ", c.Area())
	utils.Expect("Circle.Perimeter: ", c.Perimeter())

	utils.Log("测试 ✅面向对象 - 组合")
	e := Employee{
		EmployeeID: 1,
		Person: Person{
			Name: "victree",
			Age:  31,
		},
	}
	utils.Expect("Person.PrintInfo: ", "")
	e.PrintInfo()

	utils.Log("测试 ✅Channel - TestChanne")
	TestChannel()
	utils.Log("测试 ✅Channel - TestChannel2")
	TestChannel2()

	utils.Log("测试 ✅锁机制")
	utils.Expect("测试 ✅锁机制 - TestLock1", TestLock1())
	utils.Expect("测试 ✅原子操作 - TestAtomic", TestAtomic())
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

/*
✅面向对象
1、题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。

	在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	width  float32
	height float32
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float32 {
	return (r.width + r.height) * 2
}

type Circle struct {
	radio float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radio * c.radio
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radio
}

/*
	2、题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
					为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	考察点 ：组合的使用、方法接收者。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d \n", e.Name, e.Age, e.EmployeeID)
}

/*
✅Channel
1、题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

	考察点 ：通道的基本使用、协程间通信。
*/
func TestChannel() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println("Received: ", num)
	}

	fmt.Println("All numbers received.")
}

/*
2、题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func TestChannel2() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("Sender: 准备发送 ", i)
			ch <- i // 这里会阻塞，直到有人接收！
			fmt.Println("Sender: 成功发送 ", i)
		}

		close(ch)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Receiver: 开始接收")
	for num := range ch {
		fmt.Println("Received: ", num)
	}

	fmt.Println("All numbers received.")
}

/*
✅锁机制
1、题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/
func TestLock1() int {
	num := 0
	wg := sync.WaitGroup{}
	m := sync.Mutex{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m.Lock()
				num++
				m.Unlock()
			}
		}()
	}

	wg.Wait()
	return num
}

/*
2、题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

	考察点 ：原子操作、并发数据安全。

	var counter int64
	// 原子递增 1
	atomic.AddInt64(&counter, 1)
*/
func TestAtomic() int64 {
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()
	return count
}
