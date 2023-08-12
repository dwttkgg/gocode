package uitls

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type LogMessage struct {
	App     string `json:"app"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

var esClient *elastic.Client

func InitEs(addr string) (err error) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		// Handle error
		logs.Error("connect elasticsearch failed ,err:", err)
	}
	esClient = client

	//
	//fmt.Println("connect to es success")
	//p1 := Person{Name: "lmh", Age: 18, Married: false}
	//for i := 1; i < 10000; i++ {
	//	put1, err := client.Index().
	//		Index("user").
	//		Type("user").
	//		Id(fmt.Sprintf("%d", i)).
	//		BodyJson(p1).
	//		Do(context.Background())
	//	if err != nil {
	//		// Handle error
	//		panic(err)
	//	}
	//	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	//}
	return
}
func sendToES(topic string, data []byte) (err error) {
	msg := &LogMessage{}
	msg.Topic = topic
	msg.Message = string(data)
	_, err = esClient.Index().
		Index(topic).
		//Type(topic).
		//Id(fmt.Sprintf("%d", i)).
		BodyJson(msg).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	//logs.Debug("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	return
}
