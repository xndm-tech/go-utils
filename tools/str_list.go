package tools

import (
	"errors"
	"math/rand"

	err "github.com/pkg/errors"
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

func RmDuplicateStrLen(s1 []string, i int) []string {
	s := RmDuplicateStr(s1)
	if len(s) <= i || i < 0 {
		return s
	}
	return s[:i:i]
}

// s1对s2做差
func DifferenceStr(s1, s2 []string) []string {
	s := make([]string, 0, len(s1))
	for _, i := range s1 {
		sign := true
		for _, v := range s2 {
			if i == v {
				sign = false
				break
			}
		}
		if true == sign {
			s = append(s, i)
		}
	}
	return s
}

func DifferenceStrLen(s1, s2 []string, i int) []string {
	s := DifferenceStr(s1, s2)
	if len(s) <= i || i < 0 {
		return s
	}
	return s[:i:i]
}

// Intersect
func IntersectStrList(l1, l2 []string) []string {
	// l2放置较长的list
	var interList []string
	mapTmp := make(map[string]int, len(l2))
	for _, x := range l2 {
		mapTmp[x] = 0
	}
	for _, x := range l1 {
		if _, ok := mapTmp[x]; ok {
			interList = append(interList, x)
		}
	}
	return interList
}

func IntersectStrListLen(l1, l2 []string, Len int) []string {
	// l2放置较长的list
	var interList []string
	mapTmp := make(map[string]int, len(l2))
	for _, x := range l2 {
		mapTmp[x] = 0
	}

	for _, x := range l1 {
		if _, ok := mapTmp[x]; ok {
			interList = append(interList, x)
			if len(interList) == Len {
				return interList
			}
		}
	}
	return interList
}

// string list union
func UnionStrList(s1, s2 []string) []string {
	return append(s1, s2...)
}

func UnionStrListDup(s1, s2 []string) []string {
	return RmDuplicateStr(append(s1, s2...))
}

func UnionStrListLen(s1, s2 []string, i int) []string {
	s := UnionStrList(s1, s2)
	if len(s) <= i || i < 0 {
		return s
	}
	return s[:i:i]
}

// string list
func GetStrListNoLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, errors.New("Input paras error")
	}
	return s[maths.MinInt(size*(num-1), len(s)):maths.MinInt(num*size, len(s))], nil
}

func GetStrListRandN(s1 []string, size int) []string {
	s := append(s1[:0:0], s1...)
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	if size <= 0 || size >= len(s1) {
		return s
	}
	return s[:size:size]
}

func GetStrListLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 || size >= len(s) {
		return []string{}, err.Errorf("Input parameter error!!!, num:%d, size:%d, len(s):%d", num, size, len(s))
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

func MixStrList(s1, s2 []string) []string {
	mix := make([]string, 0, len(s1)+len(s2))
	var i, j = 0, 0
	for i+2 < len(s1) && j+2 < len(s2) {
		mix = append(mix, s1[i:i+2]...)
		mix = append(mix, s2[j:j+2]...)
		i += 2
		j += 2
	}
	mix = append(mix, s1[i:]...)
	mix = append(mix, s2[j:]...)
	return mix
}

func MixListStrV2(l ...[]string) []string {
	// 均匀混合推荐结果
	var count int
	var maxLen int
	for _, s := range l {
		count += len(s)
		maxLen = maths.MaxInt(maxLen, len(s))
	}
	mix := make([]string, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range l {
			if i < len(l[j]) {
				mix = append(mix, l[j][i])
			}
		}
	}
	return RmDuplicateStr(mix)
}

func MixListStrLenV2(Len int, l ...[]string) []string {
	// 均匀混合推荐结果
	var count int
	var maxLen int
	for _, s := range l {
		count += len(s)
		maxLen = maths.MaxInt(maxLen, len(s))
	}
	mix := make([]string, 0, count)
	for i := 0; i < maxLen; i++ {
		for j := range l {
			if i < len(l[j]) {
				mix = append(mix, l[j][i])
			}
		}
	}
	return RmDuplicateStrLen(mix, Len)
}
