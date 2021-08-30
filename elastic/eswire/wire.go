//+build wireinject

package eswire

import (
	"apiserver/elastic/esdao"
	"apiserver/elastic/esservice"
	"apiserver/pkg/elasticsearch"
	"github.com/google/wire"
)

func InitUserService() *esservice.UserService {
	wire.Build(elasticsearch.NewEsClient, esdao.NewUserES, esservice.NewUserService)
	return &esservice.UserService{}
}
