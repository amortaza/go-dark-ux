package pane

type State struct {
	PaneId string

	Label_ string
}

var g_stateById  map[string] *State

func ensureState(paneId string) *State {
	state, ok := g_stateById[paneId]

	if !ok {
		state = &State{PaneId: paneId, Label_: "OK"}

		g_stateById[paneId] = state
	}

	return state
}
