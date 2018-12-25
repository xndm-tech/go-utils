package candidates

import (
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

	GetDifferenceLen(s []string, len int) []string

	GetSliceNoLoop(size, num int) ([]string, error)

	GetSliceLoop(size, num int) ([]string, error)
}
