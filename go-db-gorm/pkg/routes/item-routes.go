package routes

import (
	"go-db-gorm/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterItemRoutes = func(router *mux.Router) {
	router.HandleFunc("/item/", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/item/", controllers.GetItems).Methods("GET")
	router.HandleFunc("/item/{itemId}", controllers.GetItemById).Methods("GET")
	router.HandleFunc("/item/{itemId}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/item/{itemId}", controllers.DeleteItem).Methods("DELETE")

}
