package io

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestWriteTo(t *testing.T) {
	t.Run("writeto test", func(t *testing.T) {
		str := []string{}
		for i := 0; i < 10; i++ {
			s := fmt.Sprintf("%d-%s", i, "aaa")
			str = append(str, s)
		}

		got := WriteTo(os.Stdout, str)
		var want error = nil
		assert.Equal(t, got, want)
	})

	t.Run("writeto test1", func(t *testing.T) {
		str := []string{}
		for i := 0; i < 3; i++ {
			s := fmt.Sprintf("%d-%s", i, "aaa")
			str = append(str, s)
		}

		want := ArrString{[]string{"0-aaa", "1-aaa", "2-aaa"}}
		var a ArrString
		a.WriteTo(str)
		got := a
		assert.Equal(t, got, want)
	})
}
