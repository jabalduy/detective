package game

type GameState struct {
	Case CaseInfo

	Suspects   []Suspect
	Clues      []Clue
	Locations  []Location
	Objects    []Object
	Dialogues  []Dialogue
	Facts      []Fact
	Deductions []Deduction
	Motives    []Motive

	FoundClues       map[int]bool
	FoundFacts       map[int]bool
	SearchedObjects  map[int]bool
	AskedDialogues   map[DialogueID]bool
	OpenDialogues    map[DialogueID]bool
	SolvedDeductions map[int]bool
	FoundMotives     map[int]bool
}

func InitOpenDialogues(dialogues []Dialogue) map[DialogueID]bool {
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
