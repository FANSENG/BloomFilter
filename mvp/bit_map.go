package mvp

import (
	"fmt"
)

type BitMap struct {
	bitMap []byte
	length int
}

const (
	byteSize = 8
)

// NewBitMap n: 总位数
func NewBitMap(n uint) (*BitMap, error) {
	if n%byteSize != 0 {
		return nil, ErrBitMapLength(n)
	}
	resp := &BitMap{
		bitMap: make([]byte, n/byteSize+1),
	}
	resp.length = len(resp.bitMap)
	return resp, nil
}

func (bt *BitMap) Len() int {
	return bt.length
}

func (bt *BitMap) Set(val uint) error {
	if val > uint(bt.Len()*byteSize) {
		return bt.ErrValueLength()
	}
	bt.bitMap[val/byteSize] |= 1 << (val % byteSize)
	return nil
}

func (bt *BitMap) Del(val uint) bool {
	if val > uint(bt.Len()*byteSize) {
		return false
	}
	bt.bitMap[val/byteSize] &= 0 << (val % byteSize)
	return true
}

func (bt *BitMap) NotExist(val uint) bool {
	if val > uint(bt.Len()*byteSize) {
		return false
	}
	return bt.bitMap[val/byteSize]&(1<<(val%byteSize)) == 0
}

func (bt *BitMap) ErrValueLength() error {
	return fmt.Errorf("[Set] Value bigger than %v(Length) * %v(Bytesize)", bt.Len(), byteSize)
}

func ErrBitMapLength(len uint) error {
	return fmt.Errorf("[NewBitMap] %v(Length) cannot be divided by %v(Bytesize)", len, byteSize)
}
