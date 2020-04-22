package oss

import (
	"apiserver/pkg/oss/client"
	"apiserver/pkg/oss/common"
)

func Init() {
	client.InitAliClient()
}

func SelectClient(name string) common.OSSClient {
	switch name {
	case "ali":
		return client.DefaultAliClient()
	}
	return nil
}
