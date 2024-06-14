package main

import (
	"log"

	"github.com/dystrow/budget/api/category"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	e := echo.New()

	category.Init(e)

	e.Logger.Fatal(e.Start(":7777"))
}
