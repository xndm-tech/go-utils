package HTTP

import (
	"errors"
	"net/http"
	"time"

	"github.com/xndm-recommend/go-utils/conf_read"
	"github.com/xndm-recommend/go-utils/errors_"
)

type httpYamklData struct {
	Url      string   `yaml:"Url"`
	Para     []string `yaml:"Para"`
	Time_out int      `yaml:"Time_out"`
}

func getHttpFromConf(this *conf_read.ConfigEngine, sectionName string) *httpYamklData {
	login := new(httpYamklData)
	httpLogin := this.GetStruct(sectionName, login)
	return httpLogin.(*httpYamklData)
}

func createHttpConns(this *HttpInfo, sLogin *httpYamklData) {
	if 0 == sLogin.Time_out {
		errors_.CheckFatalErr(errors.New("can't read http post timeout"))
	}
	this.HttpClient = &http.Client{Timeout: time.Duration(sLogin.Time_out) * time.Millisecond}
	this.Url = sLogin.Url
	for _, param := range sLogin.Para {
		this.Para = append(this.Para, param)
	}
	this.TimeOut = sLogin.Time_out
}
