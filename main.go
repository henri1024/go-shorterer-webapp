package main

import (
	"log"
	"shorterer/controller"
	"shorterer/repository"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found\n")
	}
}

func main() {

	var (
		db  repository.DB
		err error
	)

	if db, err = repository.MakeDB(); err != nil {
		log.Fatal("error : \v\n", err)
	}

	controller := controller.NewShortUrlController(db)

	router := NewRouter(controller)

	log.Fatal(router.Router.Run())
}
