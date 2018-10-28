package etcdconfig

import (
	"context"
	"fmt"
	"time"

	"github.com/kyawmyintthein/runtime-log-level-demo/config"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

var cli *clientv3.Client

func Init(generalConfig *config.GeneralConfig) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   generalConfig.ETCDHosts,
	})
	if err != nil {
		logrus.WithError(err).Error("Unable to connect etcd")
	}
	fmt.Println(cli.Endpoints())

	go func() {
		watchChan := cli.Watch(context.Background(), "us_dev", clientv3.WithPrefix())
		for true {
			select {
			case result := <-watchChan:
				for _, ev := range result.Events {
					fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					if string(ev.Kv.Key) == "us_dev/log_level" {
						generalConfig.LogLevel = (string(ev.Kv.Value))
						l, _ := logrus.ParseLevel(generalConfig.LogLevel)
						logrus.SetLevel(l)
					}
				}

			}
		}
	}()
}
