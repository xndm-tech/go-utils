package tools

import "github.com/xndm-recommend/go-utils/maths"

func StrToInterfaces(s []string) []interface{} {
	ifs := make([]interface{}, len(s))
	for ind, v := range s {
		ifs[ind] = v
	}
	return ifs
}

func IsInStrSlice(s []string, item string) bool {
	if 0 == len(s) {
		return false
	}
	for _, singleItem := range s {
		if item == singleItem {
			return true
		}
	}
	return false
}

func RmDuplicateStr(s []string) []string {
	dup := make([]string, len(s))
	for _, v := range s {
		if !IsInStrSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

func RmDuplicateStrLen(s []string, l int) []string {
	dup := make([]string, len(s))
	for _, v := range s {
		if !IsInStrSlice(dup, v) {
			dup = append(dup, v)
		}
		if len(dup) == l {
			break
		}
	}
	return dup
}

func DifferenceStr(s1, s2 []string) (x []string) {
	for _, i := range s1 {
		sign := true
		for _, v := range s2 {
			if i == v {
				sign = false
				break
			}
		}
		if true == sign {
			x = append(x, i)
		}
	}
	return x
}

func DifferenceStrLen(l1, l2 []string, outLen int) (x []string) {
	if outLen < 0 {
		return DifferenceStr(l1, l2)
	}
	for _, i := range l1 {
		sign := true
		for _, v := range l2 {
			if i == v {
				sign = false
				break
			}
		}
		if true == sign {
			x = append(x, i)
			if outLen == len(x) {
				return x
			}
		}
	}
	return x
}

func DifferenceAllowDup(list1, list2 []string, outLen int) (x []string) {
	x = DifferenceStrLen(list1, list2, outLen)
	if outLen < 0 {
		return x
	} else if len(x) < outLen {
		return append(x, list1[:maths.MinInt(outLen-len(x), len(list1))]...)
	} else {
		return x[:outLen]
	}
}

func UnionList(list1, list2 []string) (x []string) {
	return RmDuplicateStr(append(list1, list2...))
}

func UnionListLen(list1, list2 []string, outLen int) (x []string) {
	xTmp := UnionList(list1, list2)
	if outLen < 0 {
		return xTmp
	}
	return xTmp[:maths.MinInt(len(xTmp), outLen)]
}

func UnionListAllowDup(list1, list2 []string, outLen int) (u []string) {
	u = append(list1, list2...)
	if outLen < 0 {
		return u
	}
	return u[:maths.MinInt(len(u), outLen)]
}
