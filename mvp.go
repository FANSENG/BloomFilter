package main

import (
	"fmt"

	"fs1n.anything.bloomfilter/mvp"
)

func init() {
	mvp.InitHashMethod()
}

func TestBloom() {
	bloomMap, err := mvp.NewBloomMap(1024*1024, mvp.MD5, mvp.SHA128, mvp.SHA256)
	if err != nil {
		panic(err)
	}
	bloomMap.Put(123)
	bloomMap.Put(456)
	bloomMap.Put(789)
	fmt.Printf("Did bloomMap not exist %v?: %v", 123, bloomMap.NotExist(123))
	fmt.Printf("Did bloomMap not exist %v?: %v", 345, bloomMap.NotExist(345))
	fmt.Printf("Did bloomMap not exist %v?: %v", 12434234, bloomMap.NotExist(12434234))
}

func main() {
	TestBloom()
}
