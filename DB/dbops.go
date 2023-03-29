package DB

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

//leveldb的key和value都是byte数组类型
var DB *leveldb.DB

func init() {
	DB, _ = leveldb.OpenFile("./apis.db", nil)
}

//根据key获取value
func GetValueByKey(key string) string {
	data, err := DB.Get([]byte(key), nil)
	if err != nil {
		panic(err)
	}
	return string(data)
}

//设置key和value
func SetKeyAndValue(key string, value []byte) bool {
	err := DB.Put([]byte(key), value, nil)
	fmt.Println(key, value)
	return err == nil
}

//根据key删除key value
func DelKeyAndValue(key string) bool {
	_, err := DB.Get([]byte(key), nil)
	if err == leveldb.ErrNotFound {
		return false
	}
	err2 := DB.Delete([]byte(key), nil)
	return err2 == nil
}

//list出所有的api "key",好吧，其实它是value，但也是key，此key非彼key
func ListAllApi() map[string]string {
	iter := DB.NewIterator(nil, nil)
	m := make(map[string]string)
	for iter.Next() {
		m[string(iter.Key())] = string(iter.Value())
	}
	return m
}
