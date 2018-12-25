package versions

/*
有关yaml配置文件的封装,读取算法版本号
*/
type Version struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}
