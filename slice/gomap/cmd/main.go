package main

import (
	"fmt"
	"sort"
)

func Count(str string, m map[rune]int) {
	for _, v := range str {
		m[v]++
	}
}

func SortMap(m map[int]int) {
	var keys sort.IntSlice
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)
}

func main() {
	m := map[rune]int{}
	str := "zhdhebdadssfdsd"

	Count(str, m)
	for k, v := range m {
		fmt.Println(string(k), v)
	}

	m1 := map[int]int{}
	for _, k := range str {
		m1[int(k)]++
	}
	SortMap(m1)
	fmt.Println(m1)
}
