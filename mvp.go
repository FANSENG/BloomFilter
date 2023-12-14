package main

import (
	"fmt"

	"fs1n.anything.bloomfilter/mvp"
)

func init() {
	mvp.InitHashMethod()
}

func TestBloom() {
	bloomMap, err := mvp.NewBloomMap(1024, mvp.MD5)
	if err != nil {
		panic(err)
	}
	bloomMap.Put([]byte("123"))
	bloomMap.Put([]byte("456"))
	bloomMap.Put([]byte("789"))
	fmt.Printf("Did bloomMap not exist %v?: %v\n", 123, bloomMap.NotExist([]byte("123")))
	fmt.Printf("Did bloomMap not exist %v?: %v\n", 345, bloomMap.NotExist([]byte("345")))
	fmt.Printf("Did bloomMap not exist %v?: %v\n", 789, bloomMap.NotExist([]byte("789")))
	fmt.Printf("Did bloomMap not exist %v?: %v\n", 12434234, bloomMap.NotExist([]byte("234r879fsdhckcsbkxcz")))
	// fmt.Println(bloomMap.ToString())
}

func main() {
	TestBloom()
}
