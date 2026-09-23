package server

import (
	"detective/internal/game"
	"encoding/json"
	"log"
	"net/http"
)

type DeductionRequest struct {
	FactIDs []int `json:"fact_ids"`
}

type DeductionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func deductionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request DeductionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if len(request.FactIDs) < 2 {
		http.Error(w, "at least two facts are required", http.StatusBadRequest)
		return
	}

	session, err := getGame(w, r)
	if err != nil {
		log.Printf("deduction: get game: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	session.Mu.Lock()
	defer session.Mu.Unlock()

	state := session.State

	if state == nil {
		http.Error(w, "game not started", http.StatusConflict)
		return
	}

	deduction := game.FindDeduction(state, request.FactIDs)

	if deduction == nil {
		writeDeductionResponse(w, DeductionResponse{
			Success: false,
			Message: "Эти факты не приводят к новому выводу",
		})
		return
	}

	if !game.SolveDeduction(state, *deduction) {
		writeDeductionResponse(w, DeductionResponse{
			Success: false,
			Message: "Эту дедукцию сейчас невозможно выполнить",
		})
		return
	}

	writeDeductionResponse(w, DeductionResponse{
		Success: true,
		Message: "Новый вывод получен",
	})
}

func writeDeductionResponse(w http.ResponseWriter, response DeductionResponse) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("deduction: encode response: %v", err)
	}
}
