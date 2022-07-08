package main

import (
	"log"

	dictionary "github.com/DwarfWizzard/practice-dictionary"
	"github.com/DwarfWizzard/practice-dictionary/internal/service"
	"github.com/DwarfWizzard/practice-dictionary/internal/storage/mock"
	"github.com/DwarfWizzard/practice-dictionary/internal/transport/rest"
)

func main() {
	srv := new(dictionary.Server)

	storage := mock.NewDictionaryMock()
	service := service.NewService(&service.ServiceConfig{
		DictStorage: storage,
	})
	transport := rest.NewHandler(&rest.HandlerConfig{
		DictService: service.DictService,
	})

	if err := srv.Run(":8000", transport.InitRoutes); err != nil {
		log.Fatal(err.Error())
	}


}
