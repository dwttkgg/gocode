package tailf

import (
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct {
	LogPath string `json:"log_path"`
	Topic   string `json:"topic"`
}
type TailObj struct {
	Tailobj  *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}
type TextMsg struct {
	Msg   string
	Topic string
}
type TailObjMgr struct {
	Tailobjs []*TailObj
	msgchan  chan *TextMsg
	lcok     sync.Mutex
}

func UpdateConf(coll []CollectConf) (err error) {
	tailobjs.lcok.Lock()
	defer tailobjs.lcok.Unlock()
	for _, conf := range coll {
		var isRunning = true
		for _, v := range tailobjs.Tailobjs {
			if conf.LogPath == v.conf.LogPath {
				isRunning = true
				break
			}
		}
		if isRunning {
			continue
		}
		createNewTask(conf)
	}
	var tailOb []*TailObj
	for _, v := range tailobjs.Tailobjs {
		v.status = StatusDelete
		for _, oneConf := range coll {
			if oneConf.LogPath == v.conf.LogPath {

				v.status = StatusNormal
				break
			}
		}
		if v.status == StatusDelete {
			v.exitChan <- 1
		}
		tailOb = append(tailOb, v)
	}
	tailobjs.Tailobjs = tailOb
	return
}

var tailobjs *TailObjMgr

func GetOneLine() (msg *TextMsg) {
	msg = <-tailobjs.msgchan
	return
}
func createNewTask(conf CollectConf) {

	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(conf.LogPath, config)
	if err != nil {
		logs.Error("connect filename[%s] is failed .err:%v", conf.LogPath, err)
		return
	}
	obj.Tailobj = tails
	tailobjs.Tailobjs = append(tailobjs.Tailobjs, obj)
	go readFormTail(obj)
}
func InitTail(conf []CollectConf, chansize int) (err error) {
	if len(conf) == 0 {
		logs.Error("invalid config for log collect,conf:%v", conf)
		return
	}
	tailobjs = &TailObjMgr{
		msgchan: make(chan *TextMsg, chansize),
	}
	for _, v := range conf {
		createNewTask(v)
		//obj := &TailObj{
		//	conf: v,
		//}
		//config := tail.Config{
		//	ReOpen:    true,                                 // 重新打开
		//	Follow:    true,                                 // 是否跟随
		//	Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		//	MustExist: false,                                // 文件不存在不报错
		//	Poll:      true,
		//}
		//tails, err := tail.TailFile(v.LogPath, config)
		//if err != nil {
		//	logs.Error("tail file failed:%v", conf)
		//	return err
		//}
		//obj.Tailobj = tails
		//tailobjs.Tailobjs = append(tailobjs.Tailobjs, obj)
		//go readFormTail(obj)
	}
	return
}

func readFormTail(tail *TailObj) {
	for true {
		select {
		case line, ok := <-tail.Tailobj.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tail.Tailobj.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			msg := &TextMsg{
				Msg:   line.Text,
				Topic: tail.conf.Topic,
			}
			tailobjs.msgchan <- msg
		case <-tail.exitChan:
			logs.Warn("tail obj will exited ,conf:%v", tail.conf)
			return

		}
	}
}
