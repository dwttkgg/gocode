package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func etcdExample() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.30.101:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	// etcd设置一个值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := cli.Put(ctx, "/api/v1", "node")
	cancel()
	if err != nil {
		fmt.Println("etcd put value failed ,err:", err)
		return
	}
	fmt.Println("etcd response is：", res)

	// etcd get a value
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gres, err := cli.Get(ctx, "/api/v1")
	cancel()
	if err != nil {
		fmt.Println("etcd get value failed ,err:", err)
		return
	}
	for _, v := range gres.Kvs {
		fmt.Printf("%s value is %s", v.Key, v.Value)
	}
}

const (
	Etcdkey = "/api/v1/namespace/pod/192.168.30.1"
)

type LogConf struct {
	Path    string `json:"path""`
	Topic   string `json:"topic"`
	SendQps int
}

func etcdSetConf() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.30.101:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

	var logConf []LogConf
	logConf = append(logConf, LogConf{
		Path:  "./logs/logagent.log",
		Topic: "nginx_log",
	})
	logConf = append(logConf, LogConf{
		Path:  "/home/admin/elasticsearch/logs/my-application.log",
		Topic: "es_log",
	})
	data, err := json.Marshal(logConf)
	if err != nil {
		logs.Error("json marshal error ", err)
	}
	// etcd设置一个值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := cli.Put(ctx, Etcdkey, string(data))
	cancel()
	if err != nil {
		fmt.Println("etcd put value failed ,err:", err)
		return
	}
	fmt.Println("etcd response is：", res)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	gres, err := cli.Get(ctx, Etcdkey)
	cancel()
	if err != nil {
		fmt.Println("etcd get value failed ,err:", err)
		return
	}
	for _, v := range gres.Kvs {
		fmt.Printf("%s value is %s", v.Key, v.Value)
	}
}
func main() {
	//etcdExample()
	etcdSetConf()
}
