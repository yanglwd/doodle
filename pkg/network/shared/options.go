package shared

import "net"

type NetCallback func(net.Conn) error

type NetServerOptions struct {
	Network    string // TCP / KCP
	LisentAddr string

	OnAccept        NetCallback
	OnNewConnection NetCallback
	OnClose         NetCallback
}

type NetCientOptions struct {
	Network    string // TCP / KCP
	LocalAddr  string // 从本地哪个 IP 发起连接
	RemoteAddr string // 远端 IP

	OnConnect NetCallback
	OnClose   NetCallback
}
