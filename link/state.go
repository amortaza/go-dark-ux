package link

type State struct {

	linkId   string

	label    string

	state 	 int

	left, top, width, height int

	text_width, text_height int

	dirty bool

	onClick func(interface{})

	payload string
}

var g_stateByLinkId map[string] *State

func ensureState(linkId string) *State {

	state, ok := g_stateByLinkId[linkId]

	if !ok {

		state = &State{

			linkId:   linkId,
			label:    "OK",
			width:    -1, // -1 indicates that user has not set, so use Text Metrics to use reasonable value.
			height:   -1 }

		g_stateByLinkId[linkId] = state
	}

	return state
}
