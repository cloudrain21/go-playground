package main

import "fmt"

const num = 3

func main() {
	s1 := make([]string, num)
	var s2 []string
	i := 0

	for i < num {
		s1[i] = fmt.Sprintf("%s-%d", "aaaa", i)
		i++
	}
	MyPrint(s1)

	for i :=0; i<num; i++ {
		s2 = append(s2, s1[i])
	}
	MyPrint(s2)

	s3 := make([]string, num)
	copy(s3, s2)
	MyPrint(s3)

	var s4 []string
	s4 = append(s4, s3...)
	MyPrint(s4)

	s5 := []string{"aaa", "bbb", "ccc"}
	MyPrint(s5)
}

func MyPrint(sarr []string) {
	fmt.Println("len : ", len(sarr))
	for i, sitem := range sarr {
		fmt.Println(i, " : ", sitem)
	}
}
