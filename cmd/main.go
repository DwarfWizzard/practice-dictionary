package main

import (
	"log"
	"os"

	dictionary "github.com/DwarfWizzard/practice-dictionary"
	"github.com/DwarfWizzard/practice-dictionary/internal/service"
	"github.com/DwarfWizzard/practice-dictionary/internal/transport/rest"
	"github.com/DwarfWizzard/practice-dictionary/internal/storage/sqlite"
)

func main() {
	port := os.Getenv("PORT")
	dbPath := os.Getenv("DB_PATH")

	srv := new(dictionary.Server)

	db, err := sqlite.NewSQLiteConn(dbPath+"dictionary.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(dbPath+"dictionary.db")

	storage := sqlite.NewStorage(db)
	service := service.NewService(&service.ServiceConfig{
		DictStorage: storage.Dict,
	})
	dictionaryTransport := rest.NewHandler(&rest.HandlerConfig{
		DictService: service.Dict,
	})

	if err := srv.Run(":"+port, dictionaryTransport.InitRoutes); err != nil {
		log.Fatalf(err.Error())
	}


}
