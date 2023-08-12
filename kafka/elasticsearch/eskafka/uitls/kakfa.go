package uitls

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"strings"
	"sync"
)

// kafka consumer
// var wg sync.WaitGroup
type KafkaClient struct {
	client sarama.Consumer
	add    string
	topic  string
	wg     sync.WaitGroup
}

var kafkaClient *KafkaClient

func InitKafka(KafkaAddr string, topic string) (err error) {
	kafkaClient = &KafkaClient{}
	consumer, err := sarama.NewConsumer(strings.Split(KafkaAddr, ","), nil)
	if err != nil {
		logs.Error("fail to start consumer, err:%v\n", err)
		return
	}
	kafkaClient.client = consumer
	kafkaClient.add = KafkaAddr
	kafkaClient.topic = topic
	fmt.Println("-----------------", kafkaClient)
	return err
}
