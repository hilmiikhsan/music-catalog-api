package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/configs"
	"github.com/hilmiikhsan/music-catalog/pkg/internal_sql"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("failed to initialize config: ", err)
	}

	cfg = configs.Get()

	db, err := internal_sql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	router := gin.Default()

	router.Run(cfg.Service.Port)
}
