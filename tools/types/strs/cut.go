package strs

import (
	"fmt"

	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/tools/maths"
)

// cut strs
func CutStrList(s []string, l int) []string {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

// cut strs
func CutStrListAndFilling(s []string, f []string, l int) []string {
	if cut := CutStrList(s, l); len(cut) < l {
		tmp := append(s, f...)
		return UniqueStrsLen(tmp, l)
	} else {
		return cut
	}
}

// string list Loop
func GetStrListNoLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, fmt.Errorf("Input paras error!!!")
	}
	minThr := maths.MinInt(size*(num-1), len(s))
	maxThr := maths.MinInt(size*num, len(s))
	return s[minThr:maxThr], nil
}

func GetStrListLoop(s []string, size, num int) ([]string, error) {
	if num <= 0 || size <= 0 || size >= len(s) {
		return []string{}, fmt.Errorf("Input parameter error!!!")
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
