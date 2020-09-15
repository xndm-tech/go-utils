package bitmap

import (
	"testing"

	"github.com/xndm-recommend/go-utils/tools/maths"
)

/**
 * num/8得到byte[]的index
 * @param num
 * @return
 */
func TestBitMapDifferenceSelfInt32XXXXXXXX(t *testing.T) {
	a := maths.GenNRandInt(3000, 100)
	bitMapDifferenceSelfInt32XXXXXXXX(a, 10)
}

func BenchmarkBitMapDifferenceSelfInt32XXXXXXXX(b *testing.B) {
	b.ResetTimer()
	a := maths.GenNRandInt(3000, 100)
	//b.Log(a)
	for i := 0; i < b.N; i++ {
		bitMapDifferenceSelfInt32XXXXXXXX(a, 10)
	}
}
