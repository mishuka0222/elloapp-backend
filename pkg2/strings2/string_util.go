package strings2

import (
	"strconv"
)

func JoinInt32List(s []int32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, int64(s[i]), 10)
		// buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint32List(s []uint32, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, uint64(s[i]), 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinInt64List(s []int64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendInt(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

func JoinUint64List(s []uint64, sep string) string {
	l := len(s)
	if l == 0 {
		return ""
	}

	buf := make([]byte, 0, l*2+len(sep)*l+len(sep)*(l-1))
	for i := 0; i < l; i++ {
		buf = strconv.AppendUint(buf, s[i], 10)
		buf = append(buf, sep...)
		if i != l-1 {
			buf = append(buf, sep...)
		}
	}
	return string(buf)
}

// IsAlNumString returns true if an alpha numeric string consists of characters a-zA-Z0-9
//
// https://github.com/nebulaim/telegramd/issues/99
//
// there are some issue in username validation
// 1 - issue : username can be only numeric [must check first character for alphabet]
// 2 - issue :underscore not allowed [must edit IsAlNumString and add underscore support to it.
//
//	by the way underscore can't repeat without any character between them. for example s_a_b_a. ]
func IsAlNumString(s string) bool {
	c := 0
	prevtmp := ' '
	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			c++
			break
		case 'a' <= r && r <= 'z':
			c++
			break
		case 'A' <= r && r <= 'Z':
			c++
			break
		case prevtmp != '_' && '_' == r:
			c++
			break
		}
		prevtmp = r
	}
	return len(s) == c
}
