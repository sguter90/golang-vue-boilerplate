package app

import (
	"fmt"
	"gitea.flying-lama.com/apomedica/product-management/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	ServePort int
}

type App struct {
	Db     *sqlx.DB
	Config AppConfig
}

func (app *App) readAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	port, err := strconv.Atoi(os.Getenv("APP_SERVE_PORT"))

	config.ServePort = port

	return config, err
}

func (app *App) Usage() {
	output := ` + + + + + + 
     + App CLI +
     + + + + + +
     
     Usage: app [command]
     
     Following commands are available:
     	serve		Run application webserver
		migrate		Run database migrations
      
    `
	fmt.Println(output)
}

func NewApp() App {
	a := App{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}

	conn, _, err := repository.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	a.Db = conn

	config, err := a.readAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	a.Config = *config

	return a
}
