package tools

import (
	"math/rand"

	"github.com/xndm-recommend/go-utils/maths"
)

func NRandInt(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func IntToInterface(i []int) []interface{} {
	ifs := make([]interface{}, len(i))
	for ind, v := range i {
		ifs[ind] = v
	}
	return ifs
}

func InterfaceToInt(i []interface{}) []int {
	ints := make([]int, len(i))
	for ind, v := range i {
		ints[ind] = v.(int)
	}
	return ints
}

func IsInIntSlice(s []int, item int) bool {
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

// int slice自去重
func RmDuplicateInt(s []int) []int {
	dup := make([]int, 0, len(s))
	for _, v := range s {
		if !IsInIntSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

func RmDuplicateIntLen(s []int, i int) []int {
	if i <= 0 {
		return RmDuplicateInt(s)
	}
	return RmDuplicateInt(s)[:i]
}

// s1对s2做差
func DifferenceInt(s1, s2 []int) []int {
	dup := make([]int, 0, len(s1))
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

func DifferenceIntLen(s1, s2 []int, i int) []int {
	ints := DifferenceInt(s1, s2)
	if i < 0 {
		return ints
	}
	return ints[:maths.MinInt(len(ints), i)]
}

// int list union
func UnionIntList(s1, s2 []int) []int {
	return append(s1, s2...)
}

func UnionIntListDup(s1, s2 []int) []int {
	return RmDuplicateInt(append(s1, s2...))
}

func UnionIntListLen(s1, s2 []int, i int) []int {
	u := UnionIntList(s1, s2)
	if i < 0 {
		return u
	}
	return u[:maths.MinInt(len(u), i)]
}