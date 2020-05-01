package main

import (
	"github.com/cloudrain21/go-playground/downloader/util"
	"log"
)

func main() {
	logger := util.GetInstance()

	fileSaver := util.FileUrlSaver{"urls.txt", []string{}}
	err := fileSaver.AcquireUrls()
	if err != nil {
		log.Fatal(err)
	}

	fileSaver.PrintUrls()

	logger.Println("xxx")
}
