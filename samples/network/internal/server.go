package internal

type TinyServer struct {
	ListenAddr string
}

func (s *TinyServer) Init(addr string) error {
	return nil
}

func (s *TinyServer) Serve() error {
	return nil
}

func (s *TinyServer) Stop() error {
	return nil
}
