package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"logagent/tailf"
	"logagent/utils"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
	keys   []string
}

var etcdCli *EtcdClient

func InitEtcd(addr string, key string) (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		// handle error!
		logs.Error("connect to etcd failed, err:%v", err)
		return
	}
	etcdCli = &EtcdClient{client: cli}
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	var collectConf []tailf.CollectConf
	//fmt.Println("---------------------------", utils.Ip)
	for _, ipaddr := range utils.Ip {

		etcdkey := fmt.Sprintf("%s%s", key, ipaddr)
		//fmt.Println("etcdkey is :", etcdkey)
		etcdCli.keys = append(etcdCli.keys, etcdkey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdkey)
		if err != nil {
			continue
		}
		cancel()
		logs.Debug("resp from etcd %v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdkey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("json resp.kvs.value unmarshal failed ,err:", err)
					continue
				}
				logs.Debug("log config is :", collectConf)
			}
		}
	}
	initEtcdWatcher()
	return
}
func initEtcdWatcher() {
	for _, v := range etcdCli.keys {
		go watchKey(v)
	}
}
func watchKey(key string) {
	for {
		rch := etcdCli.client.Watch(context.Background(), key)
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Debug("the key deleted :", key)
					continue
				}
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("put value is failed ,err:", err)
						getConfSucc = false
						continue
					}
					logs.Debug("the key ")
				}

				fmt.Printf("%s %q :%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc == true {
				tailf.UpdateConf(collectConf)
			}
		}
	}
}
