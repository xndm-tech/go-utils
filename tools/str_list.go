package tools

import (
	"errors"
	"math/rand"

	"github.com/xndm-recommend/go-utils/maths"
)

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

// string list
func GetStrListNoLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, errors.New("Input paras error")
	}
	return s[maths.MinInt(size*(num-1), len(s)):maths.MinInt(num*size, len(s))], nil
}

func GetStrListLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 || size >= len(s) {
		return []string{}, errors.New("Input parameter error!!!")
	}
	start := (size * (num - 1)) % len(s)
	end := (num * size) % len(s)
	if start < end {
		return s[start:end:end], nil
	} else {
		out := make([]string, 0, size)
		out = append(out, s[start:len(s):len(s)]...)
		out = append(out, s[:end:end]...)
		return out, nil
	}
}

// split list
func SplitStrList(s []string) (s1, s2 []string) {
	for i := 0; i < len(s); i++ {
		if 0 == i%2 {
			s1 = append(s1, s[i])
		} else {
			s2 = append(s2, s[i])
		}
	}
	return
}

func ShuffleStrList(s []string) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
