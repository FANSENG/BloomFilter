package mvp

import (
	"fmt"
	"hash/fnv"
)

func doHash(method string, value ...[]byte) uint {
	res := MD5Hash(value...)
	fmt.Printf("value is %v, Hash is %v\n", value, res)
	return res
}

func MD5Hash(value ...[]byte) uint {
	h := fnv.New32()
	for _, b := range value {
		h.Write(b)
	}
	return uint(h.Sum32())
}
