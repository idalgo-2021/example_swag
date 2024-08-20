package app

import (
	"fmt"
	"net/http"

	"github.com/idalgo-2021/example_swag/internal/repository"
	"github.com/idalgo-2021/example_swag/internal/restapi/handlers"
	"github.com/idalgo-2021/example_swag/internal/restapi/router"
)

type App struct {
	httpServer *http.Server
}

func NewApp() (*App, error) {
	a := &App{}

	if err := a.initDeps(); err != nil {
		return nil, fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	return a, nil
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initHTTPServer,
	}

	for _, f := range inits {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initHTTPServer() error {

	// Создание репозитория
	repo := repository.NewItemRepository()
	// Создание хендлеров
	itemHandler := handlers.NewItemHandler(repo)
	// Создание роутера
	handler := router.NewRouter(itemHandler)

	a.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	return nil
}

func (a *App) Run() error {
	fmt.Printf("Starting server at port 8080\n")
	return a.httpServer.ListenAndServe()
}
