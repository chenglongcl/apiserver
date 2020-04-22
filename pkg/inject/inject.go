package inject

import (
	"apiserver/service/bll"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/facebookgo/inject"
	"github.com/lexkong/log"
)

// Object 注入对象
type Object struct {
	Common          *bll.Common
	AliYunOssClient *oss.Client
}

var Obj *Object

// Init 初始化依赖注入
func Init() {
	g := new(inject.Graph)

	//common new
	Common := new(bll.Common)
	_ = g.Provide(&inject.Object{Value: Common})

	if err := g.Populate(); err != nil {
		log.Errorf(err, "初始化依赖注入发生错误")
	}
	Obj = &Object{
		Common: Common,
	}
	return
}
