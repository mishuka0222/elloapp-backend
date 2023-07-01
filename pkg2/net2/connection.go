package net2

type Connection interface {
	GetConnID() uint64
	IsClosed() bool
	Close() error
	Codec() Codec
	Receive() (interface{}, error)
	Send(msg interface{}) error
}

type closeCallback interface {
	// func(Connection)
	OnConnectionClosed(Connection)
}
