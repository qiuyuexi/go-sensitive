package server

import (
	"fmt"
	"go-sensitive/internal/pkg/ahocorasick"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"time"
	"context"
)

func Watch() {

	go func() {
		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{"127.0.0.1:2379"},
			DialTimeout: 5 * time.Second,
		})

		if err != nil {
			fmt.Println(err)
		}
		defer cli.Close()

		buidlTimestamp := strconv.Itoa(time.Now().Second())
		watchKey := "acdict_build_time"

		putTimeOutCtx, cancel := context.WithTimeout(context.Background(), time.Second)

		_, putErr := cli.Put(putTimeOutCtx, watchKey, buidlTimestamp)
		if putErr != nil {
			fmt.Println(putErr)
			return
		}
		cancel()

		for {
			rch := cli.Watch(context.Background(), watchKey)
			for wresp := range rch {
				for _, ev := range wresp.Events {
					fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					if string(ev.Kv.Value) != buidlTimestamp {
						ahocorasick.RebuildAhocorasickDict()
						fmt.Println("ac自动机更新")
						buidlTimestamp = strconv.Itoa(time.Now().Second())
						cli.Put(context.Background(), watchKey, buidlTimestamp)
					}
				}
			}
		}
	}()

}
