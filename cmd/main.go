package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/configs"
	membershipsHandler "github.com/hilmiikhsan/music-catalog/internal/handler/memberships"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/hilmiikhsan/music-catalog/internal/repository/memberships"
	membershipsSvc "github.com/hilmiikhsan/music-catalog/internal/service/memberships"
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

	db.AutoMigrate(&memberships.User{})

	router := gin.Default()

	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)

	membershipsHandler := membershipsHandler.NewHandler(router, membershipsSvc)

	membershipsHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
