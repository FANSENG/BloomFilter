package mvp

import (
	"fmt"
)

type BitMap []byte

const (
	byteSize = 8
)

var (
	ErrBitMapLength = fmt.Errorf("[NewBitMap] Length cannot be divided by %v", byteSize)
	ErrValueLength  = fmt.Errorf("[Set] Value bigger than length * %v", byteSize)
)

// NewBitMap n: 总位数
func NewBitMap(n uint) (BitMap, error) {
	if n%byteSize != 0 {
		return nil, ErrBitMapLength
	}
	return make([]byte, n/byteSize+1), nil
}

func (bt BitMap) Set(val uint) error {
	if val > uint(len(bt)*byteSize) {
		return ErrValueLength
	}
	bt[val/byteSize] |= 1 << (val % byteSize)
	return nil
}

func (bt BitMap) Del(val uint) bool {
	if val > uint(len(bt)*byteSize) {
		return false
	}
	bt[val/byteSize] &= 0 << (val % byteSize)
	return true
}

func (bt BitMap) NotExist(val uint) bool {
	if val > uint(len(bt)*byteSize) {
		return false
	}
	return bt[val/byteSize]&(1<<(val%byteSize)) == 0
}
