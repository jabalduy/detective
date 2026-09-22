package game

func requirementMet(state *GameState, req Requirement) bool {
	switch req.Type {
	case "clue_found":
		return state.FoundClues[req.ID]

	case "object_searched":
		return state.SearchedObjects[req.ID]

	case "dialogue_asked":
		dialogueID := DialogueID{
			SusID:  req.SusID,
			DialID: req.DialID,
		}

		return state.AskedDialogues[dialogueID]

	default:
		return false
	}
}

func RequirementMet(state *GameState, requirements []Requirement) bool {
	for _, req := range requirements {
		if !requirementMet(state, req) {
			return false
		}
	}

	return true
}
