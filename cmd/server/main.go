package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func (a *Application) Serve() error {
	port := os.Getenv("PORT")
	log.Printf("Server is working on %s port!\n", port)

	serv := &http.Server{
		Addr: ":" + port,
		// Handler: router,
	}
	return serv.ListenAndServe()
}

func main() {
	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	// TODO: подключение к базе данных

	app := Application{
		Config: cfg,
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
