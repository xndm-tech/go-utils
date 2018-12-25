package tool_utils

import "github.com/zhanglanhui/go-utils/utils/math_utils"

func IsInSlice(list []string, item string) bool {
	if 0 == len(list) {
		return false
	}
	for _, singleItem := range list {
		if item == singleItem {
			return true
		}
	}
	return false
}

func RmDuplicate(list []string) (out []string) {
	for _, v := range list {
		if !IsInSlice(out, v) {
			out = append(out, v)
		}
	}
	return out
}

func RmDuplicateLen(list []string, outLen int) (out []string) {
	for _, v := range list {
		if !IsInSlice(out, v) {
			out = append(out, v)
		}
		if outLen == len(out) {
			break
		}
	}
	return out
}

func Difference(l1, l2 []string) (x []string) {
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
		}
	}
	return x
}

func DifferenceLen(l1, l2 []string, outLen int) (x []string) {
	if outLen < 0 {
		return Difference(l1, l2)
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
	x = DifferenceLen(list1, list2, outLen)
	if outLen < 0 {
		return x
	} else if len(x) < outLen {
		return append(x, list1[:math_utils.MinInt(outLen-len(x), len(list1))]...)
	} else {
		return x[:outLen]
	}
}

func UnionList(list1, list2 []string) (x []string) {
	return RmDuplicate(append(list1, list2...))
}

func UnionListLen(list1, list2 []string, outLen int) (x []string) {
	xTmp := UnionList(list1, list2)
	if outLen < 0 {
		return xTmp
	}
	return xTmp[:math_utils.MinInt(len(xTmp), outLen)]
}

func UnionListAllowDup(list1, list2 []string, outLen int) (xTmp []string) {
	xTmp = append(list1, list2...)
	if outLen < 0 {
		return xTmp
	}
	return xTmp[:math_utils.MinInt(len(xTmp), outLen)]
}
