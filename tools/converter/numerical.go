package converter

import (
	"strconv"
)

/**
int32 转换
*/
func Int32ToInt(i int32) int {
	return int(i)
}

func Int32ToInt64(i int32) int64 {
	return int64(i)
}

/**
int64 转换
*/
func Int64ToInt(i int64) int {
	return int(i)
}

/**
int 转换
*/
func IntToInt32(i int) int32 {
	return int32(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

func Int64ToInt32(i int64) int32 {
	return int32(i)
}

func IntToFloat32(a int) float32 {
	return strconv.Itoa(a)
}

func Int32ToFloat32(a int32) float32 {
	return strconv.Itoa(a)
}

func Int64ToFloat32(a int64) float32 {
	return strconv.Itoa(a)
}
