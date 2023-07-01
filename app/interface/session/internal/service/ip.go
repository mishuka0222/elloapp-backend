package service

import (
	"github.com/zeromicro/go-zero/core/netx"
	"os"
)

const (
	envPodIp = "POD_IP"
)

var (
	internalIp string
)

func init() {
	internalIp = os.Getenv(envPodIp)
	if len(internalIp) == 0 {
		internalIp = netx.InternalIp()
	}
}
