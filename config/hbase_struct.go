package config

type HBaseDbData struct {
	ZK   string `yaml:"ZooKeeperQuorum"`
	User string `yaml:"User"`
}

func (this *ConfigEngine) GetHBaseFromConf(name string) *HBaseDbData {
	db := new(HBaseDbData)
	login := this.GetStruct(name, db)
	return login.(*HBaseDbData)
}
