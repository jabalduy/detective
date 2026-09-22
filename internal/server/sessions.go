package server

import (
	"crypto/rand"
	"detective/internal/cases"
	"detective/internal/game"
	"encoding/hex"
	"net/http"
	"sync"
)

var games = make(map[string]*game.GameState)
var gamesMu sync.Mutex
var gameMu sync.Mutex

type Session struct {
	State *game.GameState
	Mu    sync.RWMutex
}

func newSessionID() string {
	bytes := make([]byte, 16)

	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)
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

func getGame(w http.ResponseWriter, r *http.Request) (*game.GameState, error) {
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

func newGame() (*game.GameState, error) {
	caseDef, err := cases.LoadCase("cases/case_017")
	if err != nil {
		return nil, err
	}

	return &game.GameState{
		Case: caseDef.CaseInfo,

		Suspects:  caseDef.Suspects,
		Clues:     caseDef.Clues,
		Locations: caseDef.Locations,
		Objects:   caseDef.Objects,
		Dialogues: caseDef.Dialogues,

		FoundClues:      make(map[int]bool),
		SearchedObjects: make(map[int]bool),
		AskedDialogues:  make(map[game.DialogueID]bool),

		OpenDialogues: game.InitOpenDialogues(caseDef.Dialogues),
	}, nil
}
