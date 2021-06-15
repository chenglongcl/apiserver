package redisgo

import (
	"github.com/chenglongcl/log"
	"github.com/chenglongcl/redisgo"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
)

var redis *redisgo.Cacher

func Init() error {
	var err error
	redis, err = redisgo.New(
		redisgo.Options{
			Network:   viper.GetString("redis_conf.network"),
			Addr:      viper.GetString("redis_conf.address"),
			Password:  viper.GetString("redis_conf.password"),
			MaxActive: 500,
			MaxIdle:   16,
			Prefix:    viper.GetString("redis_conf.prefix"),
			Marshal:   jsoniter.Marshal,
			Unmarshal: jsoniter.Unmarshal,
		})
	if err != nil {
		log.Errorf(err, "Redis connection failed:%s", viper.GetString("redis_conf.address"))
	}
	return nil
}

func My() *redisgo.Cacher {
	return redis
}
