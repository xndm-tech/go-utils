package gohbase_test

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

const (
	Config_path = "../../../config/test.yaml"
)

func TestHbaseConnectionFromConfig1(t *testing.T) {
	db := &gohbase.HBaseDbV2Info{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	getRes, _ := db.GetsByOptions("recommend_samh:alg_nrt_als", []string{"0000abeccf9ebab5:comic", "0004fc27232bd4bf:comic"},
		[]string{"info:user_rowkey", "info:recs"})
	fmt.Println(getRes)
}
