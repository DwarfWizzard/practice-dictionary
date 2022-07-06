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

	dictionaryStorage := mock.NewDictionaryMock()
	dictionaryService := service.NewService(&service.ServiceConfig{
		DictStorage: dictionaryStorage,
	})
	dictionaryTransport := rest.NewHandlers(&rest.HandlerConfig{
		DictService: dictionaryService.Dict,
	})

	if err := srv.Run(":8000", dictionaryTransport.InitRoutes); err != nil {
		log.Fatalf(err.Error())
	}
}
