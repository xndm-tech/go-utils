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

func (this *MysqlDbInfo) QueryIdList(sql string) (ids []string, err error) {
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

func (this *MysqlDbInfo) QueryIdIntList(sql string) (ids []int, err error) {
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

func (this *MysqlDbInfo) QueryIdListLen(sql string, len int) (ids []string, err error) {
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

func (this *MysqlDbInfo) QueryStruct(sql string, pars ...interface{}) (err error) {
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

func (this *MysqlDbInfo) QueryIdMap(sql string) (result map[string]string, err error) {
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

// 根据sql得到col: value的map的切片
func (this *MysqlDbInfo) QueryMap(sql string) (result []map[string]string) {
	result = make([]map[string]string, 0)
	var err error
	rows, err := this.SqlDataDb.Query(sql)
	if nil != err {
		errors_.CheckCommonErr(err)
		return nil
	}
	defer rows.Close()
	cols, err := rows.Columns() //  [列名]
	if nil != err {
		errors_.CheckCommonErr(err)
		return nil
	}
	// columns储存所有的列的值, columnPointers的每个元素为columns相应元素的指针
	columns := make([]string, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i, _ := range columns {
		columnPointers[i] = &columns[i]
	}
	for rows.Next() {
		// 把每条记录的信息读到columnPointers里
		if err := rows.Scan(columnPointers...); err != nil {
			errors_.CheckCommonErr(err)
			return nil
		}
		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]string, 0) // m = {colName: value}
		for i, colName := range cols {
			val := columnPointers[i].(*string)
			m[colName] = *val
		}
		result = append(result, m)
	}
	return
}
