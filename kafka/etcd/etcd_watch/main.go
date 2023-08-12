package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.30.11:2379", "192.168.30.12:2379", "192.168.30.13:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		// handle error!W
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	rch := cli.Watch(context.Background(), "/api/v1")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
