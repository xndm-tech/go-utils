package hbases

import "testing"

const (
	HOST1 = "http://ld-bp17y8n3j6f45p944-proxy-hbaseue.hbaseue.rds.aliyuncs.com:9190"
	// 用户名
	USER1 = "root"
	// 密码
	PASSWORD1 = "root"
)

func TestHBaseThrift211(t *testing.T) {

	xx := new(HBaseDbInfoThrift)
	xx.NewTHttpClient(HOST1, USER1, PASSWORD1)
	xx.GetsByOption2("recommend_samh:item_base", "10000:booklist")

}
