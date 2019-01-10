package type_convert

import "strconv"

func IntToStr(a int) string {
	return strconv.Itoa(a)
}

func StrToInt(a string) (int, error) {
	return strconv.Atoi(a)
}

func Int64ToStr(a int64) string {
	return strconv.FormatInt(a, 10)
}

func StrToInt64(a string) (int64, error) {
	return strconv.ParseInt(a, 10, 64)
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
