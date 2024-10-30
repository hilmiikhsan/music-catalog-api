package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/music-catalog/internal/configs"
	membershipsHandler "github.com/hilmiikhsan/music-catalog/internal/handler/memberships"
	tracksHandler "github.com/hilmiikhsan/music-catalog/internal/handler/tracks"
	"github.com/hilmiikhsan/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/hilmiikhsan/music-catalog/internal/repository/memberships"
	"github.com/hilmiikhsan/music-catalog/internal/repository/spotify"
	membershipsSvc "github.com/hilmiikhsan/music-catalog/internal/service/memberships"
	tracksSvc "github.com/hilmiikhsan/music-catalog/internal/service/tracks"
	"github.com/hilmiikhsan/music-catalog/pkg/httpclient"
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

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)
	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipsSvc := membershipsSvc.NewService(cfg, membershipsRepo)
	tracksSvc := tracksSvc.NewService(spotifyOutbound)

	membershipsHandler := membershipsHandler.NewHandler(router, membershipsSvc)
	membershipsHandler.RegisterRoute()

	tracksHandler := tracksHandler.NewHandler(router, tracksSvc)
	tracksHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
