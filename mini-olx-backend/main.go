package main

import (
	"log"
	"mini-olx-backend/database"
	"mini-olx-backend/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

func main() {
	database.CreateConn()

	router := mux.NewRouter()
	router.HandleFunc("/api/ads", handlers.List).Methods(http.MethodGet)
	router.HandleFunc("/api/ads", handlers.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/ads/{id}", handlers.Delete).Methods(http.MethodDelete)

	router.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	allowedMethods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodOptions,
	}

	handler := cors.
		New(cors.Options{AllowedMethods: allowedMethods}).
		Handler(router)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)

	if err != nil {
		log.Fatal(err)
	}

}
