package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web1_log"
	msg.Value = sarama.StringEncoder("this is a test log 11")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"192.168.30.101:9092"}, config)
	fmt.Println(msg)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	//time.Sleep(time.Second)
	fmt.Printf("partition:%v offset:%v\n", partition, offset)
}
