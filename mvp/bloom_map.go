package mvp

import (
	"errors"
)

type BloomMap struct {
	Values  map[string]BitMap
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
	errInvalidHashMethod = errors.New("[GetBloomMap] invalid hash method name")
	errWarnHashMethod    = errors.New("[HaveMethod] BlooMap didnt set this method")
)

func InitHashMethod() {
	HashMap[MD5] = true
	HashMap[SHA128] = true
	HashMap[SHA256] = true
}

func NewBloomMap(length uint, hashMethods ...string) (*BloomMap, error) {
	res := &BloomMap{}
	var err error
	for _, hashMethod := range hashMethods {
		if _, ok := HashMap[hashMethod]; ok {
			res.Methods[hashMethod] = true
			res.Values[hashMethod], err = NewBitMap(length)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, errInvalidHashMethod
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

func (b *BloomMap) Put(value uint) error {
	for method := range b.Methods {
		b.setHashValue(method, doHash(method, value))
	}
	return nil
}

func (b *BloomMap) NotExist(value uint) bool {
	for method := range b.Methods {
		if b.notExist(method, value) {
			return true
		}
	}
	return false
}

func (b *BloomMap) notExist(method string, value uint) bool {
	hashValue := doHash(method, value)
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

/**
 * Value -> HashValue -> Int64Value -> BloomFilter[Int64Value/8] | 1 << Int64Value%8
 */
