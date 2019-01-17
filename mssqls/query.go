package mssqls

import (
	"github.com/xndm-recommend/go-utils/errors_"
)

//http://jmoiron.github.io/sqlx/
func (this *MssqlDbInfo) QueryIdList(sql string, para ...interface{}) (ids []string, err error) {
	err = this.SqlDataDb.Select(&ids, sql, para...)
	if err != nil {
		err = this.SqlDataDb.Select(&ids, sql, para...)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
	}
	return
}

func (this *MssqlDbInfo) QueryIdIntList(sql string, para ...interface{}) (ids []int, err error) {
	err = this.SqlDataDb.Select(&ids, sql, para...)
	if err != nil {
		err = this.SqlDataDb.Select(&ids, sql, para...)
		if err != nil {
			errors_.CheckCommonErr(err)
			return
		}
	}
	return
}

func (this *MssqlDbInfo) QueryStruct(sql string, dest ...interface{}) (err error) {
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
	err = rows.Scan(dest...)
	errors_.CheckCommonErr(err)
	return
}

func (this *MssqlDbInfo) QueryIdMap(sql string) (result map[string]string, err error) {
	result = make(map[string]string, 0)
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
		result[key] = val
	}
	return
}
