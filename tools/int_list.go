package tools

import (
	"math/rand"
)

func NRandInt(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func IntToInterfaces(i []int) []interface{} {
	ifs := make([]interface{}, len(i))
	for ind, v := range i {
		ifs[ind] = v
	}
	return ifs
}

func InterfacesToInt(i []interface{}) []int {
	ints := make([]int, len(i))
	for ind, v := range i {
		ints[ind] = v.(int)
	}
	return ints
}
