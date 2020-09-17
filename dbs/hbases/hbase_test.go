package hbases

//func TestHbaseConnection(t *testing.T) {
//	client := gohbase.NewClient("emr-worker-1.cluster-181564")
//	//fmt.Println(client)
//	//// Values maps a ColumnFamily -> Qualifiers -> Values.
//	//values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
//	//putRequest, err := hrpc.NewPutStr(context.Background(), "item", "comic_info", values)
//	//errs.CheckCommonErr(err)
//	//rsp, err := client.Put(putRequest)
//	//fmt.Println(rsp)
//
//	getRequest, err := hrpc.NewGetStr(context.Background(), "item", "7119")
//	errs.CheckCommonErr(err)
//	getRes, err := client.Get(getRequest)
//	fmt.Println(getRes)
//}

//func TestHbaseGets(t *testing.T) {
//	var hbase = &gohbase.HBHelper{}
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
