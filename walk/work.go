package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("/Users/kakao_ent/type", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
	fmt.Println(err)
}
