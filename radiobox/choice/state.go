package radiobox


var gStateByNode  map[string] *State

func ensureState(RadioboxId string) *State {
	state, ok := gStateByNode[RadioboxId]

	if !ok {
		state = &State{RadioboxId: RadioboxId, Label_: "OK", Width_: 95, Height_: 30, IsEnabled: true}

		gStateByNode[RadioboxId] = state
	}

	return state
}
