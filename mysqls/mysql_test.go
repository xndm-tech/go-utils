package mysqls_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errs"
	. "github.com/xndm-recommend/go-utils/mysqls"
	"github.com/xndm-recommend/go-utils/tools"
)

const (
	Config_path = "../config/test.yaml"
)

func TestMysqlDbInfo_QueryIdList(b *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	db := MysqlDbInfo{}
	db.GetDbConnFromConf(&c, "Comic_data")
	ids, _ := db.SelectStrList("select cartoon_id from cartoon limit 1")
	b.Log(ids)
	sql := "Select cartoon_id From cartoon limit 1"
	a := strings.ToLower(sql)
	a = tools.SplitStrSep(a, "select", "from")
	fmt.Println(tools.ContainStrNum(a, ","))
}

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	dbinfo := MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	//var cartoon_id1, cartoon_id2 int

	var names []string
	_ = dbinfo.GetDb().Select(&names, "SELECT cartoon_id FROM cartoon LIMIT 10")
	t.Log(names)
	t.Log(dbinfo.SelectStrList("select cartoon_id from cartoon limit ?", "10"))
}
