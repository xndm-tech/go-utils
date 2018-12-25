package tools

import "github.com/xndm-recommend/go-utils/maths"

func StrToInterface(s []string) []interface{} {
	ifs := make([]interface{}, len(s))
	for ind, v := range s {
		ifs[ind] = v
	}
	return ifs
}

func InterfaceToStr(s []interface{}) []string {
	strings := make([]string, len(s))
	for ind, v := range s {
		strings[ind] = v.(string)
	}
	return strings
}

func IsInStrSlice(s []string, item string) bool {
	if 0 == len(s) {
		return false
	}
	for _, sItem := range s {
		if item == sItem {
			return true
		}
	}
	return false
}

// string slice自去重
func RmDuplicateStr(s []string) []string {
	dup := make([]string, 0, len(s))
	for _, v := range s {
		if !IsInStrSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

func RmDuplicateStrLen(s []string, i int) []string {
	if i <= 0 {
		return RmDuplicateStr(s)
	}
	return RmDuplicateStr(s)[:i]
}

// s1对s2做差
func DifferenceStr(s1, s2 []string) []string {
	dup := make([]string, 0, len(s1))
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

func DifferenceStrLen(s1, s2 []string, i int) []string {
	strings := DifferenceStr(s1, s2)
	if i < 0 {
		return strings
	}
	return strings[:maths.MinInt(len(strings), i)]
}

// string list union
func UnionStrList(s1, s2 []string) []string {
	return append(s1, s2...)
}

func UnionStrListDup(s1, s2 []string) []string {
	return RmDuplicateStr(append(s1, s2...))
}

func UnionStrListLen(s1, s2 []string, i int) []string {
	u := UnionStrList(s1, s2)
	if i < 0 {
		return u
	}
	return u[:maths.MinInt(len(u), i)]
}
