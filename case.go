package main

type CaseDefinition struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Clues     []Clue     `json:"-"`
	Locations []Location `json:"-"`
	Objects   []Object   `json:"-"`
	Suspects  []Suspect  `json:"-"`
	Dialogues []Dialogue `json:"-"`
}
