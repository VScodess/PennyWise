package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shaikhjunaidx/pennywise-backend/db"
	_ "github.com/shaikhjunaidx/pennywise-backend/docs"
	"github.com/shaikhjunaidx/pennywise-backend/internal/routes"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {

	err := runSwagInit()
	if err != nil {
		fmt.Printf("Error generating Swagger docs: %v\n", err)
		os.Exit(1)
	}

	database := db.InitDB()
	fmt.Println("Connected to the database:", database.Name())

	router := mux.NewRouter()

	routes.SetupUserRoutes(router, database)
	routes.SetupTransactionRoutes(router, database)
	routes.SetupCategoryRoutes(router, database)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/swagger/doc.json", http.FileServer(http.Dir("./docs")))

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
}

func runSwagInit() error {
	cmd := exec.Command("swag", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
