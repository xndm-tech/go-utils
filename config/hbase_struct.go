package config

type HBaseDbData struct {
	ZK        string            `yaml:"ZooKeeperQuorum"`
	User      string            `yaml:"User"`
	Namespace string            `yaml:"Namespace"`
	TableName map[string]string `yaml:"Table_name"`
}

func (this *ConfigEngine) GetHBaseFromConf(name string) *HBaseDbData {
	db := new(HBaseDbData)
	login := this.GetStruct(name, db)
	return login.(*HBaseDbData)
}
