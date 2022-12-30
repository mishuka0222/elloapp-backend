package ip

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/netx"
)

const (
	allEths  = "0.0.0.0"
	envPodIp = "POD_IP"
)

func FigureOutListenOn(listenOn string) string {
	fields := strings.Split(listenOn, ":")
	if len(fields) == 0 {
		return listenOn
	}

	host := fields[0]
	if len(host) > 0 && host != allEths {
		return listenOn
	}

	ip := os.Getenv(envPodIp)
	if len(ip) == 0 {
		ip = netx.InternalIp()
	}
	if len(ip) == 0 {
		return listenOn
	}

	return strings.Join(append([]string{ip}, fields[1:]...), ":")
}

func IpToInt(ipStr string) uint32 {
	var (
		numList            = strings.Split(ipStr, ".")
		sz                 = len(numList)
		ip1, ip2, ip3, ip4 int
	)

	if sz > 0 {
		ip1, _ = strconv.Atoi(numList[0])
	}
	if sz > 1 {
		ip2, _ = strconv.Atoi(numList[1])
	}
	if sz > 2 {
		ip3, _ = strconv.Atoi(numList[2])
	}
	if sz > 3 {
		ip4, _ = strconv.Atoi(numList[3])
	}

	return uint32(ip1&0xffff<<24 | ip2&0xffff<<16 | ip3&0xffff<<8 | ip4&0xffff)
}

func IntToIP(ip uint32) string {
	ip1 := ip >> 24 & 0xffff
	ip2 := ip << 8 >> 24 & 0xffff
	ip3 := ip << 16 >> 24 & 0xffff
	ip4 := ip << 24 >> 24 & 0xffff

	return fmt.Sprintf("%d.%d.%d.%d", ip1, ip2, ip3, ip4)
}
