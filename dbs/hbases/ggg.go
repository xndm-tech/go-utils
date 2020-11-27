package hbases

import (
	"context"
	"errors"
	"fmt"
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

var connPool pool.Pool

func init() {
	poolConfig := &pool.Config{
		InitialCap:  20,
		MaxIdle:     100,
		MaxCap:      300,
		Factory:     createClient,
		Close:       closeClient,
		Ping:        ping,
		IdleTimeout: 90 * time.Second,
	}
	connPool, _ = pool.NewChannelPool(poolConfig)
}

func closeClient(client interface{}) error {
	if client != nil {
		return client.(MyHbaseClient).Trans.Close()
	}
	return errors.New("连接为空")
}

func ping(client interface{}) error {
	if client != nil {
		_, err := client.(MyHbaseClient).Client.Exists(context.Background(), []byte("item"), hbase.NewTGet())
		return err
	}
	return errors.New("连接为空")
}

func createClient() (interface{}, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	trans, err := thrift.NewTHttpClient("http://emr-header-1:16000")
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
	// 设置用户名密码
	httClient := trans.(*thrift.THttpClient)
	httClient.SetHeader("ACCESSKEYID", "root")
	httClient.SetHeader("ACCESSSIGNATURE", "root")
	errs.CheckCommonErr(err)
	return MyHbaseClient{hbase.NewTHBaseServiceClientFactory(trans, protocolFactory), trans}, nil
}
