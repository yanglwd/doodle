package etcd

import (
	"github.com/yanglwd/doodle/pkg/etcd/config"
	"github.com/yanglwd/doodle/pkg/etcd/internal"
)

type EtcdClienter interface {
	Init() error
	Serve() error
	Stop() error
}

func New(conf config.EtcdConfig) EtcdClienter {
	ec := internal.CreateAndInit(conf)
	return ec
}
