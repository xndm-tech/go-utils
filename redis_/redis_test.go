package redis__test

import (
	"fmt"
	"go-utils/redis_"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

const (
	Config_path = "../config/test.yaml"
)

//func TestMysqlDbInfo_QueryIdList(b *testing.T) {
//	c := config.ConfigEngine{}
//	err := c.Load(Config_path)
//	errors_.CheckCommonErr(err)
//	dbinfo := mysqls.MysqlDbInfo{}
//	dbinfo.GetDbConnFromConf(&c, "Comic_data")
//
//	ids := dbinfo.QueryIdList("select cartoon_id from cartoon limit 1")
//	b.Log(ids)
//
//	sql := "Select cartoon_id From cartoon limit 1"
//	a := strings.ToLower(sql)
//
//	a = tools.SplitStrSep(a, "select", "from")
//	fmt.Println(tools.ContainStrNum(a, ","))
//	//fmt.Println(strings.Contains(a, ","))
//}

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errors_.CheckCommonErr(err)
	dbinfo := new(redis_.RedisItem)
	dbinfo.GetRedisItemFromConf(&c, "Redis_items.Uid_personal_index")
	//dbinfo.GetDbConnFromConf(&c, "Comic_data")
	fmt.Println(dbinfo)
	//StcRec.UidPersonalInterst.GetRedisItemFromConf(yamlConfig, "Redis_items.Uid_personal_interst")
	//StcRec.UidPersonalInterst = GetRedisItemFromConf(yamlConfig, "Redis_items.Uid_personal_interst")
}
