package nums

import (
	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/tools/maths"
)

var PerfThr = 300
var PerfThrDiff = 100

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop(s []int) []int {
	var dup []int
	for _, v := range s {
		if !IsInIntSlice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap(s []int) []int {
	var dup []int
	tmpMap := make(map[int]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop32(s []int32) []int32 {
	var dup []int32
	for _, v := range s {
		if !IsInInt32Slice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap32(s []int32) []int32 {
	var dup []int32
	tmpMap := make(map[int32]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// int slice自去重,通过两重循环过滤重复元素
func uniqueIntByLoop64(s []int64) []int64 {
	var dup []int64
	for _, v := range s {
		if !IsInInt64Slice(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueIntByMap64(s []int64) []int64 {
	var dup []int64
	tmpMap := make(map[int64]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = 0
			dup = append(dup, v)
		}
	}
	return dup
}

// list自去重
func UniqueInt(s []int) []int {
	if len(s) < PerfThr {
		return uniqueIntByLoop(s)
	} else {
		return uniqueIntByMap(s)
	}
}

func UniqueInt32(s []int32) []int32 {
	if len(s) < PerfThr {
		return uniqueIntByLoop32(s)
	} else {
		return uniqueIntByMap32(s)
	}
}

func UniqueIntLen(s []int, l int) []int {
	ints := UniqueInt(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func UniqueInt32Len(s []int32, l int) []int32 {
	ints := UniqueInt32(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// s1对s2做差
func differIntByLoop(s1, s2 []int) []int {
	if len(s2) == 0 {
		return s1
	}
	var ints []int
	for _, i := range s1 {
		if !IsInIntSlice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// 通过map主键唯一的特性过滤重复元素
func differIntByMap(s1, s2 []int) []int {
	var dup []int
	tmpMap := make(map[int]byte, len(s2)) // 存放不重复主键
	for _, v := range s2 {
		tmpMap[v] = 0
	}
	for _, i := range s1 {
		if _, ok := tmpMap[i]; !ok {
			dup = append(dup, i)
		}
	}
	return dup
}

func differIntByLoop32(s1, s2 []int32) []int32 {
	if len(s2) == 0 {
		return s1
	}
	var ints []int32
	for _, i := range s1 {
		if !IsInInt32Slice(s2, i) {
			ints = append(ints, i)
		}
	}
	return ints
}

// 通过map主键唯一的特性过滤重复元素
func differIntByMap32(s1, s2 []int32) []int32 {
	var dup []int32
	tmpMap := make(map[int32]byte, len(s2)) // 存放不重复主键
	for _, v := range s2 {
		tmpMap[v] = 0
	}
	for _, i := range s1 {
		if _, ok := tmpMap[i]; !ok {
			dup = append(dup, i)
		}
	}
	return dup
}

func DifferInt(s1, s2 []int) []int {
	if len(s1) < PerfThrDiff {
		return differIntByLoop(s1, s2)
	} else {
		return differIntByMap(s1, s2)
	}
}

func DifferIntLen(s1, s2 []int, l int) []int {
	ints := DifferInt(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

func DifferInt32(s1, s2 []int32) []int32 {
	if len(s1) < PerfThrDiff {
		return differIntByLoop32(s1, s2)
	} else {
		return differIntByMap32(s1, s2)
	}
}

func DifferInt32Len(s1, s2 []int32, l int) []int32 {
	ints := DifferInt32(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}
