package bitmap

import "github.com/xndm-recommend/go-utils/dbs/rediss"

type RedisBitMap struct {
	RedisCli rediss.RedisDbInfo
	RedisItem rediss.ItemInfo
	Bits     []byte
	Max      int
}



func (r *RedisBitMap)multiSetBit( name string,value bool , offsets ...int64) {

	r.RedisItem.ItemSetByte()

template.executePipelined((RedisCallback<Object>) connection -> {
for (long offset : offsets) {
connection.setBit(name.getBytes(), offset, value);
}
return null;
});
}

//public List<Boolean> multiGetBit(String name, long... offsets) {
//List<Object> results = template.executePipelined((RedisCallback<Object>) connection -> {
//for (long offset : offsets) {
//connection.getBit(name.getBytes(), offset);
//}
//return null;
//});
//List<Boolean> list = new ArrayList<>();
//results.forEach(obj -> {
//list.add((Boolean) obj);
//});
//return list;
//}