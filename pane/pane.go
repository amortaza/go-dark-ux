package pane

import (
	"github.com/amortaza/go-bellina"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postPaneId string) *State {
	paneId := bl.Current_Node.Id + "/" + postPaneId
	
	g_curState = ensureState(paneId)

	div()

	return g_curState
}

func div() {

	paneId := g_curState.PaneId

	state := g_curState

	bl.Div()
	{
		bl.Id(paneId)

		parent := bl.Current_Node.Parent
		bl.Pos( parent.left+1, parent.top+1)
		bl.Dim( parent.width-2, parent.height-2)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_default.Draw(0, 0, node.width, node.height, state.Label_)
		}, false)
	}
}

func (s *State) Label(label string) (*State){
	s.Label_ = label

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

func End() {
	bl.End()
}

