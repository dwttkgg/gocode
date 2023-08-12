package uitls

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

func Run() (err error) {
	fmt.Println("-----------------------", kafkaClient.topic)
	partitionList, err := kafkaClient.client.Partitions(kafkaClient.topic) // 根据topic取到所有的分区
	if err != nil {
		logs.Error("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := kafkaClient.client.ConsumePartition(kafkaClient.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			logs.Error("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(pc sarama.PartitionConsumer) {
			kafkaClient.wg.Add(1)
			for msg := range pc.Messages() {
				logs.Debug("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				err = sendToES(kafkaClient.topic, msg.Value)
				if err != nil {
					logs.Warn("send to elasticsearch failed ,err:%v", err)
				}
			}
			kafkaClient.wg.Done()
		}(pc)
	}
	kafkaClient.wg.Wait()
	return
}
