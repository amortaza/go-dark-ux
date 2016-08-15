package checkbox


var gStateByNode  map[string] *State

func ensureState(CheckboxId string) *State {
	state, ok := gStateByNode[CheckboxId]

	if !ok {
		state = &State{CheckboxId: CheckboxId, Label_: "OK", Width_: 95, Height_: 30, IsEnabled: true}

		gStateByNode[CheckboxId] = state
	}

	return state
}
