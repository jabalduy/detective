package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type StartResponse struct {
	Status    string `json:"status"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Intro     string `json:"intro"`
	KnownInfo string `json:"knownInfo"`
}

type InspectResponse struct {
	Object           Object `json:"object"`
	ClueFound        bool   `json:"clueFound"`
	ClueName         string `json:"clueName"`
	ClueAbout        string `json:"clueAbout"`
	DialogueUnlocked bool   `json:"dialogueUnlocked"`
}

type CaseResponse struct {
	Intro           string     `json:"intro"`
	KnownInfo       string     `json:"knownInfo"`
	FoundClues      []Clue     `json:"foundClues"`
	KnownFacts      []Dialogue `json:"knownFacts"`
	CluesFound      int        `json:"cluesFound"`
	CluesTotal      int        `json:"cluesTotal"`
	ObjectsSearched int        `json:"objectsSearched"`
	ObjectsTotal    int        `json:"objectsTotal"`
	DialoguesAsked  int        `json:"dialoguesAsked"`
	DialoguesTotal  int        `json:"dialoguesTotal"`
}

type AccuseResponse struct {
	TypeEnd string `json:"typeEnd"`
	TextEnd string `json:"textEnd"`
}

// type ObjectResponse struct {
//     Name     string `json:"name"`
//     Searched bool   `json:"searched"`
//     ObjID    int    `json:"objID"`
// }

// type LocationResponse struct {
//     Name     string `json:"name"`
//     About bool   `json:"about"`
//     LocID    int    `json:"locID"`
// }

var games = make(map[string]*GameState)
var gamesMu sync.Mutex
var gameMu sync.Mutex

// ИГРОВАЯ СЕССИЯ
func getGame(w http.ResponseWriter, r *http.Request) (*GameState, error) {
	sessionID := getSessionID(w, r)

	gamesMu.Lock()
	defer gamesMu.Unlock()

	game, ok := games[sessionID]
	if ok {
		return game, nil
	}

	game, err := newGame()
	if err != nil {
		return nil, err
	}
	games[sessionID] = game

	return game, nil
}

func newSessionID() string {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

func newGame() (*GameState, error) {
	caseDef, err := LoadCase("cases/case_017")
	if err != nil {
		return nil, err
	}

	return &GameState{
		Suspects:  caseDef.Suspects,
		Clues:     caseDef.Clues,
		Locations: caseDef.Locations,
		Objects:   caseDef.Objects,
		Dialogues: caseDef.Dialogues,

		FoundClues:      make(map[int]bool),
		SearchedObjects: make(map[int]bool),
		AskedDialogues:  make(map[DialogueID]bool),

		OpenDialogues: initOpenDialogues(caseDef.Dialogues),
	}, nil
}

func getSessionID(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("sessionID")

	if err == nil {
		return cookie.Value
	}

	sessionID := newSessionID()

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})

	return sessionID
}

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
	game, err := newGame()
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gamesMu.Lock()
	games[sessionID] = game
	gamesMu.Unlock()

	response := StartResponse{
		Status:    "started",
		Title:     "ДЕЛО №17 - ПОСЛЕДНИЙ ЭКЗЕМПЛЯР",
		Message:   "Расследование начато",
		Intro:     CaseIntro(),
		KnownInfo: CaseKnownInfo(),
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
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(game.Locations)
	if err != nil {
		http.Error(w, "Не удалось отправить комнаты", http.StatusInternalServerError)
		return
	}
}

func objectsHandler(w http.ResponseWriter, r *http.Request) {
	// ПОЛУЧИТЬ ИГРУ
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
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

	objectsResponse := make([]ObjectResponse, 0)

	for _, obj := range game.Objects {
		if obj.LocID == locID {
			newObj := ObjectResponse{
				LocID:    obj.LocID,
				Name:     obj.Name,
				About:    obj.About,
				ObjID:    obj.ObjID,
				IsClue:   obj.IsClue,
				Key:      obj.Key,
				Searched: game.SearchedObjects[obj.ObjID],
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
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
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

	for i := range game.Objects {
		if game.Objects[i].ObjID == objID {

			game.SearchedObjects[objID] = true
			response := InspectResponse{
				Object: game.Objects[i],
			}

			if game.Objects[i].IsClue {
				for clueIndex := range game.Clues {
					if game.Clues[clueIndex].ObjID == objID {
						game.FoundClues[objID] = true

						response.ClueFound = true
						response.ClueName = game.Clues[clueIndex].Name
						response.ClueAbout = game.Clues[clueIndex].About

						for dialIndex := range game.Dialogues {
							if game.Dialogues[dialIndex].ObjID == objID {
								game.Dialogues[dialIndex].IsClue = true
								game.OpenDialogues[game.Dialogues[dialIndex].ID()] = true

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
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(game.Suspects)
	if err != nil {
		http.Error(w, "Не удалось отправить подозреваемых", http.StatusInternalServerError)
		return
	}
}

func dialoguesHandler(w http.ResponseWriter, r *http.Request) {
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	susIDstr := r.URL.Query().Get("susID")

	susID, err := strconv.Atoi(susIDstr)
	if err != nil {
		http.Error(w, "Некорректный susID", http.StatusBadRequest)
		return
	}

	var dialogues []Dialogue

	for i := range game.Dialogues {
		if game.Dialogues[i].SusID == susID && game.OpenDialogues[game.Dialogues[i].ID()] {
			dialogues = append(dialogues, game.Dialogues[i])
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
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
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

	var response Dialogue

	found := false
	for i, dialogue := range game.Dialogues {
		if game.Dialogues[i].SusID == susID &&
			game.Dialogues[i].DialID == dialID &&
			game.OpenDialogues[dialogue.ID()] {
			dialKey := DialogueID{
				SusID:  dialogue.SusID,
				DialID: dialogue.DialID,
			}

			game.AskedDialogues[dialKey] = true
			response = game.Dialogues[i]
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

func initOpenDialogues(dialogues []Dialogue) map[DialogueID]bool {
	openDialogues := make(map[DialogueID]bool)

	for _, dialogue := range dialogues {
		if dialogue.InitiallyOpen {
			dialKey := DialogueID{
				SusID:  dialogue.SusID,
				DialID: dialogue.DialID,
			}

			openDialogues[dialKey] = true
		}
	}

	return openDialogues
}

// ДОСЬЕ
func caseHandler(w http.ResponseWriter, r *http.Request) {
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	response := CaseResponse{}

	for _, clue := range game.Clues {
		if game.FoundClues[clue.ObjID] {
			response.FoundClues = append(response.FoundClues, clue)
			response.CluesFound++
		}
	}

	for _, dial := range game.Dialogues {
		dialKey := DialogueID{
			SusID:  dial.SusID,
			DialID: dial.DialID,
		}

		if game.AskedDialogues[dialKey] {
			response.KnownFacts = append(response.KnownFacts, dial)
			response.DialoguesAsked++
		}
	}

	for _, obj := range game.Objects {
		if game.SearchedObjects[obj.ObjID] {
			response.ObjectsSearched++
		}
	}

	response.CluesTotal = len(game.Clues)
	response.DialoguesTotal = len(game.Dialogues)
	response.ObjectsTotal = len(game.Objects)

	response.Intro = CaseIntro()
	response.KnownInfo = CaseKnownInfo()

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить досье", http.StatusInternalServerError)
		return
	}
}

// ОБВИНЕНИЕ
func accuseHandler(w http.ResponseWriter, r *http.Request) {
	game, err := getGame(w, r)
	if err != nil {
		log.Printf("new game: %v", err)
		http.Error(w, "Не удалось загрузить игру", http.StatusInternalServerError)
		return
	}

	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
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

	keys := KeyClues(game) + KeyDial(game) + KeyObj(game)

	response := AccuseResponse{}
	if keys == 7 && susID == 4 {
		response.TypeEnd = "win"
		response.TextEnd = WinEnding()
	} else if keys > 3 && susID == 4 {
		response.TypeEnd = "unsolved"
		response.TextEnd = UnsolvedEnding()
	} else {
		response.TypeEnd = "fail"
		response.TextEnd = FailedEnding()
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить обвинение", http.StatusInternalServerError)
		return
	}
}

// ЗАПУСК СЕРВЕРА
func StartServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/start", startHandler)
	http.HandleFunc("/locations", locationsHandler)
	http.HandleFunc("/objects", objectsHandler)
	http.HandleFunc("/inspect-object", inspectObjectHandler)
	http.HandleFunc("/suspects", suspectsHandler)
	http.HandleFunc("/dialogues", dialoguesHandler)
	http.HandleFunc("/ask-dialogue", askDialogueHandler)
	http.HandleFunc("/case", caseHandler)
	http.HandleFunc("/accuse", accuseHandler)

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
