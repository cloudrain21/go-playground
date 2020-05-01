package downloader

import "os"

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
	f, err := os.OpenFile(f.fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0755)

}
