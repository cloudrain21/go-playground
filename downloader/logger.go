package downloader

import (
	"log"
	"os"
)

type logger struct {
	fileName string
	*log.Logger
}

func GetLogger(logFileName string) (*logger) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return &logger{logFileName,
		log.New(f, "my log prefix : ", log.Lshortfile)}
}
