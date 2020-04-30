package gomap

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("count test", func(t *testing.T) {
		m := map[rune]int{}
		str := "aabbccdd"
		Count(str, m)

		want := map[rune]int{'a':2, 'b':2, 'c':2, 'd':2}
		got := m
		assert.Equal(t, want, got)
	})
}

func ExampleCount() {
	m := map[rune]int{}
	str := "aabbccdd"

	Count(str, m)
	fmt.Println(m)
	// Output:
	// map[97:2 98:2 99:2 100:2]
}
