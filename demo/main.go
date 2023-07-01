package main

import (
	"errors"
	"github.com/dwttkgg/gocode/demo/myLogger"
)

func main() {
	//lv := myLogger.NewLogger("info")
	log := myLogger.NewFileType("trace", "./", "digest", 1024*1024)
	for {
		log.Debug("debug log")
		log.Fatal("fatal log")
		log.Info("info log")
		log.Error("error log,err is %v", errors.New("out of memory"))
		log.Fatal("fatal log")
		//time.Sleep(3 * time.Second)
	}
}
