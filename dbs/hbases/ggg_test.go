package hbases_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gen-go/hbase"

	"github.com/xndm-recommend/go-utils/dbs/hbases"
)

func TestHbaseGets(t *testing.T) {
	defaultCtx := context.Background()
	cli, _ := hbases.CreateClient()
	fmt.Println(hbases.ConnPool)
	t.Log(cli)
	t.Log(hbases.Ping(cli))
	kkk := cli.(hbases.MyHbaseClient)
	jj, err := kkk.Client.Get(defaultCtx, []byte("recommend_samh:alg_nrt_als"),
		&hbase.TGet{Row: []byte("0004fc27232bd4bf:comic")})
	t.Log(err)
	t.Log(jj)
}

//
//// 单行查询数据
//result, err := client.Get(defaultCtx, tableInbytes, &hbase.TGet{Row: []byte("row1")})
//fmt.Println("Get result:")
//fmt.Println(result)
