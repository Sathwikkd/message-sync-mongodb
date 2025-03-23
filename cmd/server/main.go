package main

import (
	"chat-sync-service/internal/app"
	"chat-sync-service/internal/config"
	"chat-sync-service/internal/delivery/http"
	"chat-sync-service/internal/infra"
	"chat-sync-service/internal/infra/mongodb"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig()
	db, err := infra.InitMongoDB(cfg.MongoURI, cfg.DBName)
	if err != nil {
		panic(err)
	}

	repo := mongodb.NewMessageRepository(db)
	syncUC := app.NewSyncUsecase(repo)

	e := echo.New()
	http.NewHandler(e, syncUC)

	e.Logger.Fatal(e.Start(":8080"))
}
