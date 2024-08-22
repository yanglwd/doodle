package internal

import (
	"github.com/yanglwd/doodle/pkg/network/shared"
)

type InternalClient struct {
}

func NewClient(opts shared.NetCientOptions) *InternalClient {
	c := &InternalClient{}
	if err := c.Init(opts); err != nil {
		panic(err)
	}
	return c
}

func (c *InternalClient) Init(opt shared.NetCientOptions) error {
	return nil
}

func (c *InternalClient) Serve() error {
	return nil
}

func (c *InternalClient) Cancel() error {
	return nil
}

func (c *InternalClient) Write([]byte) error {
	return nil
}

func (c *InternalClient) Read([]byte) error {
	return nil
}
