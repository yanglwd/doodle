package config

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig struct {
	clientv3.Config
}

func (cfg *EtcdConfig) ParserFromMap(kv map[string]string) {

}
