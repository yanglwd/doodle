package net

type Server interface {
	Init(string) // panic(err)
	Serve()      // panic(err)
	Cancel() error
}
