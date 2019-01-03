package tools

// s1对s2做差
func DiffInterface(s1, s2 []interface{}) []interface{} {
	dup := make([]interface{}, 0, len(s1))
	for _, i := range s1 {
		sign := true
		for _, v := range s2 {
			if i == v {
				sign = false
				break
			}
		}
		if true == sign {
			dup = append(dup, i)
		}
	}
	return dup
}
