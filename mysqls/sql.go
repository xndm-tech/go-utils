package mysqls

import (
	"github.com/xndm-recommend/go-utils/errors_"
)

// 封装用的较多的增删改查功能
func (this *MysqlDbInfo) QueryIdList(sql string) (ids []interface{}) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return ids
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
	return ids
}

// 封装用的较多的增删改查功能
func (this *MysqlDbInfo) QueryStruct(sql string) (ids []interface{}) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return ids
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
	return ids
}
