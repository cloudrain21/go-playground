package main

import "github.com/cloudrain21/go-playground/downloader/util"

func main() {
	logger := util.GetInstance()

	fileSaver := &util.FileUrlSaver{
		"urls.txt",
		[]string{},
	}

	fileSaver.AcquireUrls()
}
