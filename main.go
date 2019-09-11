package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"gin-demo/common/db"
	"gin-demo/common/log"
	"gin-demo/common/config"
)

func main() {
	config.Init()
	log.Init()
	db.DB.Init()

	defer db.DB.Close()

	gin.SetMode(viper.GetString("runmode"))

	ctx := gin.New()

	ctx.Run(viper.GetString("addr"))
}
