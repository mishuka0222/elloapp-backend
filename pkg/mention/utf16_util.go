package mention

import (
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

const (
	replacementChar = '\uFFFD'     // Unicode replacement character
	maxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
)

const (
	// 0xd800-0xdc00 encodes the high 10 bits of a pair.
	// 0xdc00-0xe000 encodes the low 10 bits of a pair.
	// the value is those 20 bits plus 0x10000.
	surr1 = 0xd800
	surr2 = 0xdc00
	surr3 = 0xe000

	surrSelf = 0x10000
)

// Encode returns the UTF-16 encoding of the Unicode code point sequence s.
func EncodeStringToUTF16Index(s string) []int {
	var (
		a = make([]int, len(s)+1)
		n = 0
		i = 0
	)
	for _, v := range s {
		rLen := utf8.RuneLen(v)
		switch {
		case 0 <= v && v < surr1, surr3 <= v && v < surrSelf:
			for i2 := 0; i2 < rLen; i2++ {
				a[i] = n
				i++
			}
			n++
		case surrSelf <= v && v <= maxRune:
			for i2 := 0; i2 < rLen; i2++ {
				a[i] = n
				i++
			}
			n += 2
		default:
			for i2 := 0; i2 < rLen; i2++ {
				a[i] = n
				i++
			}
			n++
		}
	}
	a[i] = n

	return a
}

func EncodeStringToUTF16(s string) []uint16 {
	n := 0
	for _, v := range s {
		n++
		if v >= 0x10000 {
			n++
		}
	}

	a := make([]uint16, n)
	n = 0
	for _, v := range s {
		switch {
		case 0 <= v && v < surr1, surr3 <= v && v < surrSelf:
			// normal rune
			a[n] = uint16(v)
			n++
		case surrSelf <= v && v <= maxRune:
			// needs surrogate sequence
			r1, r2 := utf16.EncodeRune(v)
			a[n] = uint16(r1)
			a[n+1] = uint16(r2)
			n += 2
		}
	}
	return a[:n]
}

func DecodeUTF16ToString(s []uint16) string {
	n := 0
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			n += utf8.RuneLen(rune(r))
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			n += utf8.RuneLen(utf16.DecodeRune(rune(r), rune(s[i+1])))
			i++
		default:
			// invalid surrogate sequence
			n += utf8.RuneLen(replacementChar)
		}
	}
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(s); i++ {
		switch r := s[i]; {
		case r < surr1, surr3 <= r:
			// normal rune
			b.WriteRune(rune(r))
		case surr1 <= r && r < surr2 && i+1 < len(s) &&
			surr2 <= s[i+1] && s[i+1] < surr3:
			// valid surrogate sequence
			b.WriteRune(utf16.DecodeRune(rune(r), rune(s[i+1])))
			i++
		default:
			// invalid surrogate sequence
			b.WriteRune(replacementChar)
		}
	}
	return b.String()
}
