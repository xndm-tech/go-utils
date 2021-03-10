package main

import (
	"fmt"
	"time"

	"github.com/popeyeio/gohbase/lib/thrift"

	"github.com/popeyeio/gohbase/gen/hbase"
	"github.com/popeyeio/gohbase/pool"
)

func main() {
	hbasePool := NewHbasePool()
	defer hbasePool.Close()

	hbaseClient, err := hbasePool.Get()
	if err != nil {
		fmt.Printf("Get error - %v\n", err)
		return
	}
	defer hbaseClient.Close()

	aa, err := GetRowxxx(hbaseClient)
	fmt.Printf("GetRowxxx aa - %v\n", aa)
	fmt.Printf("GetRowxxx error - %v\n", err)
}

func NewHbasePool() pool.Pool {
	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//trans, err := thrift.NewTHttpClient("ld-bp17y8n3j6f45p944-proxy-hbaseue.hbaseue.rds.aliyuncs.com:9190")
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, "error resolving address:", err)
	//	os.Exit(1)
	//}
	//// 设置用户名密码
	//httClient := trans.(*thrift.THttpClient)
	//httClient.SetHeader("ACCESSKEYID", "root")
	//httClient.SetHeader("ACCESSSIGNATURE", "root")
	kkk := thrift.NewTHttpClientTransportFactory("http://ld-bp17y8n3j6f45p944-proxy-hbaseue.hbaseue.rds.aliyuncs.com:9190")
	opts := []pool.Option{
		pool.WithTransportFactory(kkk),
		pool.WithUpdatePickerInterval(time.Second * 10),
		pool.WithSocketTimeout(time.Second * 5),
		pool.WithMaxActive(8),
		pool.WithMaxIdle(8),
		pool.WithIdleTimeout(time.Second * 5),
		pool.WithCleanUpInterval(time.Second * 30),
		pool.WithBlockMode(true),
	}
	return pool.NewPool(opts...)
}

func GetRowxxx(cli pool.Client) ([]*hbase.TRowResult_, error) {
	return cli.GetRow("recommend_samh:alg_nrt_ar_fpgrowth", "5c341a7dda910511:comic", nil)
}
