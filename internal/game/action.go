package game

func CanPerformAction(state *GameState, action Action) bool {
	return RequirementsMet(state, action.Requirements)
}

func PerformAction(state *GameState, action Action) bool {
	if !CanPerformAction(state, action) {
		return false
	}

	ApplyEffects(state, action.Effects)
	UpdateOpenDialogues(state)

	return true
}
