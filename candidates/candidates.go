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
	Cids []interface{}
	Len  int
	Ind  map[interface{}]int
}

func getCandidateIds(db *mysqls.MysqlDbInfo, sql string) (ids []interface{}) {
	rows, err := db.SqlDataDb.Query(sql)
	if err != nil {
		errors_.CheckCommonErr(err)
		return ids
	}
	defer rows.Close()
	for rows.Next() {
		var tmpId interface{}
		err := rows.Scan(&tmpId)
		errors_.CheckCommonErr(err)
		if nil == err {
			ids = append(ids, tmpId)
		}
	}
	return ids
}

func getIndexId(ids []interface{}) map[interface{}]int {
	idInd := make(map[interface{}]int, len(ids))
	for ind, id := range ids {
		idInd[id] = ind
	}
	return idInd
}

func (this *Candidate) GenCandidateFromdb(db *mysqls.MysqlDbInfo, sql string) {
	this.Cids = getCandidateIds(db, sql)
	this.Ind = getIndexId(this.Cids)
	this.Len = len(this.Cids)
}

func (this *Candidate) GenCandidateFromList(s []interface{}) {
	this.Cids = s
	this.Ind = getIndexId(this.Cids)
	this.Len = len(this.Cids)
}

func (this *Candidate) GetElementInd(s interface{}) int {
	if ind, ok := this.Ind[s]; ok {
		return ind
	} else {
		return ERR_INDEX
	}
}

func (this *Candidate) GetCids() []interface{} {
	return this.Cids
}

func (this *Candidate) GetInd() map[interface{}]int {
	return this.Ind
}

func (this *Candidate) GetLen() int {
	return this.Len
}

func (this *Candidate) IsInSlice(item interface{}) bool {
	if _, ok := this.Ind[item]; ok {
		return true
	} else {
		return false
	}
}

func (this *Candidate) GetDifference(s []interface{}) []interface{} {
	return tools.DiffInterface(this.Cids, s)
}

func (this *Candidate) GetDifferenceLen(s []interface{}, len int) []interface{} {
	return tools.DiffInterfaceLen(this.Cids, s, len)
}

func (this *Candidate) GetSliceNoLoop(size, num int) ([]interface{}, error) {
	if num <= 0 || size <= 0 {
		return []interface{}{}, errors.New("Input paras error")
	}
	return this.Cids[maths.MinInt(size*(num-1), this.Len):maths.MinInt(num*size, this.Len)], nil
}

func (this *Candidate) GetSliceLoop(size, num int) ([]interface{}, error) {
	if num <= 0 || size <= 0 {
		return []interface{}{}, errors.New("Input parameter error!!!")
	}
	start := (size * (num - 1)) % this.Len
	end := (num * size) % this.Len
	return this.Cids[start:end], nil
}
