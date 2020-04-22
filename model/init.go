package model

import (
	"fmt"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"time"

	// MySql driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var database *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(150) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(10) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetConnMaxLifetime(7200 * time.Second)
}

func initDB(name string) *gorm.DB {
	if name == "" {
		return nil
	}
	return openDB(
		viper.GetString(name+".username"),
		viper.GetString(name+".password"),
		viper.GetString(name+".addr"),
		viper.GetString(name+".name"),
	)
}

func Init() {
	database = &Database{
		Self:   initDB("db"),
		Docker: initDB("docker_db"),
	}
}

func Close() {
	database.Self.Close()
	database.Docker.Close()
}

func SelectDB(name string) *gorm.DB {
	switch name {
	case "self":
		return database.Self
	}
	return nil
}
