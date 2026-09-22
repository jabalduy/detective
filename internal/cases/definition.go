package cases

import "detective/internal/game"

type CaseDefinition struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Clues     []game.Clue     `json:"-"`
	Locations []game.Location `json:"-"`
	Objects   []game.Object   `json:"-"`
	Suspects  []game.Suspect  `json:"-"`
	Dialogues []game.Dialogue `json:"-"`
}
