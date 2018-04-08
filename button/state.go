package button

var g_stateByNode map[string] *State

func ensureState(buttonId string) *State {

	state, ok := g_stateByNode[buttonId]

	if !ok {

		state = &State{	buttonId: buttonId,
						label: "OK",
						width: 95,
						height: 30,
						isEnabled: true}

		g_stateByNode[buttonId] = state
	}

	return state
}
