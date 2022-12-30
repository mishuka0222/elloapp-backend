package net2

type ClientConfig struct {
	Name      string
	ProtoName string
	AddrList  []string
	EtcdAddrs []string
	Balancer  string
}

type ServerConfig struct {
	Name      string
	ProtoName string
	Addr      string
}
