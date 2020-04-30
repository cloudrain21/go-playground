package gomap

func Count(str string, m map[rune]int) {
	for _, v := range str {
		m[v]++
	}
}
