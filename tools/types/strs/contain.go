package strs

import (
	"reflect"
	"strings"

	"github.com/xndm-recommend/go-utils/common/consts"
)

// split
func ContainStrNum(str string, sep string) int {
	return len(strings.Split(str, sep)) - 1
}

func IsContainStr(slice []string, item string) bool {
	if consts.ZERO == len(slice) {
		return false
	}
	for _, sItem := range slice {
		if item == sItem {
			return true
		}
	}
	return false
}

func IsNotContainStr(slice []string, item string) bool {
	return !IsContainStr(slice, item)
}

// Returns true if slice strings contains any of the strings.
func IsContainAny(slice []string, checks []string) bool {
	for _, s := range checks {
		if IsContainStr(slice, s) {
			return true
		}
	}
	return false
}

func IsContainsMap(maps []map[string]string, item map[string]string) bool {
	for _, m := range maps {
		if reflect.DeepEqual(m, item) {
			return true
		}
	}
	return false
}
