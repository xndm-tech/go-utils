package mysqls

import (
	"strings"

	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/tools"
)

// 封装用的较多的增删改查功能
func checkNumSql(sql string) int {
	sql_low := strings.ToLower(sql)
	return tools.ContainStrNum(tools.SplitStrSep(sql_low, "select", "from"), ",")
}

func (this *MysqlDbInfo) QueryIdList(sql string) (ids []interface{}) {
	if checkNumSql(sql) != 0 {
		return
	}
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tmpId string
		err := rows.Scan(&tmpId)
		errors_.CheckCommonErr(err)
		if nil == err {
			ids = append(ids, tmpId)
		}
	}
	return
}

func (this *MysqlDbInfo) QueryStruct(sql string, out ...*interface{}) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return
	}
	defer rows.Close()
	for _, o := range out {
		var tmp string
		err = rows.Scan(&tmp)
		errors_.CheckCommonErr(err)
		if nil == err {
			*o = tmp
		}
	}
	return
}
