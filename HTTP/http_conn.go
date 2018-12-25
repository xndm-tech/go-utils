package HTTP

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/xndm-recommend/go-utils/conf_read"
	"github.com/xndm-recommend/go-utils/errors"
)

type HttpInfo struct {
	HttpClient *http.Client
	Url        string
	Para       []string
	TimeOut    int
}

// 线上设置url参数
func (this *HttpInfo) SetUrlPara(values ...interface{}) string {
	var url_tmp string = this.Url
	u, err := url.Parse(url_tmp)
	errors.CheckCommonErr(err)
	for i, val := range values {
		sVal, err := val.(string)
		if false == err {
			sVal = strconv.Itoa(val.(int))
		}
		q := u.Query()
		if len(this.Para) <= i {
			errors.CheckCommonErr(errors.New("Set Url Para error"))
		}
		q.Set(this.Para[i], sVal)
		u.RawQuery = q.Encode()
	}
	return u.String()
}

func GetHttpConnFromConf(this *conf_read.ConfigEngine, SectionName string) *HttpInfo {
	HttpInfo := new(HttpInfo)
	sLogin := getHttpFromConf(this, SectionName)
	createHttpConns(HttpInfo, sLogin)
	return HttpInfo
}
