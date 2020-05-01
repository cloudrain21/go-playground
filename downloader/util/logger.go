package util

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

type mylogger struct {
	filename string
	*log.Logger
}

var logger *mylogger
var once sync.Once

func GetInstance() *mylogger {
	once.Do(func() {
		logger = createLogger("mylogger.log")
	})
	return logger
}

func createLogger(fname string) *mylogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &mylogger{
		filename: fname,
		Logger:   log.New(file, "My app Name ", log.Lshortfile),
	}
}