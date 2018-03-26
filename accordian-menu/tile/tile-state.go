package accordian_menu_tile

import "github.com/amortaza/go-bellina"

type State struct {

	title string
	isOpen bool
}

var g_stateById map[string] *State

func init() {

	g_stateById = make(map[string] *State)
}

func ensureState(nodeId string) *State {

	state, ok := g_stateById[nodeId]

	if !ok {

		state = &State{isOpen: true}

		g_stateById[nodeId] = state

		bl.OnFreeNode(onFreeNode)
	}

	return state
}

func onFreeNode(nodeId string) {

	delete(g_stateById, nodeId)
}

func (state *State) Title(title string) {

	state.title = title
}