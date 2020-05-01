package downloader

//import (
//	"log"
//	"os"
//)
//
//type logger struct {
//	fileName string
//	*log.Logger
//}
//
//var logger *logger
//
//func GetLogger(logFileName string) (*logger) {
//	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return &logger{logFileName,
//		log.New(f, "my log prefix : ", log.Lshortfile)}
//}


import (
	"log"
	"os"
	"sync"
)

type logger struct {
	filename string
	*log.Logger
}

var logger *logger
var once sync.Once

// start loggeando
func GetInstance() *logger {
	once.Do(func() {
		logger = createLogger("mylogger.log")
	})
	return logger
}

func createLogger(fname string) *logger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &logger{
		filename: fname,
		Logger:   log.New(file, "My app Name ", log.Lshortfile),
	}
}