package gohbase

import (
	"errors"

	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

type HBaseDbV2Info struct {
	address   string
	_client   *goh.HClient
	Namespace string
	TableName map[string]string
}

func getOneRow(data []*Hbase.TRowResult) []string {
	if data == nil {
		return nil
	}
	var rowValue = make([]string, consts.ZERO)
	for _, x := range data {
		for _, v := range x.Columns {
			rowValue = append(rowValue, string(v.Value))
		}
		break
	}
	return rowValue
}

func getRows(data []*Hbase.TRowResult) [][]string {
	if data == nil {
		return nil
	}
	var rowsValue = make([][]string, consts.ZERO)
	for _, x := range data {
		var value = make([]string, consts.ZERO)
		for _, v := range x.Columns {
			value = append(value, string(v.Value))
		}
		rowsValue = append(rowsValue, value)
	}
	return rowsValue
}

func (hb *HBaseDbV2Info) ConnectHBase(address string) {
	client, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		errs.CheckCommonErr(err)
		return
	}
	hb = &HBaseDbV2Info{
		address: address,
		_client: client,
	}
}

func (hb *HBaseDbV2Info) connectHBase(db *config.HBaseDbV2Data) {
	client, err := goh.NewTcpClient(db.Thrift, goh.TBinaryProtocol, false)
	if err != nil {
		errs.CheckCommonErr(err)
		return
	}
	hb = &HBaseDbV2Info{
		address:   db.Thrift,
		_client:   client,
		Namespace: db.Namespace,
		TableName: db.TableName,
	}
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbV2Info) GetsByOption(table, rowkey string, columns []string) ([]string, error) {
	if data, err := hb._client.GetRowWithColumns(table, []byte(rowkey), columns, nil); err != nil {
		return nil, err
	} else {
		return getOneRow(data), nil
	}
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbV2Info) GetsByOptions(table string, rowkeys []string, columns []string) ([][]string, error) {
	var rows = make([][]byte, len(rowkeys))
	for i, k := range rowkeys {
		rows[i] = []byte(k)
	}
	if data, err := hb._client.GetRowsWithColumns(table, rows, columns, nil); err != nil {
		return nil, err
	} else {
		return getRows(data), nil
	}
}

func (this *HBaseDbV2Info) Ping() error {
	if this._client != nil {
		return nil
	}
	return errors.New("连接为空")
}

func (this *HBaseDbV2Info) GetClient() *goh.HClient {
	return this._client
}

func (this *HBaseDbV2Info) GetDbConnFromConf(c *config.ConfigEngine, name string) {
	this.connectHBase(c.GetHBaseV2FromConf(name))
}
