package io

import (
	"fmt"
	"io"
)

type ArrString struct {
	Str []string
}

func (a *ArrString) WriteTo(lines []string) error {
	for _, s := range lines {
		a.Str = append(a.Str, s)
	}
	return nil
}

func WriteTo(w io.Writer, lines []string) error {
	for i, str := range lines {
		_, err := fmt.Fprintf(w, "%d : %s\n", i, str)
		if err != nil {
			return err
		}
	}
	return nil
}
