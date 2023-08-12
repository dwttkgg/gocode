package uitls

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace:":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
func InitLogger(LogPath string, LogLevel string) (err error) {

	config := make(map[string]interface{})
	config["filename"] = LogPath
	config["level"] = convertLogLevel(LogLevel)

	bt, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed ,json marshal is error:", err)
	}
	logs.SetLogger(logs.AdapterFile, string(bt))

	//logs.Debug("this is a test ,my name is :", "std01")
	//logs.Trace("this is a Trace ,my name is :", "std02")
	//logs.Warn("this is a Warn ,my name is :", "std03")

	return
}
