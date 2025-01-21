package main

import (
	"log"
	api "milliy/api"
	"milliy/api/handler"
	"milliy/config"
	"milliy/logs"
	"milliy/service"
	"milliy/storage/postgres"
	"milliy/upload"

	"github.com/casbin/casbin/v2"
)

func main() {
	conf := config.Load()
	hand := NewHandler()
	router := api.Router(hand)
	log.Printf("server is running...")
	log.Fatal(router.Run(conf.Server.HTTP_PORT))
}

func NewHandler() handler.HandlerInterface {
	//minio
	uploader, err := upload.NewMinioUploader()
	if err != nil {
		log.Fatal(err)
	}

	// conf := config.Load()
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}
	storage := postgres.NewPostgresStorage(db)
	twit := service.NewTwitService(storage)
	user := service.NewUserService(storage)

	logs := logs.NewLogger()
	enforcer, err := casbin.NewEnforcer("casbin/model.conf", "casbin/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	return &handler.Handler{
		User:     user,
		Twit:     twit,
		Log:      logs,
		Enforcer: enforcer,
		MINIO:    uploader,
	}
}
