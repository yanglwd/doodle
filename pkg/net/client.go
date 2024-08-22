package net

type Cient interface {
	Init(string) // panic(err)
	Serve()      // panic(err)
	Cancel() error
}
