package main

import (
	"fmt"
	"go-db-gorm/pkg/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting up...")

	r := gin.Default()
	routes.RegisterItemRoutes(r)

	http.Handle("/", r)

	// Does it matter whihc one we use?
	log.Fatal(http.ListenAndServe("localhost:9090", r))
	// r.Run("localhost:9090")
}
