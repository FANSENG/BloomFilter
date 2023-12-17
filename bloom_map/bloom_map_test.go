package bloom_map

import (
	"errors"
	"testing"

	"fs1n.anything.bloomfilter/base"
	"fs1n.anything.bloomfilter/consts"
	"fs1n.anything.bloomfilter/utils"
	"github.com/bytedance/mockey"
	"github.com/smartystreets/goconvey/convey"
)

func init() {
	InitHashMethod()
	utils.InitHashInstance()
}

func TestBloomMapNotExistValue(t *testing.T) {
	mockey.PatchConvey("TestBloomMapNotExistValue", t, func() {
		bloomMap, err := NewBloomMap(1024,
			consts.MD5,
			consts.SHA1,
			consts.SHA256,
			consts.SHA512,
		)
		if err != nil {
			panic(err)
		}
		bloomMap.Put([]byte("123"))
		bloomMap.Put([]byte("456"))
		bloomMap.Put([]byte("789"))
		convey.So(bloomMap.NotExist([]byte("123")), convey.ShouldEqual, false)
		convey.So(bloomMap.NotExist([]byte("456")), convey.ShouldEqual, false)
		convey.So(bloomMap.NotExist([]byte("789")), convey.ShouldEqual, false)
		convey.So(bloomMap.NotExist([]byte("357")), convey.ShouldEqual, true)
		convey.So(bloomMap.NotExist([]byte("asviuhiadkj12")), convey.ShouldEqual, true)
	})
}

func TestBloomMapErr(t *testing.T) {
	mockey.PatchConvey("TestBloomMapDoHashErr", t, func() {
		var errDoHash = errors.New("Unit test do hash error!")
		mockey.Mock(utils.DoHash).Return(1, errDoHash).Build()
		bloomMap, err := NewBloomMap(1024,
			consts.MD5,
			consts.SHA1,
			consts.SHA256,
			consts.SHA512,
		)
		if err != nil {
			panic(err)
		}
		err = bloomMap.Put([]byte("123"))
		convey.So(err, convey.ShouldEqual, errDoHash)
		convey.So(bloomMap.NotExist([]byte("123")), convey.ShouldEqual, true)
	})

	mockey.PatchConvey("TestBloomMapSetHashErr", t, func() {
		var errSetHash = errors.New("Unit test set hash error!")
		bloomMap, err := NewBloomMap(1024,
			consts.MD5,
			consts.SHA1,
			consts.SHA256,
			consts.SHA512,
		)
		if err != nil {
			panic(err)
		}
		mockey.Mock((*base.BitMap).Set).Return(errSetHash).Build()
		err = bloomMap.Put([]byte("123"))
		convey.So(err, convey.ShouldEqual, errSetHashValue)
		convey.So(bloomMap.NotExist([]byte("123")), convey.ShouldEqual, true)
	})
}
