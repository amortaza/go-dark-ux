package radio


var gStateByNode  map[string] *State

func ensureState(RadioId string) *State {
	state, ok := gStateByNode[RadioId]

	if !ok {
		state = &State{RadioId: RadioId, Label_: "OK", Width_: 95, IsEnabled: true}

		gStateByNode[RadioId] = state
	}

	return state
}
