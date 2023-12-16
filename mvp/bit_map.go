package mvp

import (
	"fmt"

	"fs1n.anything.bloomfilter/consts"
)

type BitMap struct {
	bitMap    []byte
	length    uint
	maxNumber uint
}

// NewBitMap n: 总位数
func NewBitMap(length uint) (*BitMap, error) {
	if length%consts.ByteSize != 0 {
		return nil, ErrBitMapLength(length)
	}
	resp := &BitMap{
		bitMap: make([]byte, length/consts.ByteSize),
	}
	resp.length = uint(len(resp.bitMap))
	resp.maxNumber = uint(resp.length) * consts.ByteSize
	return resp, nil
}

func (bt *BitMap) Len() uint {
	return bt.length
}

func (bt *BitMap) Set(val uint) error {
	if val > bt.maxNumber {
		return bt.ErrValueLength()
	}
	bt.bitMap[val/consts.ByteSize] |= 1 << (val % consts.ByteSize)
	return nil
}

// Delete operations should not be used in basic Bloom filters.
func (bt *BitMap) Delete(val uint) bool {
	if val > bt.maxNumber {
		return false
	}
	bt.bitMap[val/consts.ByteSize] &= 0 << (val % consts.ByteSize)
	return true
}

func (bt *BitMap) NotExist(val uint) bool {
	if val > bt.maxNumber {
		return false
	}
	return bt.bitMap[val/consts.ByteSize]&(1<<(val%consts.ByteSize)) == 0
}

func (bt *BitMap) ErrValueLength() error {
	return fmt.Errorf("[Set] Value bigger than %v(Length) * %v(Bytesize)", bt.Len(), consts.ByteSize)
}

func ErrBitMapLength(len uint) error {
	return fmt.Errorf("[NewBitMap] %v(Length) cannot be divided by %v(Bytesize)", len, consts.ByteSize)
}
