package redis

type RedisClienter interface {
	Init() error
	Serve() error
	Stop() error
}

func New(conf map[string]string) RedisClienter {
	return nil
}
