package nums

// int list union
func MergeInts(s1, s2 []int) []int {
	return append(s1, s2...)
}

func MergeIntsDup(s1, s2 []int) []int {
	return UniqueInt(MergeInts(s1, s2))
}

func MergeIntsLen(s1, s2 []int, l int) []int {
	return UniqueIntLen(MergeInts(s1, s2), l)
}

func MergeInts32(s1, s2 []int32) []int32 {
	return append(s1, s2...)
}

func MergeInts32Dup(s1, s2 []int32) []int32 {
	return UniqueInt32(MergeInts32(s1, s2))
}

func MergeInts32Len(s1, s2 []int32, l int) []int32 {
	return UniqueInt32Len(MergeInts32(s1, s2), l)
}
