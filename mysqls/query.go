package mysqls

import (
	sql_ "database/sql"

	"github.com/xndm-recommend/go-utils/errors_"
)

//http://jmoiron.github.io/sqlx/
func (this *MysqlDbInfo) QueryIdList(sql string, para ...interface{}) (ids []string, err error) {
	err = this.SqlDataDb.Select(&ids, sql, para...)
	if err != nil {
		err = this.SqlDataDb.Select(&ids, sql, para...)
		if err != nil {
			errors_.CheckCommonErr(err)
		}
		return
	}
	return
}

func (this *MysqlDbInfo) QueryIdIntList(sql string, para ...interface{}) (ids []int, err error) {
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

func (this *MysqlDbInfo) QueryStruct(sql string, dest ...interface{}) (err error) {
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

func (this *MysqlDbInfo) QueryIdMap(sql string) (result map[string]string, err error) {
	result = make(map[string]string, 0)
	// 查询数据
	var key, val sql_.NullString
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
		result[key.String] = val.String
	}
	return
}
