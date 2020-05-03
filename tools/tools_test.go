package tools_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/common/slices"
	"github.com/xndm-recommend/go-utils/tools"
)

func TestRmDuplicateIntLen(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 4, 4, 6}
	fmt.Println(tools.RmDuplicateIntLen(a, 9))
}

func TestCutSlice(t *testing.T) {
	a := slices.GenNRandInt(10, 10)
	t.Log(a)
	t.Log(a[:0:0])
}

func TestCutSlice2(t *testing.T) {
	a := slices.GenNRandInt(10, 10)
	t.Log(a)
	t.Log(a[10:10])
}

func TestMixListIntLenV2(t *testing.T) {
	a := slices.GenNRandInt(30, 20)
	b := slices.GenNRandInt(30, 20)
	t.Log(a)
	t.Log(b)
	t.Log(tools.MixListIntLenV2(10, a, b))
}

func TestMixListInt32LenV2(t *testing.T) {
	a := slices.GenNRandInt32(30, 20)
	b := slices.GenNRandInt32(30, 20)
	t.Log(a)
	t.Log(b)
	t.Log(tools.MixListInt32LenV2(10, a, b))
}

func BenchmarkRmDuplicateIntLen(b *testing.B) {
	a := slices.GenNRandInt(300, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		tools.RmDuplicateIntLen(a, 10)
	}
}

func BenchmarkRemoveRepByLoop(b *testing.B) {
	a := slices.GenNRandInt(300, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		slices.RemoveRepByLoop(a)
	}
}
