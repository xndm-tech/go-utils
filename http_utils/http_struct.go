package http_utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/zhanglanhui/go-utils/utils/conf_utils"
	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

type httpYamklData struct {
	Url      string   `yaml:"Url"`
	Para     []string `yaml:"Para"`
	Time_out int      `yaml:"Time_out"`
}

func getHttpFromConf(this *conf_utils.ConfigEngine, SectionName string) *httpYamklData {
	login := new(httpYamklData)
	httpLogin := this.GetStruct(SectionName, login)
	return httpLogin.(*httpYamklData)
}

func createHttpConns(this *HttpInfo, sLogin *httpYamklData) {
	if 0 == sLogin.Time_out {
		err_utils.CheckFatalErr(errors.New("can't read http post timeout"))
	}
	this.HttpClient = &http.Client{Timeout: time.Duration(sLogin.Time_out) * time.Millisecond}
	this.Url = sLogin.Url
	for _, param := range sLogin.Para {
		this.Para = append(this.Para, param)
	}
	this.TimeOut = sLogin.Time_out
}
