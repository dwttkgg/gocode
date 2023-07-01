package myLogger

//
//import (
//	"errors"
//	"fmt"
//	"github.com/dwttkgg/gocode/demo/runtime"
//	"strings"
//	"time"
//)
//
//type LogerLevel uint16
//
////const (
////	DEBUG LogerLevel = iota
////	TRACE
////	INFO
////	WARNING
////	ERROR
////	FATAL
////	UNKNOWN
////)
//
//type Logger struct {
//	Lever LogerLevel
//}
//
//// 构造函数
//func NewLogger(leverstr string) Logger {
//	lv, err := parseLogLever(leverstr)
//	if err != nil {
//		panic(err)
//	}
//	return Logger{Lever: lv}
//}
//func parseLogLever(msg string) (LogerLevel, error) {
//	msg = strings.ToLower(msg)
//	switch msg {
//	case "debug":
//		return DEBUG, nil
//	case "trace":
//		return TRACE, nil
//	case "info":
//		return INFO, nil
//	case "warning":
//		return WARNING, nil
//	case "error":
//		return ERROR, nil
//	case "fatal":
//		return FATAL, nil
//	default:
//		err := errors.New("无效的日志级别")
//		return UNKNOWN, err
//	}
//
//}
//func getLogLever(lv LogerLevel) string {
//
//	switch lv {
//	case 0:
//		return "DEBUG"
//	case 1:
//		return "TRACE"
//	case 2:
//		return "INFO"
//	case 3:
//		return "WARNING"
//	case 4:
//		return "ERROR"
//	case 5:
//		return "FATAL"
//	default:
//		return "UNKNOWN"
//	}
//
//}
//func (l Logger) compareLv(level LogerLevel) bool {
//	return l.Lever <= level
//}
//func log(f *fileType, lv LogerLevel, format string, a ...interface{}) {
//	msg := fmt.Sprintf(format, a...)
//	n := time.Now()
//	funcname, filename, line := runtime.Getinfo()
//	fmt.Fprintf(f.fileObj, "[%s] [%s] [%s %s %d] %s \n", n.Format("2006-01-02T15:04:05 -070000"), getLogLever(lv), filename, funcname, line, msg)
//	if lv >= ERROR {
//		fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s %s %d] %s \n", n.Format("2006-01-02T15:04:05 -070000"), getLogLever(lv), filename, funcname, line, msg)
//
//	}
//}
//func (l Logger) Info(msg string, a ...interface{}) {
//	if l.compareLv(INFO) {
//		log(INFO, msg, a...)
//	}
//}
//func (l Logger) Debug(msg string, a ...interface{}) {
//	if l.compareLv(DEBUG) {
//		log(DEBUG, msg, a...)
//	}
//}
//func (l Logger) Error(msg string, a ...interface{}) {
//	if l.compareLv(ERROR) {
//		log(ERROR, msg, a...)
//	}
//}
//func (l Logger) Warn(msg string, a ...interface{}) {
//	if l.compareLv(WARNING) {
//		log(WARNING, msg, a...)
//	}
//}
//func (l Logger) Fatal(msg string, a ...interface{}) {
//	if l.compareLv(FATAL) {
//		log(FATAL, msg, a...)
//	}
//}
