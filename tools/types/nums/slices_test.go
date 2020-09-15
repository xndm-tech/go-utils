package nums_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/tools/types/nums"
)

func TestRmDuplicateIntLen(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 4, 4, 6}
	fmt.Println(nums.UniqueIntLen(a, 9))
}

//func TestRemoveRepByLoop(t *testing.T) {
//	a := maths.GenNRandInt(50, 10)
//	t.Log(a)
//	t.Log(nums.uniqueIntByLoop(a))
//}
//
//func TestRemoveRepByMap(t *testing.T) {
//	a := maths.GenNRandInt(50, 10)
//	t.Log(a)
//	t.Log(nums.uniqueIntByMap(a))
//}
//
//func BenchmarkRemoveRepByMap(b *testing.B) {
//	b.ResetTimer()
//	a := maths.GenNRandInt(3000, 100)
//	//b.Log(a)
//	for i := 0; i < b.N; i++ {
//		nums.uniqueIntByMap(a)
//	}
//}

//
//func BenchmarkRemoveRepByRedis(b *testing.B) {
//	b.ResetTimer()
//	a := maths.GenNRandInt(30000, 100)
//	//b.Log(a)
//	for i := 0; i < b.N; i++ {
//		nums.RemoveRepBySet(a)
//	}
//}
