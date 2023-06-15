package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type Logconfig struct {
	KafkaAddr string
	EsAddr    string
	LogPath   string
	LogLevel  string
}

var logConfig *Logconfig

func initConfig(ConfType string, filename string) (err error) {
	conf, err := config.NewConfig(ConfType, filename)
	if err != nil {
		fmt.Println("config.NewConfig is error:", err)
	}
	logConfig = &Logconfig{}
	logConfig.LogLevel = conf.String("logs:log_level")
	if len(logConfig.LogLevel) == 0 {
		logConfig.LogLevel = "debug"
	}
	logConfig.LogPath = conf.String("logs::log_path")
	if len(logConfig.LogPath) == 0 {
		logConfig.LogPath = "./logs"
	}
	logConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(logConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka address")
		return
	}
	logConfig.EsAddr = conf.String("etcd::addr")
	if len(logConfig.EsAddr) == 0 {
		err = fmt.Errorf("invalid elasticsearch address")
		return
	}

	return
}
