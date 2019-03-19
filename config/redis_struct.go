package config

type RedisData struct {
	Master_host []string `yaml:"Master_host"`
	Master_port []string `yaml:"Master_port"`
	Slave_host  []string `yaml:"Slave_host"`
	Slave_port  []string `yaml:"Slave_port"`
	Password    string   `yaml:"Password"`
	Nodes       int      `yaml:"Nodes"`
	Data_time   int      `yaml:"Data_time"`
	Pool_size   int      `yaml:"Pool_size"`
}

func (this *ConfigEngine) GetRedisDataFromConf(name string) *RedisData {
	login := new(RedisData)
	redisLogin := this.GetStruct(name, login)
	return redisLogin.(*RedisData)
}
