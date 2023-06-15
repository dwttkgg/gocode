package utils

import (
	"errors"
	"fmt"
	"logagent/tailf"

	"github.com/astaxie/beego/config"
)

var appconf *Config

type Config struct {
	LogPath   string
	LogLevel  string
	Chan_size int
	KafkaAddr string
	CotConf   []tailf.CollectConf
	Etcdaddr  string
	Etcdkey   string
}

func loadCollectConf(conf config.Configer, appconf *Config) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		errors.New("invalid collect:log_path")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		errors.New("invalid collect:topic")
		return
	}
	appconf.CotConf = append(appconf.CotConf, cc)
	return
}
func Loadconf(typename string, filename string) (appconf *Config, err error) {
	conf, err := config.NewConfig(typename, filename)
	if err != nil {
		fmt.Println("config.NewConfig is error:", err)
	}
	appconf = &Config{}
	appconf.LogLevel = conf.String("logs:log_level")
	if len(appconf.LogLevel) == 0 {
		appconf.LogLevel = "debug"
	}
	appconf.LogPath = conf.String("logs::log_path")
	if len(appconf.LogPath) == 0 {
		appconf.LogPath = "./logs"
	}
	appconf.Chan_size, err = conf.Int("collect::chan_size")
	if err != nil {
		appconf.Chan_size = 100
	}
	appconf.KafkaAddr = conf.String("kafka::server_addr")
	if len(appconf.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka address")
		return
	}
	appconf.Etcdaddr = conf.String("etcd::addr")
	if len(appconf.Etcdaddr) == 0 {
		err = fmt.Errorf("invalid etcd address")
		return
	}
	appconf.Etcdkey = conf.String("etcd::key")
	if len(appconf.Etcdkey) == 0 {
		err = fmt.Errorf("invalid etcd key")
		return
	}
	err = loadCollectConf(conf, appconf)
	if err != nil {
		fmt.Println("loadCollectConf is err :", err)
	}
	return
}
