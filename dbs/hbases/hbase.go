package hbases

import (
	"context"
	"errors"
	"fmt"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/hrpc"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

type HBHelper struct {
	Zkquorum string
	Option   string
	_client  gohbase.Client
}

func (hb *HBHelper) ConnectHBase(account string, zkquorum string) {
	auth := gohbase.Auth("KERBEROS")
	user := gohbase.EffectiveUser(account)
	options := []gohbase.Option{auth, user}
	hb.Zkquorum = zkquorum
	hb._client = gohbase.NewClient(zkquorum, options...)
}

//通过hb.PutsByRowkeyVersion(table, rowkey, values, hrpc.Timestamp(timestamp))调用，其中timestamp是time.Time类型，options也可以是其他 func(hrpc.Call)的函数
func (hb *HBHelper) PutsByRowkeyVersion(table, rowKey string, values map[string]map[string][]byte, options func(hrpc.Call) error) (err error) {
	putRequest, err := hrpc.NewPutStr(context.Background(), table, rowKey, values, options)
	errs.CheckCommonErr(err)
	_, err = hb._client.Put(putRequest)
	errs.CheckCommonErr(err)

	return
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBHelper) GetsByOption(table string, rowkey string, options func(hrpc.Call) error) (*hrpc.Result, error) {
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey, options)
	errs.CheckCommonErr(err)
	res, err := hb._client.Get(getRequest)
	errs.CheckCommonErr(err)
	defer func() {
		if err := recover(); err != nil {
			switch fmt.Sprintf("%v", err) {
			case "runtime error: index out of range":
				err = errors.New("NoSuchRowKeyOrQualifierException")
			case "runtime error: invalid memory address or nil pointer dereference":
				err = errors.New("NoSuchColFamilyException")
			default:
				err = fmt.Errorf("%v", err)
			}
		}
	}()
	return res, nil
}
