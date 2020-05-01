package main

import (
	"fmt"
	"time"
)

func func1(str []string) {
	str = append(str, "ddd")
}

func func2(str *[]string) {
	*str = append(*str, "ddd")
}

func f1(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func f2(arr ...int) {
	fmt.Println(arr)
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("MyInt : %d", m)
}

func main() {
	strArr := []string{"aaa", "bbb", "ccc"}
	func1(strArr)
	fmt.Println(strArr)

	func2(&strArr)
	fmt.Println(strArr)

	arr := []int{1, 2, 3}
	f1(arr)

	f2(arr...)
	f2(1, 2, 3)

	v := MyInt(100)
	fmt.Println(v)

	CountDown(5)

	flag := 1
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("after 3 seconds...")
		flag = 0
	})

	for flag == 1 {
		time.Sleep(time.Second)
	}
	fmt.Println("timer stop...")
	timer.Stop()
}

func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}
