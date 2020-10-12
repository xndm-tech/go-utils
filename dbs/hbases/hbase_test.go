package hbases

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/filter"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/errs"

	"github.com/xndm-recommend/go-utils/dbs/hbases/gohbase/hrpc"
)

const (
	Config_path = "../../config/test.yaml"
)

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
	getRes, _ := db.GetsByOption("recommender:item", "100012", hrpc.Families(f))
	fmt.Println(getRes)
}

func TestHbaseConnectionFromConfig1(t *testing.T) {
	db := &HBaseDbInfo{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	f := map[string][]string{"comicInfo": []string{"cgender"}}
	f1 := filter.NewPrefixFilter([]byte("1"))
	getRes, _ := db.GetsByScanOption("recommender:item", hrpc.Families(f), hrpc.Filters(f1))
	fmt.Println(getRes)
}
