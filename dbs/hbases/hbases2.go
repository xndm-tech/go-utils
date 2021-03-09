package hbases

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/tsuna/gohbase"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gen-go/hbase"
)

type HBaseDbInfoThrift struct {
	_host       string
	_user       string
	_passwd     string
	_defaultCtx context.Context
	_client     *hbase.THBaseServiceClient
}

func (hb *HBaseDbInfoThrift) NewTHttpClient(host, user, passwd string, options ...gohbase.Option) {
	hb._defaultCtx = context.Background()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	trans, err := thrift.NewTHttpClient(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	// 设置用户名密码
	httClient := trans.(*thrift.THttpClient)
	httClient.SetHeader("ACCESSKEYID", user)
	httClient.SetHeader("ACCESSSIGNATURE", passwd)
	hb._client = hbase.NewTHBaseServiceClientFactory(trans, protocolFactory)
}

//func (hb *HBaseDbInfoThrift) connectHBase2(db *config.HBaseDbData) {
//	hb._host = db.ZK
//	hb.Namespace = db.Namespace
//	hb.TableName = db.TableName
//	hb._client = gohbase.NewClient(db.ZK, gohbase.RpcQueueSize(db.QueueSize), gohbase.CompressionCodec("snappy"))
//}
//
////通过hb.PutsByRowkeyVersion(table, rowkey, values, hrpc.Timestamp(timestamp))调用，其中timestamp是time.Time类型，options也可以是其他 func(hrpc.Call)的函数
//func (hb *HBaseDbInfoThrift) PutsByRowkeyVersion2(table, rowKey string, values map[string]map[string][]byte, options func(hrpc.Call) error) (err error) {
//	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values, options)
//	errs.CheckCommonErr(err)
//	_, err = hb._client.Put(putRequest)
//	errs.CheckCommonErr(err)
//
//	return
//}
//

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbInfoThrift) GetsByOption2(table, rowkey string) (*hbase.TResult_, error) {
	// 单行查询数据
	result, err := hb._client.Get(hb._defaultCtx, []byte(table), &hbase.TGet{Row: []byte(rowkey)})
	//fmt.Println("Get result:")
	//fmt.Println(result)
	return result, err
}

//
////指定表，通过options筛选数据，例如Families函数，或者filter函数
//func (hb *HBaseDbInfoThrift) GetsByScanOption2(table string, options ...func(hrpc.Call) error) (rsp []*hrpc.Result, err error) {
//	var (
//		scanRequest *hrpc.Scan
//		res         *hrpc.Result
//		err1        error
//	)
//	scanRequest, err = hrpc.NewScanStr(context.Background(), table, options...)
//	if nil != err {
//		errs.CheckCommonErr(err)
//		return nil, err
//	}
//	scanner := hb._client.Scan(scanRequest)
//	for {
//		res, err1 = scanner.Next()
//		if err1 == io.EOF || res == nil {
//			break
//		}
//		if nil != err1 {
//			errs.CheckCommonErr(err1)
//			return nil, err1
//		}
//		rsp = append(rsp, res)
//	}
//	return
//}
//
//func (this *HBaseDbInfoThrift) Ping2() error {
//	if this._client != nil {
//		return nil
//	}
//	return errors.New("连接为空")
//}
//
//func (this *HBaseDbInfoThrift) GetClient2() gohbase.Client {
//	return this._client
//}
//
//func (this *HBaseDbInfoThrift) GetDbConnFromConf2(c *config.ConfigEngine, name string) {
//	this.connectHBase2(c.GetHBaseFromConf(name))
//}
