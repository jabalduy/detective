package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type StartResponse struct {
	Status  string `json:"status"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type InspectResponse struct {
	Object           Object `json:"object"`
	ClueFound        bool   `json:"clueFound"`
	ClueName         string `json:"clueName"`
	ClueAbout        string `json:"clueAbout"`
	DialogueUnlocked bool   `json:"dialogueUnlocked"`
}

var game *GameState

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func startHandler(w http.ResponseWriter, r *http.Request) {

	game = &GameState{
		Suspects:  CreateSuspects(),
		Clues:     CreateClue(),
		Locations: CreateLocations(),
		Objects:   CreateObjects(),
		Dialogues: CreateDialogue(),
	}

	response := StartResponse{
		Status:  "started",
		Title:   "ДЕЛО №17 - ПОСЛЕДНИЙ ЭКЗЕМПЛЯР",
		Message: "Расследование начато",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить ответ", http.StatusInternalServerError)
		return
	}
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(game.Locations)
	if err != nil {
		http.Error(w, "Не удалось отправить комнаты", http.StatusInternalServerError)
		return
	}
}

func objectsHandler(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	locIDStr := r.URL.Query().Get("locID")

	locID, err := strconv.Atoi(locIDStr)
	if err != nil {
		http.Error(w, "Некорректный locID", http.StatusBadRequest)
		return
	}

	var objects []Object

	for _, obj := range game.Objects {
		if obj.LocID == locID {
			objects = append(objects, obj)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(objects)
	if err != nil {
		http.Error(w, "Не удалось отправить объекты", http.StatusInternalServerError)
		return
	}
}

func inspectObjectHandler(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	objIDstr := r.URL.Query().Get("objID")

	objID, err := strconv.Atoi(objIDstr)
	if err != nil {
		http.Error(w, "Некорректный objID", http.StatusBadRequest)
		return
	}

	for i := range game.Objects {
		if game.Objects[i].ObjID == objID {

			game.Objects[i].Searched = true
			response := InspectResponse{
				Object: game.Objects[i],
			}

			if game.Objects[i].IsClue {
				for clueIndex := range game.Clues {
					if game.Clues[clueIndex].ObjID == objID {
						game.Clues[clueIndex].Found = true

						response.ClueFound = true
						response.ClueName = game.Clues[clueIndex].Name
						response.ClueAbout = game.Clues[clueIndex].About

						for dialIndex := range game.Dialogues {
							if game.Dialogues[dialIndex].ObjID == objID {
								game.Dialogues[dialIndex].IsClue = true
								game.Dialogues[dialIndex].IsOpen = true

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

func StartServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/locations", locationsHandler)
	http.HandleFunc("/objects", objectsHandler)
	http.HandleFunc("/inspect-object", inspectObjectHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
