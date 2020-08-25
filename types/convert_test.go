package types_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/types"
)

func TestStrToFloat32(b *testing.T) {
	a := "1.2"
	fmt.Println(types.StrToFloat32(a, 9))
}

func TestStrToFloat64(b *testing.T) {
	a := "1.2"
	fmt.Println(types.StrToFloat64(a, 9))
}
