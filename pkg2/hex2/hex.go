package hex2

import (
	"encoding/hex"
)

func HexDumpSize(size int, buf []byte) string {
	if size > len(buf) {
		size = len(buf)
	}
	return hex.EncodeToString(buf[:size])
}

func HexDump(buf []byte) string {
	return HexDumpSize(128, buf)
}
