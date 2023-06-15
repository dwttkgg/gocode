package utils

import (
	"logagent/kafka"
	"logagent/tailf"
	"time"

	"github.com/astaxie/beego/logs"
)

func ServerRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err := sendToKakfa(msg)
		if err != nil {
			logs.Error("msg send kafka is error :", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKakfa(msg *tailf.TextMsg) (err error) {
	//logs.Debug("read msg:%s,topic:%s", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
