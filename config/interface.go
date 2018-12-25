package config

type ConfRead interface {
	Load(path string) error

	Get(name string) interface{}

	GetString(name string) string

	GetInt(name string) int

	GetBool(name string) bool

	GetFloat64(name string) float64

	GetStructInt(name string, s string) int

	GetStructStr(name string, s string) string

	GetStruct(name string, s interface{}) interface{}

	GetMySqlFromConf(sectionName string) *MysqlDbYamlData
}