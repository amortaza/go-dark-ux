package pane


var gStateByNode  map[string] *State

func ensureState(paneId string) *State {
	state, ok := gStateByNode[paneId]

	if !ok {
		state = &State{PaneId: paneId, Label_: "OK", Width_: 95, Height_: 30}

		gStateByNode[paneId] = state
	}

	return state
}
