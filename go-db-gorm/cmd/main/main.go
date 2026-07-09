package godbgorm

import (
	"fmt"
	"go-db-gorm/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Starting up...")

	r := mux.NewRouter()
	routes.RegisterItemRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
