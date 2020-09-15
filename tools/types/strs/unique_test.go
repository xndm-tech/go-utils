package strs_test

import (
	"fmt"
	"testing"

	strs2 "github.com/xndm-recommend/go-utils/tools/types/nums"

	"github.com/xndm-recommend/go-utils/tools/types/strs"
)

func TestRmDuplicateIntLen(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 4, 4, 6}
	fmt.Println(strs2.RmDuplicateIntLen(a, 9))
}

func TestRightStr(t *testing.T) {
	t.Log(strs.RightStr("abcd", 20))
}

func TestJoinStrSlice(t *testing.T) {
	t.Log(strs.JoinStrSlice("", []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"}))
	t.Log(strs.JoinStrs("", "1", "2", "3", "1", "2", "3", "1", "2", "3"))
}

func BenchmarkJoinStrSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strs.JoinStrSlice("|", []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"})
	}
}
