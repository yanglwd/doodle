package network

import (
	"github.com/yanglwd/doodle/pkg/network/internal"
	"github.com/yanglwd/doodle/pkg/network/shared"
)

type Client interface {
	Init(shared.NetCientOptions) error
	Serve() error
	Cancel() error

	Write([]byte) error
	Read([]byte) error
}

type Server interface {
	Init(shared.NetServerOptions) error
	Serve() error
	Cancel() error

	Write([]byte) error
	Read([]byte) error
}

func NewClient(opts shared.NetCientOptions) Client {
	return internal.NewClient(opts)
}

func NewServer(opts shared.NetServerOptions) Server {
	return internal.NewServer(opts)
}
