package accordian_menu_tile

import "github.com/amortaza/go-bellina"

type State struct {

	tileId string
	title string
	isOpen bool
	headerHeight, bodyHeight int
}

var g_stateById map[string] *State

func init() {

	g_stateById = make(map[string] *State)
}

func ensureState(tileId string) *State {

	state, ok := g_stateById[tileId]

	if !ok {

		state = &State{tileId: tileId, isOpen: true}

		g_stateById[tileId] = state

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