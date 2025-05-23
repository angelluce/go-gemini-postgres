package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-gemini-postgres/database"
	"go-gemini-postgres/services"
)

type PromptRequest struct {
	ItemID int `json:"item_id"`
}

func GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := database.GetAllItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener todos los items de la base de datos: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GenerateTextResponse(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Cuerpo de solicitud incorrecta", http.StatusBadRequest)
		return
	}

	item, err := database.GetItemByID(req.ItemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el item de la base de datos: %v", err), http.StatusInternalServerError)
		return
	}

	quemadoPrompt := "¿Podrías darme un resumen corto y atractivo de este producto?"
	fullPrompt := fmt.Sprintf("%s. Información adicional de la base de datos: Nombre: %s, Descripción: %s", quemadoPrompt, item.Name, item.Description)

	geminiResponse, err := services.GenerateTextFromGemini(fullPrompt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar texto desde Gemini: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"response": geminiResponse})
}

func GenerateAudioResponse(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Cuerpo de solicitud incorrecta", http.StatusBadRequest)
		return
	}

	item, err := database.GetItemByID(req.ItemID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener el item de la base de datos: %v", err), http.StatusInternalServerError)
		return
	}

	quemadoPrompt := "Dame una descripción corta de este artículo."
	fullPrompt := fmt.Sprintf("%s. Información adicional de la base de datos: Nombre: %s, Descripción: %s", quemadoPrompt, item.Name, item.Description)

	geminiResponse, err := services.GenerateTextFromGemini(fullPrompt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar texto desde Gemini: %v", err), http.StatusInternalServerError)
		return
	}

	audioBytes, err := services.GenerateAudioFromTextGoogleCloud(geminiResponse)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al generar audio: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(audioBytes)))
	w.Write(audioBytes)
}
