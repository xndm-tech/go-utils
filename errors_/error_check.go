package errors_

/*
有关报错打印的封装
*/
import (
	"reflect"
	"runtime"

	"github.com/cihub/seelog"
)

//错误处理函数
func CheckCommonErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		seelog.Warn(file, ":", line, err)
	}
}

//错误处理函数
func CheckFatalErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		seelog.Error("Important error:", file, ":", line, err)
		panic(err)
	}
}

func CheckEmptyValue(val interface{}) {
	if reflect.TypeOf(val).Kind() == reflect.Int {
		if val.(int) == 0 {
			panic("this value shouldn't be 0")
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Int64 {
		if val.(int64) == 0 {
			panic(`this value shouldn't be 0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.String {
		if val.(string) == "" {
			panic(`this value shouldn't be ""`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float32 {
		if val.(float32) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float64 {
		if val.(float64) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Slice {
		if len(val.([]interface{})) == 0 {
			panic(`this value shouldn't be empty slice`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Map {
		if len(val.(map[interface{}]interface{})) == 0 {
			panic(`this value shouldn't be empty map`)
		}
	}
}
