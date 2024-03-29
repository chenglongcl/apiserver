// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package eswire

import (
	"apiserver/elastic/esdao"
	"apiserver/elastic/esservice"
	"apiserver/pkg/elasticsearch"
)

// Injectors from wire.go:

func InitUserService() *esservice.UserService {
	client := elasticsearch.NewEsClient()
	userES := esdao.NewUserES(client)
	userService := esservice.NewUserService(userES)
	return userService
}
