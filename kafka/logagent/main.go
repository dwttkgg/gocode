package main

import (
	"fmt"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/tailf"
	"logagent/utils"

	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "./conf/logagent.conf"
	typename := "ini"
	//  加载配置文件信息
	appconf, err := utils.Loadconf(typename, filename)
	if err != nil {
		fmt.Println("load conf failed,err:", err)
		panic("load conf error")
		return
	}
	logs.Debug("load conf  success")
	//  初始化日志
	err = utils.InitLogger(appconf)
	if err != nil {
		fmt.Println("init Logger failed,err:", err)
		panic("init Logger error")
		return
	}
	utils.InitIp()
	logs.Debug("init Logger success , config:", appconf)
	// 初始化读取组件
	err = tailf.InitTail(appconf.CotConf, appconf.Chan_size)
	if err != nil {
		logs.Error("init Tail is error:", err)
	}
	logs.Debug("init Tail  success")

	// 初始化kafka
	err = kafka.InitKafka(appconf.KafkaAddr)
	if err != nil {
		logs.Error("init kafka is error:", err)
	}
	logs.Debug("init kafka success")
	// 初始化etcd
	err = etcd.InitEtcd(appconf.Etcdaddr, appconf.Etcdkey)
	if err != nil {
		logs.Error("init etcd is error:", err)
	}
	logs.Debug("init etcd success")

	logs.Debug("initialize all  success")
	err = utils.ServerRun()
	if err != nil {
		logs.Error("server run  is error:", err)
	}
	logs.Info("program exited\n")
}
