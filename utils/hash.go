package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"fs1n.anything.bloomfilter/consts"
)

var (
	hashInstanceMD5    hash.Hash
	hashInstanceSHA1   hash.Hash
	hashInstanceSHA256 hash.Hash
	hashInstanceSHA512 hash.Hash
)

func InitHashInstance() {
	hashInstanceMD5 = md5.New()
	hashInstanceSHA1 = sha1.New()
	hashInstanceSHA256 = sha256.New()
	hashInstanceSHA512 = sha512.New()
}

func DoHash(method string, value []byte) (uint, error) {
	// TODO Add hash cache.
	var res uint
	var err error = nil
	switch method {
	case consts.MD5:
		res = MD5Hash(value)
	case consts.SHA1:
		res = SHA1Hash(value)
	case consts.SHA256:
		res = SHA256Hash(value)
	case consts.SHA512:
		res = SHA512Hash(value)
	default:
		err = fmt.Errorf("[doHash] %v is error method", method)
	}
	return res, err
}

func MD5Hash(value []byte) uint {
	hashInstanceMD5.Reset()
	hashInstanceMD5.Write(value)
	return byteToUint(hashInstanceMD5.Sum(nil))
}

func SHA1Hash(value []byte) uint {
	hashInstanceSHA1.Reset()
	hashInstanceSHA1.Write(value)
	return byteToUint(hashInstanceSHA1.Sum(nil))
}

func SHA256Hash(value []byte) uint {
	hashInstanceSHA256.Reset()
	hashInstanceSHA256.Write(value)
	return byteToUint(hashInstanceSHA256.Sum(nil))
}

func SHA512Hash(value []byte) uint {
	hashInstanceSHA512.Reset()
	hashInstanceSHA512.Write(value)
	return byteToUint(hashInstanceSHA512.Sum(nil))
}

func byteToUint(value []byte) uint {
	var res uint
	for i := 0; i < 4 && i < len(value); i++ {
		res |= uint(value[i])
		res <<= 8
	}
	return res
}
