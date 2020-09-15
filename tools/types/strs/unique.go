package strs

import (
	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/tools/maths"
)

var PerfThr = 300

// int slice自去重,通过两重循环过滤重复元素
func uniqueStrsByLoop(s []string) []string {
	dup := make([]string, consts.ZERO, len(s))
	for _, v := range s {
		if !IsContainStr(dup, v) {
			dup = append(dup, v)
		}
	}
	return dup
}

// 通过map主键唯一的特性过滤重复元素
func uniqueStrsByMap(s []string) []string {
	dup := make([]string, consts.ZERO, len(s))
	tmpMap := make(map[string]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = consts.ZERO
			dup = append(dup, v)
		}
	}
	return dup
}

// 自去重
func UniqueStrs(s []string) []string {
	if len(s) < PerfThr {
		return uniqueStrsByLoop(s)
	} else {
		return uniqueStrsByMap(s)
	}
}

func UniqueStrsLen(s []string, l int) []string {
	ints := UniqueStrs(s)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(ints))
	return ints[:size:size]
}

// string间做去重
func DifferStrs(s1, s2 []string) []string {
	if len(s2) == 0 {
		return s1
	}
	strs := make([]string, 0, len(s1))
	for _, i := range s1 {
		if !IsContainStr(s2, i) {
			strs = append(strs, i)
		}
	}
	return strs
}

func DifferStrsLen(s1, s2 []string, l int) []string {
	strs := DifferStrs(s1, s2)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(strs))
	return strs[:size:size]
}

// 长度大于30使用
func bitMapFilterInt32(s, filter []int32, bitArr gobal.BitMap, l int) []int32 {
	for _, i := range filter {
		index := i >> 3
		pos := i & 0x07
		bitArr.Bits[index] ^= 1 << pos
	}
	output := make([]int32, 0)

	for _, i := range s {
		index := i >> 3
		pos := i & 0x07
		if bitArr.Bits[index]&(1<<pos) != 0 {
			output = append(output, i)
		}
	}
	for _, i := range filter {
		index := i >> 3
		pos := i & 0x07
		bitArr.Bits[index] ^= 1 << pos
	}
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}

// 自去重长度大于35的时候使用
func bitMapDifferenceSelfInt32(s []int32, bitList []byte, l int) []int32 {
	output := make([]int32, 0)
	for _, i := range s {
		index := i >> 3
		pos := i & 0x07
		if bitList[index]&(1<<pos) == 0 {
			bitList[index] |= 1 << pos
			output = append(output, i)
		}
	}

	for _, i := range s {
		index := i >> 3
		bitList[index] = 0
	}

	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}
