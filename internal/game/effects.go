package game

func applyEffect(state *GameState, effect Effect) {
	switch effect.Type {
	case "discover_clue":
		state.FoundClues[effect.ID] = true

	case "open_dialogue":
		id := DialogueID{
			SusID:  effect.SusID,
			DialID: effect.DialID,
		}

		state.OpenDialogues[id] = true
	}
}

func ApplyEffects(state *GameState, effects []Effect) {
	for _, effect := range effects {
		applyEffect(state, effect)
	}
}
