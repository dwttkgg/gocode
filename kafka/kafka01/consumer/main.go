package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

// kafka consumer
var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.30.101:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("nginx_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
			defer wg.Done()
		}(pc)
	}
	wg.Wait()
	time.Sleep(time.Second * 5)
}
