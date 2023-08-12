package uitls

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

type Logconfig struct {
	KafkaAddr string
	EsAddr    string
	LogPath   string
	LogLevel  string
	Topic     string
}

var logConfig *Logconfig

func InitConfig(ConfType string, filename string) (logConfig *Logconfig, err error) {
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
	logConfig.EsAddr = conf.String("es::addr")
	if len(logConfig.EsAddr) == 0 {
		err = fmt.Errorf("invalid elasticsearch address")
		return
	}
	logConfig.Topic = conf.String("kafka::Topic")
	if len(logConfig.Topic) == 0 {
		err = fmt.Errorf("invalid Topic address")
		return
	}

	return
}
