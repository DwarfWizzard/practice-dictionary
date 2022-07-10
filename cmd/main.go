package main

import (
	"log"

	dictionary "github.com/DwarfWizzard/practice-dictionary"
	"github.com/DwarfWizzard/practice-dictionary/internal/service"
	"github.com/DwarfWizzard/practice-dictionary/internal/storage/sqlite"
	"github.com/DwarfWizzard/practice-dictionary/internal/transport/rest"
)

func main() {
	srv := new(dictionary.Server)

	db, err := sqlite.NewSQLite3("./db/dictionary.db")
	if err != nil {
		log.Fatal(err)
	}

	storage := sqlite.NewStorage(db)
	service := service.NewService(&service.ServiceConfig{
		DictStorage: storage.Dict,
	})
	dictionaryTransport := rest.NewHandler(&rest.HandlerConfig{
		DictService: service.Dict,
	})

	if err := srv.Run(":8000", dictionaryTransport.InitRoutes); err != nil {
		log.Fatalf(err.Error())
	}


}
