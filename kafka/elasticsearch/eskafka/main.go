package main

import (
	"eskakfa/uitls"
	"github.com/astaxie/beego/logs"
)

func main() {
	logConfig, err := uitls.InitConfig("ini", "../../logagent/conf/log_transfer.conf")
	if err != nil {
		panic(err)
	}
	//fmt.Println(logConfig)
	err = uitls.InitLogger(logConfig.LogPath, logConfig.LogLevel)
	if err != nil {
		panic(err)
	}
	logs.Debug("init logger succ")

	err = uitls.InitKafka(logConfig.KafkaAddr, logConfig.Topic)
	if err != nil {
		logs.Error("init kafka failed ,err:", err)
	}
	logs.Debug("init kafka succ")
	err = uitls.InitEs(logConfig.EsAddr)
	if err != nil {
		logs.Error("init elasticsearch failed,err:", err)
	}
	logs.Debug("init elasticsearch succ")

	err = uitls.Run()
	if err != nil {
		logs.Error("server run failed ,err:", err)
	}
	logs.Debug("server run succ")
}
