package strs

// string list union
func MergeStrs(s1, s2 []string) []string {
	return append(s1, s2...)
}

func MergeStrsDup(s1, s2 []string) []string {
	return UniqueStrs(MergeStrs(s1, s2))
}

func MergeStrsLen(s1, s2 []string, l int) []string {
	return UniqueStrsLen(MergeStrs(s1, s2), l)
}

func MergeStrs32(s1, s2 []int32) []int32 {
	return append(s1, s2...)
}
