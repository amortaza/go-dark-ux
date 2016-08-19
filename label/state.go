package label


var gStateByNode  map[string] *State

func ensureState(labelId string) *State {
	state, ok := gStateByNode[labelId]

	if !ok {
		state = &State{LabelId: labelId, Label_: "OK", Width_: 95, Height_: 30}

		gStateByNode[labelId] = state
	}

	return state
}
