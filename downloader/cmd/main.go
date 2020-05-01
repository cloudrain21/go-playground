package main

import (
	"bufio"
	"github.com/cloudrain21/go-playground/downloader/util"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var logger = util.GetInstance()

func main() {
	urlSaver := util.NewFileUrlSaver("urls.txt")
	err := urlSaver.AcquireUrls()
	if err != nil {
		log.Fatal(err)
	}

	logger.Println(urlSaver.GetUrls())
	download(urlSaver.GetUrls())
}

func download(urls []string) {
	for _, url := range urls {
		logger.Printf("start download... : %s\n", url)

		res, err := http.Get(url)
		if err != nil {
			logger.Fatal(err)
		}
		defer res.Body.Close()

		filename := "image/" + filepath.Base(url)

		f, err := os.Create(filename)
		if err != nil {
			logger.Fatal(err)
		}
		defer f.Close()

		r := bufio.NewReader(res.Body)
		len, err := io.Copy(f, r)
		if err != nil {
			logger.Fatal(err)
		}

		logger.Printf("end download... : %s len(%d)\n", url, len)
	}
}
