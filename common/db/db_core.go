package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"gin-demo/common/log"
)

type Database struct {
	Local *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Local: GetLocalDB(),
	}
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local",
	)

	db, err := gorm.Open("mysql", config)

	setupDB(db)

	if err != nil {
		log.Log.Default.Error(
			"Database open failed. ",
			zap.String("err", err.Error()),
		)
	}

	return db
}

func GetLocalDB() *gorm.DB {
	return InitLocalDB()
}

func InitLocalDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(viper.GetInt("max_open_connections"))
	db.DB().SetMaxIdleConns(viper.GetInt("max_idle_connections"))
	db.DB().SetConnMaxLifetime(viper.GetDuration("connection_max_life_time"))

	err := db.DB().Ping()

	if err != nil {
		log.Log.Default.Error(
			"Database setup failed. ",
			zap.String("err", err.Error()),
		)
	}
}

func (db *Database) Close() {
	DB.Local.Close()
}
