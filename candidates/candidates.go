package candidates

import (
	"github.com/pkg/errors"
	"github.com/xndm-recommend/go-utils/errors_"
	"github.com/xndm-recommend/go-utils/maths"
	"github.com/xndm-recommend/go-utils/mysqls"
	"github.com/xndm-recommend/go-utils/tools"
)

const (
	ERR_INDEX = -1
)

type Candidate struct {
	Cids []string
	Len  int
	Ind  map[string]int
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

func (this *Candidate) GenCandidateFromdb(db *mysqls.MysqlDbInfo, sql string) {
	this.Cids = getCandidateIds(db, sql)
	this.Ind = getIndexId(this.Cids)
	this.Len = len(this.Cids)
}

func (this *Candidate) GenCandidateFromList(s []string) {
	this.Cids = s
	this.Ind = getIndexId(this.Cids)
	this.Len = len(this.Cids)
}

func (this *Candidate) GetElementInd(s string) int {
	if ind, ok := this.Ind[s]; ok {
		return ind
	} else {
		return ERR_INDEX
	}
}

func (this *Candidate) GetCids() []string {
	return this.Cids
}

func (this *Candidate) GetInd() map[string]int {
	return this.Ind
}

func (this *Candidate) GetLen() int {
	return this.Len
}

func (this *Candidate) IsInSlice(item string) bool {
	if _, ok := this.Ind[item]; ok {
		return true
	} else {
		return false
	}
}

func (this *Candidate) GetDifference(s []string) []string {
	return tools.DifferenceStr(this.Cids, s)
}

func (this *Candidate) GetDifferenceLen(s []string, len int) []string {
	return tools.DifferenceStrLen(this.Cids, s, len)
}

func (this *Candidate) GetSliceNoLoop(size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, errors.New("Input paras error")
	}
	return this.Cids[size*(num-1) : maths.MinInt(num*size, this.Len)], nil
}

func (this *Candidate) GetSliceLoop(size, num int) ([]string, error) {
	if num <= 0 || size <= 0 {
		return []string{}, errors.New("Input paras error")
	}
	start := (size * (num - 1)) % this.Len
	end := (num * size) % this.Len
	return this.Cids[start:end], nil
}
