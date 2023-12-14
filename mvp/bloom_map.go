package mvp

import (
	"errors"
	"fmt"
)

type BloomMap struct {
	Values  map[string]*BitMap
	Methods map[string]interface{}
	Length  uint
}

const (
	MD5    string = "md5"
	SHA128 string = "sha128"
	SHA256 string = "sha256"
)

var (
	HashMap map[string]interface{}
)

var (
	errWarnHashMethod = errors.New("[HaveMethod] BlooMap didnt set this method")
	errSetHashValue   = errors.New("[Put] set Hash Err")
)

func InitHashMethod() {
	HashMap = make(map[string]interface{})
	HashMap[MD5] = true
	HashMap[SHA128] = true
	HashMap[SHA256] = true
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
		err := b.setHashValue(method, doHash(method, value))
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
	hashValue := doHash(method, value)
	// ! And Here
	return b.Values[method].NotExist(hashValue % b.Length)
}

func (b *BloomMap) setHashValue(hashMethod string, value uint) error {
	if !b.HaveMethod(hashMethod) {
		return errWarnHashMethod
	}
	// ! Here
	if err := b.Values[hashMethod].Set(value % b.Length); err != nil {
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

/**
 * Value -> HashValue -> Int64Value -> BloomFilter[Int64Value/8] | 1 << Int64Value%8
 */
