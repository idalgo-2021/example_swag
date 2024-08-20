package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/idalgo-2021/example_swag/internal/restapi/handlers"

	_ "github.com/idalgo-2021/example_swag/docs"
)

func NewRouter(itemHandler *handlers.ItemHandler) *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.HandleFunc("/items", itemHandler.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id:[0-9]+}", itemHandler.GetItemByID).Methods("GET")
	router.HandleFunc("/items", itemHandler.GetAllItems).Methods("GET")

	return router
}
