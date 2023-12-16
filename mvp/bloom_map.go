package mvp

import (
	"errors"
	"fmt"

	"fs1n.anything.bloomfilter/consts"
)

/**
 * Value -> HashValue -> Int64Value -> BloomFilter[Int64Value/8] | 1 << Int64Value%8
 */

type BloomMap struct {
	Values  map[string]*BitMap
	Methods map[string]interface{}
	Length  uint
}

var (
	HashMap map[string]interface{}
)

var (
	errWarnHashMethod = errors.New("[HaveMethod] BlooMap didnt set this method")
	errSetHashValue   = errors.New("[Put] set hash err")
)

func InitHashMethod() {
	HashMap = make(map[string]interface{})
	HashMap[consts.MD5] = true
	HashMap[consts.SHA1] = true
	HashMap[consts.SHA256] = true
	HashMap[consts.SHA512] = true
}

func BuildDefaultBloomMap() *BloomMap {
	res := &BloomMap{}
	res.Length = 0
	res.Methods = make(map[string]interface{})
	res.Values = make(map[string]*BitMap)
	return res
}

func NewBloomMap(length uint, hashMethods ...string) (*BloomMap, error) {
	if length == 0 {
		panic("[NewBloomMap] length cannot be zero")
	}
	res := BuildDefaultBloomMap()
	res.Length = length
	var err error
	for _, hashMethod := range hashMethods {
		if _, ok := HashMap[hashMethod]; ok {
			res.Methods[hashMethod] = true
			res.Values[hashMethod], err = NewBitMap(length)
			if err != nil {
				// TODO: Access the log module.
				fmt.Println("[NewBloomMap] Create BloomMap Error")
				return nil, err
			}

		} else {
			panic("[NewBloomMap] invalid hash method name")
		}
	}
	return res, nil
}

func (b *BloomMap) HaveMethod(hashMethod string) bool {
	if _, ok := b.Methods[hashMethod]; ok {
		return true
	}
	return false
}

func (b *BloomMap) Put(value []byte) error {
	for method := range b.Methods {
		hashValue, err := doHash(method, value)
		if err != nil {
			return err
		}
		err = b.setHashValue(method, hashValue%b.Length)
		if err != nil {
			fmt.Println(errSetHashValue, err)
		}
	}
	return nil
}

func (b *BloomMap) NotExist(value []byte) bool {
	for method := range b.Methods {
		if b.notExist(method, value) {
			return true
		}
	}
	return false
}

func (b *BloomMap) notExist(method string, value []byte) bool {
	hashValue, err := doHash(method, value)
	hashValue = hashValue % b.Length
	if err != nil {
		return true
	}
	return b.Values[method].NotExist(hashValue)
}

func (b *BloomMap) setHashValue(hashMethod string, value uint) error {
	if !b.HaveMethod(hashMethod) {
		return errWarnHashMethod
	}
	if err := b.Values[hashMethod].Set(value); err != nil {
		return err
	}
	return nil
}

func (b *BloomMap) ToString() map[string][]byte {
	res := make(map[string][]byte)
	for key, value := range b.Values {
		res[key] = value.bitMap
	}
	return res
}
