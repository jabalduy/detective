package cases

import "detective/internal/game"

type CaseDefinition struct {
	game.CaseInfo

	Clues     []game.Clue     `json:"-"`
	Locations []game.Location `json:"-"`
	Objects   []game.Object   `json:"-"`
	Suspects  []game.Suspect  `json:"-"`
	Dialogues []game.Dialogue `json:"-"`
}

type Endings struct {
	Win      string `json:"win"`
	Unsolved string `json:"unsolved"`
	Fail     string `json:"fail"`
}
