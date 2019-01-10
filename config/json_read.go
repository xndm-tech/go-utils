package config

import (
	"encoding/json"

	"github.com/xndm-recommend/go-utils/errors_"
)

// 从json string中获得map
func Json_parse(jsonText string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonText), &result)
	errors_.CheckCommonErr(err)
	return
}
