package main

import (
	"github.com/cloudrain21/go-playground/downloader/util"
	"log"
)

func main() {
	logger := util.GetInstance()

	urlSaver := util.NewFileUrlSaver("urls.txt")
	err := urlSaver.AcquireUrls()
	if err != nil {
		log.Fatal(err)
	}

	download(urlSaver)

	logger.Println("ttttt")
}

func download(urlSaver util.UrlSaver) {
}
