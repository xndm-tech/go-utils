package hbases

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/xndm-recommend/go-utils/tools/errs"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cihub/seelog"
	"github.com/silenceper/pool"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gen-go/hbase"
)

type MyHbaseClient struct {
	Client *hbase.THBaseServiceClient
	Trans  thrift.TTransport
}

var ConnPool pool.Pool

func init() {
	poolConfig := &pool.Config{
		InitialCap:  20,
		MaxIdle:     100,
		MaxCap:      300,
		Factory:     CreateClient,
		Close:       CloseClient,
		Ping:        Ping,
		IdleTimeout: 90 * time.Second,
	}
	ConnPool, _ = pool.NewChannelPool(poolConfig)
}

func CloseClient(client interface{}) error {
	if client != nil {
		return client.(MyHbaseClient).Trans.Close()
	}
	return errors.New("连接为空")
}

func Ping(client interface{}) error {
	if client != nil {
		_, err := client.(MyHbaseClient).Client.Exists(context.Background(), []byte("item"), hbase.NewTGet())
		return err
	}
	return errors.New("连接为空")
}

func CreateClient() (interface{}, error) {
	var HOST = "hb-bp1m205ju7p5l24ey-001.hbase.rds.aliyuncs.com"
	var PORT = "9099"
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	trans, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	//trans, err := thrift.NewTSocket(net.JoinHostPort("http://emr-header-1", "16000"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	errs.CheckCommonErr(err)
	if err != nil {
		_ = seelog.Error("hbase连接异常")
		_ = seelog.Error(err)
		return nil, err
	}
	////// 设置用户名密码
	//httClient := trans.(*thrift.THttpClient)
	//httClient.SetHeader("ACCESSKEYID", "root")
	//httClient.SetHeader("ACCESSSIGNATURE", "root")
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening "+HOST, err)
		os.Exit(1)
	}
	return MyHbaseClient{hbase.NewTHBaseServiceClientFactory(trans, protocolFactory), trans}, nil
}
