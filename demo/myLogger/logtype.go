package myLogger

import (
	"errors"
	"fmt"
	"github.com/dwttkgg/gocode/demo/runtime"
	"os"
	"path"
	"strings"
	"time"
)

type LogerLevel uint16

const (
	DEBUG LogerLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
	UNKNOWN
)

func parseLogLever(msg string) (LogerLevel, error) {
	msg = strings.ToLower(msg)
	switch msg {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}

}
func getLogLever(lv LogerLevel) string {

	switch lv {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "WARNING"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return "UNKNOWN"
	}

}

type fileType struct {
	Level      LogerLevel
	logPath    string
	logName    string
	fileObj    *os.File
	errFileObj *os.File
	maxSize    int64
}

func NewFileType(lv, lp, ln string, size int64) *fileType {
	loglv, err := parseLogLever(lv)
	if err != nil {
		panic(err)
	}
	fl := &fileType{
		Level:   loglv,
		logPath: lp,
		logName: ln,
		maxSize: size,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *fileType) initFile() (err error) {
	fullname := path.Join(f.logPath, f.logName)
	file, err := os.OpenFile(fullname+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开日志失败")
		return
	}
	errfile, err := os.OpenFile(fullname+"-err.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开日志失败")
		return
	}
	f.fileObj = file
	f.errFileObj = errfile
	return
}
func (f *fileType) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
func (l fileType) compareLv(level LogerLevel) bool {
	return l.Level <= level
}
func (f fileType) checkSize(file *os.File) bool {
	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info error")
		return false
	}
	return fileinfo.Size() > f.maxSize
}
func (f *fileType) log(lv LogerLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	n := time.Now()
	funcname, filename, line := runtime.Getinfo()
	if lv < ERROR {
		// 大于设置的文件大小,则需要进行切割
		if f.checkSize(f.fileObj) {
			f.fileObj = f.splitFile(f.fileObj)
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s %s %d] %s \n", n.Format("2006-01-02T15:04:05 -070000"), getLogLever(lv), filename, funcname, line, msg)
	}
	if lv >= ERROR {
		if f.checkSize(f.errFileObj) {
			f.errFileObj = f.splitFile(f.errFileObj)
		}
		fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s %s %d] %s \n", n.Format("2006-01-02T15:04:05 -070000"), getLogLever(lv), filename, funcname, line, msg)
	}
}
func (f *fileType) splitFile(fobj *os.File) *os.File {
	// 获取时间 并拼接到新名字
	nowStr := time.Now().Format("20060102150405000")
	logName := fobj.Name()
	newName := fmt.Sprintf("%s-%s", logName, nowStr)
	// 关闭旧文件,备
	fobj.Close()
	// 重命名文件
	os.Rename(logName, newName)
	// 然后打开一个新文件 将新打开的文件对象传给fileobj
	file, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
	}
	return file
}
func (f *fileType) Info(msg string, a ...interface{}) {
	if f.compareLv(INFO) {
		f.log(INFO, msg, a...)
	}
}
func (f *fileType) Debug(msg string, a ...interface{}) {
	if f.compareLv(DEBUG) {
		f.log(DEBUG, msg, a...)
	}
}
func (f *fileType) Error(msg string, a ...interface{}) {
	if f.compareLv(ERROR) {
		f.log(ERROR, msg, a...)
	}
}
func (f *fileType) Warn(msg string, a ...interface{}) {
	if f.compareLv(WARNING) {
		f.log(WARNING, msg, a...)
	}
}
func (f *fileType) Fatal(msg string, a ...interface{}) {
	if f.compareLv(FATAL) {
		f.log(FATAL, msg, a...)
	}
}
func (f *fileType) Trace(msg string, a ...interface{}) {
	if f.compareLv(TRACE) {
		f.log(TRACE, msg, a...)
	}
}
