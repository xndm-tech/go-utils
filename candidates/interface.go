package candidates

import (
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/mysqls"
)

// LRUCache is the interface for simple LRU cache.
type CandidateMethod interface {
	GenCandidateFromdb(db *mysqls.MysqlDbInfo, sql string)

	GenCandidateFromList(l []string)

	// 获取备选集元素对应位置
	GetElementInd(s string) int

	GetCids() []string

	GetInd() map[string]int

	GetLen() int

	IsInSlice(item string) bool

	GetDifference(s []string) []string

	GetSliceNoLoop(size, num int) ([]string, error)

	GetSliceLoop(size, num int) ([]string, error)
}

func getCandidateIds(db *mysqls.MysqlDbInfo, sql string) (ids []string) {
	rows, err := db.SqlDataDb.Query(sql)
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

func getIndexId(ids []string) map[string]int {
	IdInd := make(map[string]int, len(ids))
	for ind, id := range ids {
		IdInd[id] = ind
	}
	return IdInd
}
