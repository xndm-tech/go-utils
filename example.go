package main

import (
	"fmt"

	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/mysqls"
)

const (
	Config_path = "config/test.yaml"
)

func main() {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errors_.CheckCommonErr(err)
	dbinfo := mysqls.MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	//var cartoon_id1, cartoon_id2 string
	//var shoucang string

	fmt.Println(dbinfo.QueryIdList("select cartoon_id,cartoon_name,shoucang from cartoon limit 1"))
}
