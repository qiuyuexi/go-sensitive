package server

import (
	"context"
	"fmt"
	"go-sensitive/internal/pkg/ahocorasick"
	"go-sensitive/internal/pkg/config"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"time"
)

var AcDictIns *acDict

type acDict struct {
	isUpdate int
	ac       [2]*ahocorasick.Ahocorasick
}

func Start() {
	config.SetWorkPath()
	buildAhocorasickDict()
	watch()
	printOut()
}

/**
生成ac自动机字典
 */
func buildAhocorasickDict() {
	sensitiveWordConfig := config.LoadSensitiveWords()
	AcDictIns = new(acDict)
	AcDictIns.isUpdate = 0
	AcDictIns.ac[0] = ahocorasick.GetAhocorasick()
	for k, v := range sensitiveWordConfig {
		AcDictIns.ac[0].Tree[k] = ahocorasick.GetTire(v)
	}
	AcDictIns.ac[1] = AcDictIns.ac[0]
}

func rebuildAhocorasickDict() {
	AcDictIns.isUpdate = 1
	sensitiveWordConfig := config.LoadSensitiveWords()
	newAcDictIns := new(acDict)
	newAcDictIns.isUpdate = 0
	newAcDictIns.ac[0] = ahocorasick.GetAhocorasick()
	for k, v := range sensitiveWordConfig {
		newAcDictIns.ac[0].Tree[k] = ahocorasick.GetTire(v)
	}
	newAcDictIns.ac[1] = newAcDictIns.ac[0]
	AcDictIns.ac[0] = newAcDictIns.ac[0]
	AcDictIns.isUpdate = 0
	AcDictIns.ac[1] = newAcDictIns.ac[1]
}

/**
获取ac自动机
 */
func GetAcDictIns() *ahocorasick.Ahocorasick {
	var acIns *ahocorasick.Ahocorasick
	if AcDictIns.isUpdate == 0 {
		acIns = AcDictIns.ac[0]
	} else {
		acIns = AcDictIns.ac[1]
	}
	return acIns
}

func watch() {

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
						rebuildAhocorasickDict()
						fmt.Println("ac自动机更新")
						buidlTimestamp = strconv.Itoa(time.Now().Second())
						cli.Put(context.Background(), watchKey, buidlTimestamp)
					}
				}
			}
		}
	}()

}

func printOut() {
	fmt.Println("启动...")
}
