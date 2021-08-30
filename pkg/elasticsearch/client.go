package elasticsearch

import (
	log2 "github.com/chenglongcl/log"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewEsClient() *elastic.Client {
	address := viper.GetString("elasticSearch.address")
	username := viper.GetString("elasticSearch.username")
	password := viper.GetString("elasticSearch.password")
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(address),
		//账号密码
		elastic.SetBasicAuth(username, password),
		//关闭sniff模式；或者设置es的地址为publish_address 地址
		elastic.SetSniff(false),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		log2.Error("Failed to create elastic client", err)
	}
	return client
}
