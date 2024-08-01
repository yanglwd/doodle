package internal

import (
	"github.com/yanglwd/doodle/pkg/etcd/config"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdClient struct {
	conf   config.EtcdConfig
	client *clientv3.Client
}

var etcdInstanceInner *EtcdClient

func CreateAndInit(conf config.EtcdConfig) *EtcdClient {
	etcdInstanceInner = &EtcdClient{conf: conf}
	if err := etcdInstanceInner.Init(); err != nil {
		panic(err)
	}
	return etcdInstanceInner
}

func (c *EtcdClient) Init() error {
	etcd, err := clientv3.New(c.conf.Config)
	if err != nil {
		return err
	}
	c.client = etcd
	return nil
}

func (c *EtcdClient) Serve() error {
	return nil
}

func (c *EtcdClient) Stop() error {
	return nil
}
