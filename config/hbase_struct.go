package config

type HBaseDbData struct {
	ZK        string            `yaml:"ZooKeeperQuorum"`
	User      string            `yaml:"User"`
	Namespace string            `yaml:"Namespace"`
	QueueSize int               `yaml:"QueueSize"`
	TableName map[string]string `yaml:"Table_name"`
}

type HBaseDbV2Data struct {
	Thrift    string            `yaml:"Thrift_address"`
	Namespace string            `yaml:"Namespace"`
	TableName map[string]string `yaml:"Table_name"`
}

func (this *ConfigEngine) GetHBaseFromConf(name string) *HBaseDbData {
	db := new(HBaseDbData)
	login := this.GetStruct(name, db)
	return login.(*HBaseDbData)
}

func (this *ConfigEngine) GetHBaseV2FromConf(name string) *HBaseDbV2Data {
	db := new(HBaseDbV2Data)
	login := this.GetStruct(name, db)
	return login.(*HBaseDbV2Data)
}
