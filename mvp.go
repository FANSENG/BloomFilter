package main

import (
	bm "fs1n.anything.bloomfilter/bloom_map"
	"fs1n.anything.bloomfilter/utils"
)

func init() {
	bm.InitHashMethod()
	utils.InitHashInstance()
}

func main() {
}
