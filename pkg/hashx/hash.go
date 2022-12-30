package hashx

// CombineInt64Hash2
// ?????
func CombineInt64Hash2(acc, id int64) int64 {
	acc ^= acc >> 21
	acc ^= acc << 35
	acc ^= acc >> 4
	return acc + id
}

// CombineInt64Hash
// ?????
func CombineInt64Hash(acc, id int64) int64 {
	acc ^= id >> 21
	acc ^= id << 35
	acc ^= id >> 4
	return acc + id
}

func HashInt64(acc int64) int32 {
	return int32(acc&0xffffffff ^ acc>>32)
}
