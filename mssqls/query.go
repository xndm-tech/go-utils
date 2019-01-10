package mssqls

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

func (this *MssqlDbInfo) QueryIdList(sql string) (ids []string, err error) {

	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		rows, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
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

func (this *MssqlDbInfo) QueryIdIntList(sql string) (ids []int, err error) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		rows, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
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
func (this *MssqlDbInfo) QueryIdListLen(sql string, len int) (ids []string, err error) {
	stmt, err := this.SqlDataDb.Prepare(sql + " LIMIT ?")
	defer stmt.Close()
	errors_.CheckCommonErr(err)
	row, err := stmt.Query(len)
	if err != nil {
		row, err = stmt.Query(len)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
	}
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

func (this *MssqlDbInfo) QueryStruct(sql string, pars ...interface{}) (err error) {
	rows, err := this.SqlDataDb.Query(sql)
	if err != nil {
		rows, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(pars...)
	errors_.CheckCommonErr(err)
	return
}

func (this *MssqlDbInfo) QueryIdMap(sql string) (out map[string]string, err error) {
	out = make(map[string]string, 0)
	// 查询数据
	var key, val string
	row, err := this.SqlDataDb.Query(sql)
	if err != nil {
		row, err = this.SqlDataDb.Query(sql)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
	}
	defer row.Close()
	errors_.CheckCommonErr(err)
	for row.Next() {
		err = row.Scan(&key, &val)
		errors_.CheckCommonErr(err)
		out[key] = val
	}
	return
}
