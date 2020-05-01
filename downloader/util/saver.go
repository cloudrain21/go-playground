package util

import (
	"bufio"
	"log"
	"os"
)

type UrlSaver interface {
	AcquireUrls() ([]string, error)
}

type FileUrlSaver struct {
	fileName string
	urls     []string
}

type MemUrlSaver struct {
	urls []string
}

func (f FileUrlSaver) AcquireUrls() error {
	file, err := os.OpenFile(f.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f.urls = append(f.urls, scanner.Text())
	}

	return nil
}

func (f FileUrlSaver) PrintUrls() {
	for _, url := range f.urls {
		log.Println(url)
	}
}
