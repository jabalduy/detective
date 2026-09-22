package game

func CanPerformAction(state *GameState, action Action) bool {
	return RequirementsMet(state, action.Requirements)
}
