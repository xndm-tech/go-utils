package bitmap

import (
	"github.com/jmoiron/sqlx"
	"github.com/xndm-recommend/go-utils/common/consts"
	"github.com/xndm-recommend/go-utils/config"
	"github.com/xndm-recommend/go-utils/tools/maths"
)

type MysqlMethod interface {
	GetDbConnFromConf(c *config.ConfigEngine, name string)
	SelectStrList(sql string, para ...interface{}) (dest []string, err error)
	SelectIntList(sql string, para ...interface{}) (dest []int, err error)
	QueryStruct(sql string, dest ...interface{}) (err error)
	QueryIdMap(sql string, para ...interface{}) (dest map[string]string, err error)
	GetDb() *sqlx.DB
	GetTableName(key string) string
}

type BitMap struct {
	sqlDataDb *sqlx.DB
	tableName map[string]string
	maxConns  int
	dbTimeOut int
}

/**
 * num/8得到byte[]的index
 * @param num
 * @return
 */
func getIndex(num int) int {
	return num >> 3
}

/**
 * num%8得到在byte[index]的位置
 * @param num
 * @return
 */
func getPosition(num int) int {
	return num & 0x07
}

/**
 * 标记指定数字（num）在bitmap中的值，标记其已经出现过<br/>
 * 将1左移position后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了
 * @param bits
 * @param num
 */
func add(bits []byte, num int) {
	bits[getIndex(num)] |= 1 << getPosition(num)
}

/**
 * 创建bitmap数组
 */
func create(n int) []byte {
	var bits = make([]byte, getIndex(n)+1)
	for i := 0; i < n; i++ {
		add(bits, i)
	}
	return bits
}

//func (this *MysqlDbInfo) GetTableName(key string) string {
//	if val, ok := this.tableName[key]; ok {
//		return val
//	} else {
//		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("key %s not in tablenames.", key)))
//		return consts.BLANK
//	}
//}

// 自去重长度大于35的时候使用
func bitMapDifferenceSelfInt32(s []int32, bitList []byte, l int) []int32 {
	output := make([]int32, 0)
	for _, i := range s {
		index := i >> 3
		pos := i & 0x07
		if bitList[index]&(1<<pos) == 0 {
			bitList[index] |= 1 << pos
			output = append(output, i)
		}
	}

	for _, i := range s {
		index := i >> 3
		bitList[index] = 0
	}

	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}

// 自去重长度大于35的时候使用
func bitMapDifferenceSelfInt32XXXXXXXX(s []int, l int) []int {
	var bitList = make([]byte, (len(s)>>3)+1)
	output := make([]int, 0)
	for _, i := range s {
		index := i >> 3
		pos := i & 0x07
		if bitList[index]&(1<<pos) == 0 {
			output = append(output, i)
		}
	}

	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
	return output[:size:size]
}
