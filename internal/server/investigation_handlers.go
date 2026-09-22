package server

import (
	"detective/internal/game"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// ЛОКАЦИИ, ОБЪЕКТЫ, РАССЛЕДОВАНИЕ
func locationsHandler(w http.ResponseWriter, r *http.Request) {
	state, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if state == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(state.Locations)
	if err != nil {
		http.Error(w, "Не удалось отправить комнаты", http.StatusInternalServerError)
		return
	}
}

func objectsHandler(w http.ResponseWriter, r *http.Request) {
	// ПОЛУЧИТЬ ИГРУ
	state, err := getGame(w, r)
	if err != nil {
		log.Printf("new state: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if state == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	// ОБРАБРОТКА ВЫБОРА КОМНАТЫ
	locIDStr := r.URL.Query().Get("locID")

	locID, err := strconv.Atoi(locIDStr)
	if err != nil {
		http.Error(w, "Некорректный locID", http.StatusBadRequest)
		return
	}

	objectsResponse := make([]game.ObjectResponse, 0)

	for _, obj := range state.Objects {
		if obj.LocID == locID {
			newObj := game.ObjectResponse{
				LocID:    obj.LocID,
				Name:     obj.Name,
				About:    obj.About,
				ObjID:    obj.ObjID,
				IsClue:   obj.IsClue,
				Key:      obj.Key,
				Searched: state.SearchedObjects[obj.ObjID],
			}
			objectsResponse = append(objectsResponse, newObj)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(objectsResponse)
	if err != nil {
		http.Error(w, "Не удалось отправить объекты", http.StatusInternalServerError)
		return
	}
}

func inspectObjectHandler(w http.ResponseWriter, r *http.Request) {
	state, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if state == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Некорректная команда", http.StatusMethodNotAllowed)
		return
	}

	objIDstr := r.URL.Query().Get("objID")

	objID, err := strconv.Atoi(objIDstr)
	if err != nil {
		http.Error(w, "Некорректный objID", http.StatusBadRequest)
		return
	}

	for i := range state.Objects {
		if state.Objects[i].ObjID == objID {

			state.SearchedObjects[objID] = true
			response := InspectResponse{
				Object: state.Objects[i],
			}

			if state.Objects[i].IsClue {
				for clueIndex := range state.Clues {
					if state.Clues[clueIndex].ObjID == objID {
						state.FoundClues[objID] = true

						response.ClueFound = true
						response.ClueName = state.Clues[clueIndex].Name
						response.ClueAbout = state.Clues[clueIndex].About

						for dialIndex := range state.Dialogues {
							if state.Dialogues[dialIndex].ObjID == objID {
								state.Dialogues[dialIndex].IsClue = true
								state.OpenDialogues[state.Dialogues[dialIndex].ID()] = true

								response.DialogueUnlocked = true
							}
						}

						break
					}
				}
			}

			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(response)
			if err != nil {
				http.Error(w, "Не удалось отправить объект", http.StatusInternalServerError)
			}

			return
		}
	}

	http.Error(w, "Объект не найден", http.StatusNotFound)
}
