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

func getOneRow(data []*Hbase.TRowResult) map[string]string {
	if data == nil {
		return nil
	}
	var rowValue = make(map[string]string)
	for _, x := range data {
		for k, v := range x.Columns {
			rowValue[k] = string(v.Value)
		}
		break
	}
	return rowValue
}

func getRows(data []*Hbase.TRowResult) []map[string]string {
	if data == nil {
		return nil
	}
	var rowsValue = make([]map[string]string, consts.ZERO, len(data))
	for _, x := range data {
		var value = make(map[string]string)
		for k, v := range x.Columns {
			value[k] = string(v.Value)
		}
		rowsValue = append(rowsValue, value)
	}
	return rowsValue
}

func (hb *HBaseDbV2Info) ConnectHBase(address string) {
	client, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		errs.CheckFatalErr(err)
		return
	}
	hb.address = address
	hb._client = client
}

func (hb *HBaseDbV2Info) connectHBase(db *config.HBaseDbV2Data) {
	client, err := goh.NewTcpClient(db.Thrift, goh.TBinaryProtocol, false)
	if err != nil {
		errs.CheckFatalErr(err)
		return
	}
	hb.address = db.Thrift
	hb._client = client
	hb.Namespace = db.Namespace
	hb.TableName = db.TableName
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbV2Info) GetsByOption(table, rowkey string, columns []string) (map[string]string, error) {
	if err := hb._client.Open(); err != nil {
		return nil, err
	}
	defer hb._client.Close()
	if data, err := hb._client.GetRowWithColumns(table, []byte(rowkey), columns, nil); err != nil {
		return nil, err
	} else {
		return getOneRow(data), nil
	}
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (hb *HBaseDbV2Info) GetsByOptions(table string, rowkeys []string, columns []string) ([]map[string]string, error) {
	var rows = make([][]byte, len(rowkeys))
	for i, k := range rowkeys {
		rows[i] = []byte(k)
	}
	if err := hb._client.Open(); err != nil {
		return nil, err
	}
	defer hb._client.Close()
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
