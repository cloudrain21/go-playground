package gomap

func NewMultiSet() map[string]int {
	return map[string]int{}
}

func Insert(m map[string]int, val string) {
	m[val]++
}

func Erase(m map[string]int, val string) {
	if v, ok := m[val]; ok && v > 0 {
		m[val]--
	}
}

func Count(m map[string]int, val string) int {
	return m[val]
}

func String(m map[string]int) string {
	s := "{ "
	for key, _ := range m {
		s += key
		s += " "
	}
	s += "}"
	return s
}
