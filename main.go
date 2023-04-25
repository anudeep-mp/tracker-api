package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anudeep-mp/tracker/router"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")

	routeHandler := router.Router()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:7777", "https://anufolio-2point0.netlify.app", "https://anudeep-m.netlify.app", "http://localhost:5173", "https://folio-tracker.netlify.app"},
		AllowedHeaders: []string{"*"},
	})

	handler := corsHandler.Handler(routeHandler)

	fmt.Println("Server is getting ready")
	log.Fatal(http.ListenAndServe(":"+port, handler))
	fmt.Printf("Listening at port %v", port)
}
