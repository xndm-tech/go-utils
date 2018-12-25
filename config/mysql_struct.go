package config

type MysqlDbData struct {
	User       string            `yaml:"user"`
	Password   string            `yaml:"password"`
	Host       string            `yaml:"host"`
	Port       string            `yaml:"port"`
	Db_name    string            `yaml:"db_name"`
	Table_name map[string]string `yaml:"table_name"`
	Max_conns  int               `yaml:"max_conns"`
	Time_out   int               `yaml:"time_out"`
}

func (this *ConfigEngine) GetMySqlFromConf(name string) *MysqlDbData {
	mysqlLogin := new(MysqlDbData)
	login := this.GetStruct(name, mysqlLogin)
	return login.(*MysqlDbData)
}
