package main

type GameState struct {
	Suspects  []Suspect
	Clues     []Clue
	Locations []Location
	Objects   []Object
	Dialogues []Dialogue

	FoundClues      map[int]bool
	SearchedObjects map[int]bool
	AskedDialogues  map[DialogueID]bool
	OpenDialogues   map[DialogueID]bool
}

func main() {
	// clues, err := LoadClues("cases/case_017/clues.json")
	// if err != nil {
	// 	panic(err)
	// }

	// for _, clue := range clues {
	// 	fmt.Println(clue.Name, clue.ObjID, clue.Key, clue.Found)
	// }

	StartServer()
}
