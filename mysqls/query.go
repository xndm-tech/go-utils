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

func (this *MysqlDbInfo) QueryIdList(sql string) (ids []string) {
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

func (this *MysqlDbInfo) QueryIdIntList(sql string) (ids []int) {
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
		var tmpId int
		err := rows.Scan(&tmpId)
		errors_.CheckCommonErr(err)
		if nil == err {
			ids = append(ids, tmpId)
		}
	}
	return
}

func (this *MysqlDbInfo) QueryIdListLen(sql string, len int) (ids []string) {
	if checkNumSql(sql) != 0 {
		return
	}
	stmt, err := this.SqlDataDb.Prepare(sql + " LIMIT ?")
	defer stmt.Close()
	errors_.CheckCommonErr(err)
	row, err := stmt.Query(len)
	defer row.Close()
	errors_.CheckCommonErr(err)
	for row.Next() {
		var tmpId string
		err = row.Scan(&tmpId)
		errors_.CheckCommonErr(err)
		ids = append(ids, tmpId)
	}
	return
}

func (this *MysqlDbInfo) QueryStruct(sql string, pars ...interface{}) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(pars...)
	errors_.CheckCommonErr(err)
	return
}

func (this *MysqlDbInfo) QueryIdMap(sql string) (mOut map[string]string, err error) {
	mOut = make(map[string]string, 0)
	if checkNumSql(sql) != 1 {
		return
	}
	// 查询数据
	var key, val string
	row, err := this.SqlDataDb.Query(sql)
	defer row.Close()
	errors_.CheckCommonErr(err)
	for row.Next() {
		err = row.Scan(&key, &val)
		errors_.CheckCommonErr(err)
		mOut[key] = val
	}
	return
}
