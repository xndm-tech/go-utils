package converter

import (
	"strconv"

	"github.com/xndm-recommend/go-utils/tools/errs"
	"github.com/xndm-recommend/go-utils/tools/types/nums"
	"github.com/xndm-recommend/go-utils/tools/types/strs"
)

/**
string 转换
*/
func IntToStr(a int) string {
	return strconv.Itoa(a)
}

func Int32ToStr(a int32) string {
	return strconv.FormatInt(int64(a), 10)
}

func Int64ToStr(a int64) string {
	return strconv.FormatInt(a, 10)
}

func FloatToStr(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func IntsToStrs(a []int) []string {
	if nums.IsEmptyInts(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = IntToStr(b)
	}
	return ret
}

func Int32sToStrs(a []int32) []string {
	if nums.IsEmptyInt32s(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = Int32ToStr(b)
	}
	return ret
}

// str2other
func StrToByte(s string) []byte {
	return []byte(s)
}

func StrToInt(a string, defaultInt int) int {
	if r, err := strconv.Atoi(a); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToInt32(a string, defaultInt int32) int32 {
	if r, err := strconv.ParseInt(a, 10, 32); err == nil {
		return int32(r)
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToInt64(a string, defaultInt int64) int64 {
	if r, err := strconv.ParseInt(a, 10, 64); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToFloat32(a string, defaultFloat float32) float32 {
	if r, err := strconv.ParseFloat(a, 32); err == nil {
		return float32(r)
	} else {
		errs.CheckCommonErr(err)
		return defaultFloat
	}
}

func StrToFloat64(a string, defaultInt float64) float64 {
	if r, err := strconv.ParseFloat(a, 64); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrsToInts(a []string, defaultInt int) []int {
	if strs.IsEmptyStrs(a) {
		return nil
	}
	ret := make([]int, len(a))
	for i, b := range a {
		ret[i] = StrToInt(b, defaultInt)
	}
	return ret
}

func StrsToInt32s(a []string, defaultInt int32) []int32 {
	if strs.IsEmptyStrs(a) {
		return nil
	}
	ret := make([]int32, len(a))
	for i, b := range a {
		ret[i] = StrToInt32(b, defaultInt)
	}
	return ret
}
