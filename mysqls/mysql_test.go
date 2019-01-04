package mysqls

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xndm-recommend/go-utils/mysqls"

	"github.com/xndm-recommend/go-utils/tools"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	//"github.com/xndm-recommend/go-utils/mysqls"
)

const (
	Config_path = "../config/test.yaml"
)

func TestMysqlDbInfo_QueryIdList(b *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errors_.CheckCommonErr(err)
	dbinfo := mysqls.MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	ids := dbinfo.QueryIdList("select cartoon_id from cartoon limit 1")
	b.Log(ids)

	sql := "Select cartoon_id From cartoon limit 1"
	a := strings.ToLower(sql)

	a = tools.SplitStrSep(a, "select", "from")
	fmt.Println(tools.ContainStrNum(a, ","))
	//fmt.Println(strings.Contains(a, ","))
}

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errors_.CheckCommonErr(err)
	dbinfo := MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	var cartoon_id1, cartoon_id2 string
	dbinfo.QueryStruct("select cartoon_id,cartoon_id from cartoon limit 1", &cartoon_id1, &cartoon_id2)
	t.Log(cartoon_id1, cartoon_id2)
}
