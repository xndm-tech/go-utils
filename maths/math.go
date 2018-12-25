package maths

import (
	"fmt"
	"reflect"
)

func Min(a, b interface{}) interface{} {
	a_val := reflect.ValueOf(a)
	b_val := reflect.ValueOf(b)
	if a_val.Kind() != b_val.Kind() {
		return fmt.Errorf("input num type diff!!!")
	} else if a_val.Kind() == reflect.Int {
		if a.(int) < b.(int) {
			return a
		} else {
			return b
		}
	} else if a_val.Kind() == reflect.Float32 {
		if a.(float32) < b.(float32) {
			return a
		} else {
			return b
		}
	} else if a_val.Kind() == reflect.Float64 {
		if a.(float64) < b.(float64) {
			return a
		} else {
			return b
		}
	} else {
		return fmt.Errorf("input num type not support!!!")
	}
}

func Max(a, b interface{}) interface{} {
	a_val := reflect.ValueOf(a)
	b_val := reflect.ValueOf(b)
	if a_val.Kind() != b_val.Kind() {
		return fmt.Errorf("input num type diff!!!")
	} else if a_val.Kind() == reflect.Int {
		if a.(int) > b.(int) {
			return a
		} else {
			return b
		}
	} else if a_val.Kind() == reflect.Float32 {
		if a.(float32) > b.(float32) {
			return a
		} else {
			return b
		}
	} else if a_val.Kind() == reflect.Float64 {
		if a.(float64) > b.(float64) {
			return a
		} else {
			return b
		}
	} else {
		return fmt.Errorf("input num type not support!!!")
	}
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinFloat(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func MaxFloat(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}
