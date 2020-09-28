package hbases

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/errs"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/hrpc"
)

const (
	Config_path = "../../config/test.yaml"
)

func TestHbaseConnection(t *testing.T) {
	db := &HBaseDbInfo{}
	db.ConnectHBase("recommender/user@EMR.181564.COM",
		"emr-worker-1.cluster-181564,emr-worker-2.cluster-181564,emr-header-1.cluster-181564")
	t.Log("db", db)
	t.Log(db.Ping())
	f := map[string][]string{"comicInfo": {"cid"}}
	getRes, _ := db.GetsByOption("default:item", "7119", hrpc.Families(f))
	fmt.Println(getRes)
}

func TestHbaseConnectionFromConfig(t *testing.T) {
	db := &HBaseDbInfo{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	f := map[string][]string{"comicInfo": []string{"cid"}}
	getRes, _ := db.GetsByOption("item", "7119", hrpc.Families(f))
	fmt.Println(getRes)
}

//func TestHbaseGets(t *testing.T) {
//	var hbase = &gohbase1.HBaseDbInfo{}
//	hbase.ConnectHBase("emr-worker-1.cluster-181564", "root-region-server")
//	t.Log("hbase._client", hbase._client)
//	f := map[string][]string{"comicInfo": []string{"cid"}}
//	getres, err := hbase.GetsByOption("item", "7119", hrpc.Families(f))
//	t.Log(err)
//	for i, v := range getres.Cells {
//		row := string(v.Row[:])
//		fam := string(v.Family[:])
//		qua := string(v.Qualifier[:])
//		value := string(v.Value[:])
//		str := fmt.Sprintf("index:%d		Rowkey: %s		family: %s		qualifies:%s		value:%s", i, row, fam, qua, value)
//		fmt.Println(str)
//	}
//}
