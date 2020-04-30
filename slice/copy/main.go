package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cap = 100000000

func getSum(arr []int) (arrsum int) {
	arrsum = 0
	for _, v := range arr {
		arrsum += v
	}
	return
}

func main() {
	s1 := make([]int, cap)
	s2 := make([]int, cap)
	s3 := make([]int, cap)
	var s4 []int

	for i := 0; i<cap; i++ {
		s1[i] = rand.Intn(100)
	}
	sum1 := getSum(s1)

	start := time.Now()
	copy(s2, s1)
	elapse := time.Since(start)
	fmt.Println(elapse)
	sum2 := getSum(s2)

	start = time.Now()
	for i, v := range s1 {
		s3[i] = v
	}
	elapse = time.Since(start)
	fmt.Println(elapse)
	sum3 := getSum(s3)

	start = time.Now()
	s4 = append(s4, s1...)
	//for _, v := range s1 {
	//	s4 = append(s4, v)
	//}
	elapse = time.Since(start)
	fmt.Println(elapse)
	sum4 := getSum(s4)

	fmt.Println(sum1, sum2, sum3, sum4)
}
