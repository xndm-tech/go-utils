package gohbase

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/errs"
)

const (
	Config_path = "../../config/test.yaml"
)

func TestHbaseConnectionFromConfig1(t *testing.T) {
	db := &HBaseDbV2Info{}
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db.GetDbConnFromConf(&c, "HBase_db2")
	t.Log(db.Ping())
	t.Log("db", db)
	t.Log("db", db.TableName)
	getRes, _ := db.GetsByOption("recommender:recall_samh_alg_nrt_comicsim_cb", "0000abeccf9ebab5:comic", []string{"recs"})
	fmt.Println(getRes)
}
