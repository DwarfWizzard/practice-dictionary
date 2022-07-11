package main

import (
	"log"

	dictionary "github.com/DwarfWizzard/practice-dictionary"
	"github.com/DwarfWizzard/practice-dictionary/internal/service"
	"github.com/DwarfWizzard/practice-dictionary/internal/transport/rest"
	"github.com/DwarfWizzard/practice-dictionary/internal/storage/sqlite"
)

func main() {
	srv := new(dictionary.Server)

	db, err := sqlite.NewSQLiteConn("./db/dictionary.db")
	if err != nil {
		log.Fatal(err)
	}

	storage := sqlite.NewStorage(db)
	service := service.NewService(&service.ServiceConfig{
		DictStorage: storage.Dict,
	})
	transport := rest.NewHandler(&rest.HandlerConfig{
		DictService: service.DictService,
	})

	if err := srv.Run(":8000", transport.InitRoutes); err != nil {
		log.Fatal(err.Error())
	}


}
