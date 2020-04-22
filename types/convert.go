package types

import (
	"strconv"

	"github.com/xndm-recommend/go-utils/errs"
)

func IntToStr(a int) string {
	return strconv.Itoa(a)
}

func StrToInt(a string, default_int int) int {
	if r, err := strconv.Atoi(a); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return default_int
	}
}

func Int64ToStr(a int64) string {
	return strconv.FormatInt(a, 10)
}

func StrToInt64(a string, default_int int64) int64 {
	if r, err := strconv.ParseInt(a, 10, 64); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return default_int
	}
}

func FloatToStr(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func MapToStringInt(interfa map[interface{}]interface{}) map[string]int {
	change := make(map[string]int)
	for key, val := range interfa {
		change[key.(string)] = val.(int)
	}
	return change
}

func MapToStringString(interfa map[interface{}]interface{}) map[string]string {
	change := make(map[string]string)
	for key, val := range interfa {
		change[key.(string)] = val.(string)
	}
	return change
}
