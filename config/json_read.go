package config

import (
	"encoding/json"

	"github.com/xndm-recommend/go-utils/tools/logs"
)

// 从json string中获得map
func Json_parse(jsonText string) (result map[string]interface{}) {
	result = make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonText), &result)
	logs.CheckCommonErr(err)
	return
}
