package config

type HttpData struct {
	Url      string   `yaml:"Url"`
	Para     []string `yaml:"Para"`
	Time_out int      `yaml:"Time_out"`
}

func (this *ConfigEngine) GetHttpFromConf(name string) *HttpData {
	login := new(HttpData)
	httpLogin := this.GetStruct(name, login)
	return httpLogin.(*HttpData)
}
