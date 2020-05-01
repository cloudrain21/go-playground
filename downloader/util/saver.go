package util

import (
	"log"
	"os"
)

type UrlSaver interface {
	AcquireUrls() ([]string, error)
}

type FileUrlSaver struct {
	fileName string
	urls []string
}

type MemUrlSaver struct {
	urls []string
}

func (f FileUrlSaver)AcquireUrls() ([]string, error) {
	_, err := os.OpenFile(f.fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return nil, nil
}
