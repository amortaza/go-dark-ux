package label

type State struct {
	S_LabelId                        string
	S_Label                          string
	S_FontSize                       int

	S_ColorRed, S_ColorGreen, S_ColorBlue, S_ColorAlpha int
	S_BackColorRed, S_BackColorGreen, S_BackColorBlue, S_BackColorAlpha int

	S_Left, S_Top, S_Width, S_Height int
}

var g_stateByNode  map[string] *State

func ensureState(labelId string) *State {
	state, ok := g_stateByNode[labelId]

	if !ok {
		state = &State{
			S_LabelId: labelId,
			S_Label: "OK",
			S_Width: 95,
			S_Height: 30,
			S_FontSize: 36.0}

		g_stateByNode[labelId] = state
	}

	return state
}
