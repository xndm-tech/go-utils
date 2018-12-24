package candidates

import (
	"github.com/zhanglanhui/go-utils/utils/err_utils"
	"github.com/zhanglanhui/go-utils/utils/mysql_utils"
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

func getCandidateIds(db *mysql_utils.MysqlDbInfo, sql string) (ids []string) {
	rows, err := db.SqlDataDb.Query(sql)
	if err != nil {
		err_utils.CheckCommonErr(err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var tmpId string
		err := rows.Scan(&tmpId)
		err_utils.CheckCommonErr(err)
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

func GenCandidateList(db *mysql_utils.MysqlDbInfo, sql string) *Candidate {
	var can *Candidate
	can.Ids = getCandidateIds(db, sql)
	can.Index = getIndexId(can.Ids)
	return can
}
