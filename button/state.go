package button


var gStateByNode  map[string] *State

func ensureState(buttonId string) *State {
	state, ok := gStateByNode[buttonId]

	if !ok {
		state = &State{ButtonId: buttonId, Label_: "OK", Width_: 95, Height_: 30}

		gStateByNode[buttonId] = state
	}

	return state
}
