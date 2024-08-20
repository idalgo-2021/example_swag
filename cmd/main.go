package main

import (
	"log"

	"github.com/idalgo-2021/example_swag/internal/app"
)

// @title Example API(by idalgo-2021)
// @version 1.0
// @description API description
// @contact.name Michael G.
// @contact.url https://github.com/idalgo-2021
// @contact.email sample@email.com
// @host localhost:8080
// @BasePath /
func main() {

	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
