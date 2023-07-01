package net2

import (
	"sync"
)

type UdpConnection struct {
	addr      string
	codec     Codec
	sendChan  chan interface{}
	sendMutex sync.RWMutex
	isAlive   bool
}

func NewUdpConnection() *UdpConnection {
	return nil
}
