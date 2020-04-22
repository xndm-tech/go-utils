package candidates

import (
	"fmt"
	"testing"

	"github.com/xndm-recommend/go-utils/candidates"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/errs"
	"github.com/xndm-recommend/go-utils/mysqls"
)

const (
	Config_path = "../config/test.yaml"
)

func TestCandidate(b *testing.T) {
	c := config.ConfigEngine{}
	err := c.Load(Config_path)
	errs.CheckCommonErr(err)
	dbinfo := mysqls.MysqlDbInfo{}
	dbinfo.GetDbConnFromConf(&c, "Comic_data")

	can := candidates.Candidate{}
	can.GenCandidateFromdb(&dbinfo, "select settled_rate from cartoon limit 1")
	b.Log(can)

	//b.Logf("hit: %d miss: %d ratio: %f", hit, miss, float64(hit)/float64(miss))
}

func TestCandidate_test(b *testing.T) {
	a := "asfdsa"
	var tmp interface{}
	tmp = a
	fmt.Println(tmp)
}
