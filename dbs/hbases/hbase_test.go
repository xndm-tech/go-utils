package hbases

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/hrpc"
)

func TestHbaseConnection(t *testing.T) {
	db := &HBHelper{}
	db.ConnectHBase("recommender/user@EMR.181564.COM",
		"emr-worker-1.cluster-181564,emr-worker-2.cluster-181564,emr-header-1.cluster-181564")
	t.Log("db", db)
	//// Values maps a ColumnFamily -> Qualifiers -> Values.
	//values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
	//putRequest, err := hrpc.NewPutStr(context.Background(), "item", "comic_info", values)
	//errs.CheckCommonErr(err)
	//rsp, err := client.Put(putRequest)
	//fmt.Println(rsp)

	//getRequest, err := hrpc.NewGetStr(context.Background(), "item", "7119")
	//errs.CheckCommonErr(err)
	f := map[string][]string{"comicInfo": []string{"cid"}}
	getRes, _ := db.GetsByOption("item", "7119", hrpc.Families(f))
	fmt.Println(getRes)
}

//func TestHbaseGets(t *testing.T) {
//	var hbase = &gohbase1.HBHelper{}
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
