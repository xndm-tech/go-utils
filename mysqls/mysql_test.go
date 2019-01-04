package mysqls

import (
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/mysqls"
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

	ids := dbinfo.QueryIdList("select cartoon_id,settled_rate from cartoon limit 1")
	b.Log(ids)

	//b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}
