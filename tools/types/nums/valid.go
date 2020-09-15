package nums

import (
	"github.com/xndm-recommend/go-utils/common/consts"
)

func IsEmptyInts(s []int) bool {
	return len(s) == consts.ZERO
}

func IsEmptyInt32s(s []int32) bool {
	return len(s) == consts.ZERO
}
