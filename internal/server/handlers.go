package server

import (
	"detective/internal/game"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	gamesMu.Lock()
	games[sessionID] = state
	gamesMu.Unlock()

	response := StartResponse{
		Status:    "started",
		Title:     "ДЕЛО №17 - ПОСЛЕДНИЙ ЭКЗЕМПЛЯР",
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

// ПОДОЗРЕВАЕМЫЕ, ДИАЛОГИ
func suspectsHandler(w http.ResponseWriter, r *http.Request) {
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

	err = json.NewEncoder(w).Encode(state.Suspects)
	if err != nil {
		http.Error(w, "Не удалось отправить подозреваемых", http.StatusInternalServerError)
		return
	}
}

func dialoguesHandler(w http.ResponseWriter, r *http.Request) {
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

	susIDstr := r.URL.Query().Get("susID")

	susID, err := strconv.Atoi(susIDstr)
	if err != nil {
		http.Error(w, "Некорректный susID", http.StatusBadRequest)
		return
	}

	var dialogues []game.Dialogue

	for i := range state.Dialogues {
		if state.Dialogues[i].SusID == susID && state.OpenDialogues[state.Dialogues[i].ID()] {
			dialogues = append(dialogues, state.Dialogues[i])
		}
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(dialogues)
	if err != nil {
		http.Error(w, "Не удалось отправить диалог", http.StatusInternalServerError)
		return
	}
}

func askDialogueHandler(w http.ResponseWriter, r *http.Request) {
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

	susIDstr := r.URL.Query().Get("susID")
	susID, err := strconv.Atoi(susIDstr)
	if err != nil {
		http.Error(w, "Некорректный susID", http.StatusBadRequest)
		return
	}

	dialIDstr := r.URL.Query().Get("dialID")
	dialID, err := strconv.Atoi(dialIDstr)
	if err != nil {
		http.Error(w, "Некорректный dialID", http.StatusBadRequest)
		return
	}

	var response game.Dialogue

	found := false
	for i, dialogue := range state.Dialogues {
		if state.Dialogues[i].SusID == susID &&
			state.Dialogues[i].DialID == dialID &&
			state.OpenDialogues[dialogue.ID()] {
			dialKey := game.DialogueID{
				SusID:  dialogue.SusID,
				DialID: dialogue.DialID,
			}

			state.AskedDialogues[dialKey] = true
			response = state.Dialogues[i]
			found = true
		}
	}

	if !found {
		http.Error(w, "Не найден диалог", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить выбранный диалог", http.StatusInternalServerError)
		return
	}
}

// ДОСЬЕ
func caseHandler(w http.ResponseWriter, r *http.Request) {
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

	response := CaseResponse{}

	for _, clue := range state.Clues {
		if state.FoundClues[clue.ObjID] {
			response.FoundClues = append(response.FoundClues, clue)
			response.CluesFound++
		}
	}

	for _, dial := range state.Dialogues {
		dialKey := game.DialogueID{
			SusID:  dial.SusID,
			DialID: dial.DialID,
		}

		if state.AskedDialogues[dialKey] {
			response.KnownFacts = append(response.KnownFacts, dial)
			response.DialoguesAsked++
		}
	}

	for _, obj := range state.Objects {
		if state.SearchedObjects[obj.ObjID] {
			response.ObjectsSearched++
		}
	}

	response.CluesTotal = len(state.Clues)
	response.DialoguesTotal = len(state.Dialogues)
	response.ObjectsTotal = len(state.Objects)

	response.Intro = state.Case.Intro
	response.KnownInfo = state.Case.KnownInfo

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить досье", http.StatusInternalServerError)
		return
	}
}

// ОБВИНЕНИЕ
func accuseHandler(w http.ResponseWriter, r *http.Request) {
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
