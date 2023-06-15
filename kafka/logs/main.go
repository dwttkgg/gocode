package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	config["level"] = logs.LevelInfo

	bt, err := json.Marshal(config)
	if err != nil {
		fmt.Println("json marshal is error:", err)
	}
	logs.SetLogger(logs.AdapterFile, string(bt))

	logs.Debug("this is a test ,my name is :", "std01")
	logs.Trace("this is a Trace ,my name is :", "std02")
	logs.Warn("this is a Warn ,my name is :", "std03")
}
