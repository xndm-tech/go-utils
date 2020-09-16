package strs

import (
	"testing"
)

func TestRightStr(t *testing.T) {
	t.Log(RightStr("abcd", 20))
}

func TestJoinStrSlice(t *testing.T) {
	t.Log(JoinStrSlice("", []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"}))
	t.Log(JoinStrs("", "1", "2", "3", "1", "2", "3", "1", "2", "3"))
}

func BenchmarkJoinStrSlice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinStrSlice("|", []string{"1", "2", "3", "1", "2", "3", "1", "2", "3"})
	}
}
