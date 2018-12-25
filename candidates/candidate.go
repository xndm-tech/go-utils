package candidates

import (
	"github.com/xndm-recommend/go-utils/errors"
	"github.com/xndm-recommend/go-utils/mysqls"
)

const ERR_INDEX = -1

type Candidate struct {
	Ids   []string
	Index map[string]int
}

func (this *Candidate) GetCandidateIndex(s string) int {
	if ind, ok := this.Index[s]; ok {
		return ind
	} else {
		return ERR_INDEX
	}
}

func (this *Candidate) GetCandidateId() []string {
	return this.Ids
}

func getCandidateIds(db *mysqls.MysqlDbInfo, sql string) (ids []string) {
	rows, err := db.SqlDataDb.Query(sql)
	if err != nil {
		errors.CheckCommonErr(err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var tmpId string
		err := rows.Scan(&tmpId)
		errors.CheckCommonErr(err)
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

func GenCandidateList(db *mysqls.MysqlDbInfo, sql string) *Candidate {
	var can *Candidate
	can.Ids = getCandidateIds(db, sql)
	can.Index = getIndexId(can.Ids)
	return can
}
