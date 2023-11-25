package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server-example/db"
	"server-example/router"
	"server-example/services"
	"server-example/tools"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (a *Application) Serve() error {
	port := os.Getenv("PORT")
	log.Printf("Server is working on %s port!\n", port)

	serv := &http.Server{
		Addr:    ":" + port,
		Handler: router.Routes(),
	}
	return serv.ListenAndServe()
}

func main() {
	if err := tools.LoadEnv(); err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := db.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.DB.Close()

	app := Application{
		Config: cfg,
		Models: services.New(db.DB),
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
