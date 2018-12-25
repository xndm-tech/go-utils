package HTTP

/*
有关Http协议GET和POST请求的封装
*/
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/zhanglanhui/go-utils/utils/err_utils"
)

//发送GET请求
//url:请求地址
//response:请求返回的内容
func HttpGet(client *http.Client, url string) (response string, ok bool) {
	resp, err := client.Get(url)
	if err != nil {
		err_utils.CheckCommonErr(err)
		return "", false
	}
	defer resp.Body.Close()
	if 200 == resp.StatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		err_utils.CheckCommonErr(err)
		return string(body), true
	} else {
		return "", false
	}
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json_struct
//content:请求放回的内容
func HttpPost(client *http.Client, url string, data interface{}, contentType string) (content string) {
	jsonStr, err := json.Marshal(data)
	err_utils.CheckCommonErr(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	err_utils.CheckCommonErr(err)
	req.Header.Set("Content-Type", contentType)
	defer req.Body.Close()
	resp, err := client.Do(req)
	err_utils.CheckCommonErr(err)
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	err_utils.CheckCommonErr(err)
	return string(result)
}
