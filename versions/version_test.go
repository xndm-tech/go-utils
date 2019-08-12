package versions

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
)

const (
	Config_path = "../config/test.yaml"
)

func TestMysqlDbInfo_QueryStruct(t *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errors_.CheckCommonErr(err)
	dbinfo := Version{}
	dbinfo.GeVersionFromConf(&c, "Version")
	fmt.Println(dbinfo.GetAlgoVersion())
}
