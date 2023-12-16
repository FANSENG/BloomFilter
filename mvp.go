package main

import (
	"fmt"

	"fs1n.anything.bloomfilter/consts"
	"fs1n.anything.bloomfilter/mvp"
)

func init() {
	mvp.InitHashMethod()
	mvp.InitHashInstance()
}

func TestBloom() {
	bloomMap, err := mvp.NewBloomMap(1024,
		consts.MD5,
		consts.SHA1,
		consts.SHA256,
		consts.SHA512,
	)
	if err != nil {
		panic(err)
	}
	bloomMap.Put([]byte("123"))
	bloomMap.Put([]byte("456"))
	bloomMap.Put([]byte("789"))
	fmt.Printf("Did bloomMap exist %v?: %v\n", 123, !bloomMap.NotExist([]byte("123")))
	fmt.Printf("Did bloomMap exist %v?: %v\n", 345, !bloomMap.NotExist([]byte("345")))
	fmt.Printf("Did bloomMap exist %v?: %v\n", 456, !bloomMap.NotExist([]byte("456")))
	fmt.Printf("Did bloomMap exist %v?: %v\n", 789, !bloomMap.NotExist([]byte("789")))
	fmt.Printf("Did bloomMap exist %v?: %v\n", "234r879fsdhckcsbkxcz", !bloomMap.NotExist([]byte("234r879fsdhckcsbkxcz")))
}

func main() {
	TestBloom()
}
