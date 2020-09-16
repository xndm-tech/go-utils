package nums

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/tools/maths"
)

const length = 500

func TestRmDuplicateIntLen(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 4, 4, 6}
	fmt.Println(UniqueIntLen(a, 9))
}

func BenchmarkUniqueIntByLoop(b *testing.B) {
	b.StopTimer()
	a := maths.GenNRandInt(length, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		uniqueIntByLoop(a)
	}
}

func BenchmarkUniqueIntByMap(b *testing.B) {
	b.StopTimer()
	a := maths.GenNRandInt(length, 100)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		uniqueIntByMap(a)
	}
}

func BenchmarkUniqueInt(b *testing.B) {
	b.ResetTimer()
	a := maths.GenNRandInt(length, 100)
	for i := 0; i < b.N; i++ {
		UniqueInt(a)
	}
}
