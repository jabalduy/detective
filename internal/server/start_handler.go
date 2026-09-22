package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// СТАРТ
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func startHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Некорректная команда", http.StatusMethodNotAllowed)
		return
	}

	sessionID := getSessionID(w, r)
	state, err := newGame()
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	session := &Session{
		State: state,
	}

	gamesMu.Lock()
	games[sessionID] = session
	gamesMu.Unlock()

	response := StartResponse{
		Status:    "started",
		Title:     state.Case.Title,
		Message:   "Расследование начато",
		Intro:     state.Case.Intro,
		KnownInfo: state.Case.KnownInfo,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить ответ", http.StatusInternalServerError)
		return
	}
}
