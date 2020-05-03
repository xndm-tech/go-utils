package slices_test

import (
	"testing"

	"github.com/xndm-recommend/go-utils/common/slices"
)

func TestRemoveRepByLoop(t *testing.T) {
	a := slices.GenNRandInt(50, 10)
	t.Log(a)
	t.Log(slices.RemoveRepByLoop(a))
}

func TestRemoveRepByMap(t *testing.T) {
	a := slices.GenNRandInt(50, 10)
	t.Log(a)
	t.Log(slices.RemoveRepByMap(a))
}

func BenchmarkRemoveRepByMap(b *testing.B) {
	a := slices.GenNRandInt(300, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		slices.RemoveRepByMap(a)
	}
}

func BenchmarkRemoveRepByLoop(b *testing.B) {
	a := slices.GenNRandInt(300, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		slices.RemoveRepByLoop(a)
	}
}
