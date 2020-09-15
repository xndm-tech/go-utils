package nums

// int list union
func UnionInts(s1, s2 []int) []int {
	return append(s1, s2...)
}

func UnionIntsDup(s1, s2 []int) []int {
	return RmDuplicateInt(UnionInts(s1, s2))
}

func UnionIntsLen(s1, s2 []int, l int) []int {
	return RmDuplicateIntLen(UnionInts(s1, s2), l)
}

func UnionInts32(s1, s2 []int32) []int32 {
	return append(s1, s2...)
}

func UnionInts32Dup(s1, s2 []int32) []int32 {
	return RmDuplicateInt32(UnionInts32(s1, s2))
}

func UnionInts32Len(s1, s2 []int32, l int) []int32 {
	return RmDuplicateInt32Len(UnionInts32(s1, s2), l)
}
