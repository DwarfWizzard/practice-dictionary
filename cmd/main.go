package dictionary

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	srv := new(dictionary.Server)

	if err := srv.Run(":8000", restHandler.InitRoutes); err != nil {
		log.Fatalf("Run server error: %s", err.Error())
	}
}
