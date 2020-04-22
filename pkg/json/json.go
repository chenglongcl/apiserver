package json

import (
	jsoniter "github.com/json-iterator/go"
	"time"
	"unsafe"
)

func Init() {
	//自定义time.Time类型JSON格式 json中时间字段字符串格式"2019-09-09 10:10:10"
	jsoniter.RegisterTypeDecoderFunc("time.Time", func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", iter.ReadString(), time.Local)
		if err != nil {
			iter.Error = err
			return
		}
		*((*time.Time)(ptr)) = t
	})
}
