package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
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

func getGame(w http.ResponseWriter, r *http.Request) *GameState {
	sessionID := getSessionID(w, r)

	gamesMu.Lock()
	defer gamesMu.Unlock()

	game, ok := games[sessionID]
	if ok {
		return game
	}

	game = newGame()
	games[sessionID] = game

	return game
}

func newSessionID() string {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
}

func newGame() *GameState {
	return &GameState{
		Suspects:  CreateSuspects(),
		Clues:     CreateClue(),
		Locations: CreateLocations(),
		Objects:   CreateObjects(),
		Dialogues: CreateDialogue(),
	}
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func startHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Некорректная команда", http.StatusMethodNotAllowed)
		return
	}

	sessionID := getSessionID(w, r)
	game := newGame()
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

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить ответ", http.StatusInternalServerError)
		return
	}
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
	gameMu.Lock()
	defer gameMu.Unlock()

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

func suspectsHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(game.Suspects)
	if err != nil {
		http.Error(w, "Не удалось отправить подозреваемых", http.StatusInternalServerError)
		return
	}
}

func objectsHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
	gameMu.Lock()
	defer gameMu.Unlock()

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
	game := getGame(w, r)
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

func dialoguesHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
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
		if game.Dialogues[i].SusID == susID && game.Dialogues[i].IsOpen {
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
	game := getGame(w, r)
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
	for i := range game.Dialogues {
		if game.Dialogues[i].SusID == susID &&
			game.Dialogues[i].DialID == dialID &&
			game.Dialogues[i].IsOpen {
			game.Dialogues[i].Asked = true
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

func caseHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
	gameMu.Lock()
	defer gameMu.Unlock()

	if game == nil {
		http.Error(w, "Игра не запущена", http.StatusBadRequest)
		return
	}

	response := CaseResponse{}

	for _, clue := range game.Clues {
		if clue.Found {
			response.FoundClues = append(response.FoundClues, clue)
			response.CluesFound++
		}
	}

	for _, dial := range game.Dialogues {
		if dial.Asked {
			response.KnownFacts = append(response.KnownFacts, dial)
			response.DialoguesAsked++
		}
	}

	for _, obj := range game.Objects {
		if obj.Searched {
			response.ObjectsSearched++
		}
	}

	response.CluesTotal = len(game.Clues)
	response.DialoguesTotal = len(game.Dialogues)
	response.ObjectsTotal = len(game.Objects)

	response.Intro = CaseIntro()
	response.KnownInfo = CaseKnownInfo()

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Не удалось отправить досье", http.StatusInternalServerError)
		return
	}
}

func accuseHandler(w http.ResponseWriter, r *http.Request) {
	game := getGame(w, r)
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
