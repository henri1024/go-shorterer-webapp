package main

import (
	"go-shorterer/app"
	"go-shorterer/controller"
	"go-shorterer/repository"
	"log"
	"os"

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

	mailWidget := app.NewEmailWidget(generateMailWidget())

	controller := controller.NewMainController(db, mailWidget)

	router := NewRouter(controller)

	log.Fatal(router.Router.Run())
}

func generateMailWidget() (string, string, string, string) {
	from := os.Getenv("MAIL_FROM")
	password := os.Getenv("MAIL_PASSWORD")
	smtphost := os.Getenv("MAIL_HOST")
	smtpport := os.Getenv("MAIL_PORT")
	return from, password, smtphost, smtpport
}
