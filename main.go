package main

import (
	"fmt"
	"log"
	"net/http"

	"go-gemini-postgres/config"
	"go-gemini-postgres/database"
	"go-gemini-postgres/handlers"
	"go-gemini-postgres/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	database.InitDB(cfg)
	defer database.DB.Close()

	services.InitGemini(cfg)
	services.InitGoogleCloudTTS(cfg)

	router := mux.NewRouter()

	router.HandleFunc("/api/items", handlers.GetAllItemsHandler).Methods("GET")
	router.HandleFunc("/api/generate/text", handlers.GenerateTextResponse).Methods("POST")
	router.HandleFunc("/api/generate/audio", handlers.GenerateAudioResponse).Methods("POST")

	fmt.Printf("Servidor iniciando en el puerto %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, router))
}
