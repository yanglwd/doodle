package internal

import (
	"net"

	"github.com/xtaci/kcp-go/v5"
	"github.com/yanglwd/doodle/pkg/network/shared"
)

type InternalServer struct {
	NetOpts shared.NetServerOptions

	ln net.Listener
}

func NewServer(opts shared.NetServerOptions) *InternalServer {
	s := &InternalServer{}
	if err := s.Init(opts); err != nil {
		panic(err)
	}
	return s
}

func (s *InternalServer) Init(opts shared.NetServerOptions) error {
	s.NetOpts = opts
	switch s.NetOpts.Network {
	case "kcp":
		return s.initKCP()
	case "tcp":
		return s.initTCP()
	}
	return shared.ErrInvalidProtocol
}

func (s *InternalServer) initTCP() error {
	ln, err := net.Listen(s.NetOpts.Network, s.NetOpts.LisentAddr)
	if err != nil {
		return err
	}

	s.ln = ln
	return nil
}

func (s *InternalServer) initKCP() error {
	ln, err := kcp.Listen(s.NetOpts.LisentAddr)
	if err != nil {
		return err
	}

	s.ln = ln
	return nil
}

func (s *InternalServer) Serve() error {
	switch s.NetOpts.Network {
	case "kcp":
		return s.serveKCP()
	default:
		return s.serveTCP()
	}
}

func (s *InternalServer) serveTCP() error {
	return nil
}

func (s *InternalServer) serveKCP() error {
	// accpetor, ok := s.ln.(*kcp.Listener)
	// if !ok {
	// 	return shared.ErrInvalidListener
	// }
	// for {
	// 	c, err := accpetor.AcceptKCP()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	return nil
}

func (s *InternalServer) Cancel() error {
	s.ln.Close()
	return nil
}

func (s *InternalServer) Write([]byte) error {
	return nil
}

func (s *InternalServer) Read([]byte) error {
	return nil
}
