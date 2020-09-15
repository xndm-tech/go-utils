package bitmap

//// 长度大于30使用
//func bitMapFilterInt32(s, filter []int32, bitArr gobal.BitMap, l int) []int32 {
//	for _, i := range filter {
//		index := i >> 3
//		pos := i & 0x07
//		bitArr.Bits[index] ^= 1 << pos
//	}
//	output := make([]int32, 0)
//
//	for _, i := range s {
//		index := i >> 3
//		pos := i & 0x07
//		if bitArr.Bits[index]&(1<<pos) != 0 {
//			output = append(output, i)
//		}
//	}
//	for _, i := range filter {
//		index := i >> 3
//		pos := i & 0x07
//		bitArr.Bits[index] ^= 1 << pos
//	}
//	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(output))
//	return output[:size:size]
//}
