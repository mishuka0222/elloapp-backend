package hashx

import (
	"fmt"
	"testing"
)

func TestCombineInt64Hash(t *testing.T) {
	var acc int64 = 0
	for i := 1; i < 2; i++ {
		acc = CombineInt64Hash(acc, int64(i))
	}
	fmt.Println(acc)

	var acc2 int64 = 0
	for i := 1; i < 2; i++ {
		acc2 = CombineInt64Hash2(acc2, int64(i))
	}
	fmt.Println(acc2)
}
