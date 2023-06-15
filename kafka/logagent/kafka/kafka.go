package kafka

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka connect failed, err:", err)
		return
	}
	//defer client.Close()
	logs.Debug("init kafka success")
	return
}
func SendToKafka(data, topic string) (err error) {

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	// 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Printf("send msg failed, err:%v,data:%v,topic:%v\n", err, data, topic)
		return
	}
	logs.Debug("partition:%v offset:%v topic:%v", partition, offset, topic)
	return
}
