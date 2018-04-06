package label

type State struct {

	labelId  string
	label    string
	fontsize int

	hasBack                                                         bool
	colorRed, colorGreen, colorBlue, colorAlpha                     int
	back_colorRed, back_colorGreen, back_colorBlue, back_colorAlpha int

	left, top, width, height int

	text_width, text_height int
}

var g_stateByLabelId map[string] *State

func ensureState(labelId string) *State {

	state, ok := g_stateByLabelId[labelId]

	if !ok {

		state = &State{

			labelId:  labelId,
			label:    "OK",
			width:    -1, // -1 indicates that user has not set, so use Text Metrics to use reasonable value.
			height:   -1, // -1 indicates that user has not set, so use Text Metrics to use reasonable value.
			fontsize: 36.0}

		g_stateByLabelId[labelId] = state
	}

	return state
}
