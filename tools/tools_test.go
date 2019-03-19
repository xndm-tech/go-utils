package tools

import (
	"fmt"
	"testing"
	//"github.com/xndm-recommend/go-utils/mysqls"
)

const (
	Config_path = "../config/test.yaml"
)

func TestContainStrNum(b *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a[:5:5])
}
