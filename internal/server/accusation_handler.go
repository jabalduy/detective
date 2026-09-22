package server

import (
	"detective/internal/game"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// ОБВИНЕНИЕ
func accuseHandler(w http.ResponseWriter, r *http.Request) {
	session, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	session.Mu.RLock()
	defer session.Mu.RUnlock()
	state := session.State

	if state == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Некорректная команда", http.StatusMethodNotAllowed)
		return
	}

	susIDstr := r.URL.Query().Get("susID")

	susID, err := strconv.Atoi(susIDstr)
	if err != nil {
		http.Error(w, "Некорректный susID", http.StatusBadRequest)
		return
	}

	keys := game.KeyClues(state) +
		game.KeyDial(state) +
		game.KeyObj(state)

	response := AccuseResponse{}
	if keys == 7 && susID == 4 {
		response.TypeEnd = "win"
		response.TextEnd = state.Case.Endings.Win
	} else if keys > 3 && susID == 4 {
		response.TypeEnd = "unsolved"
		response.TextEnd = state.Case.Endings.Unsolved
	} else {
		response.TypeEnd = "fail"
		response.TextEnd = state.Case.Endings.Fail
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить обвинение", http.StatusInternalServerError)
		return
	}
}
