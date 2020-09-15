package nums

import (
	"math/rand"
	"time"
)

// Shuffle
func ShuffleIntList(s []int) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func ShuffleInt32List(s []int32) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func ReverseIntList(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseInt32List(s []int32) []int32 {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
