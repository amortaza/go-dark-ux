package hscroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/hscroll/handle"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	hscrollId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(hscrollId)

	return g_curState
}

func div() {

	hscrollId := g_curState.HScrollId

	state := g_curState

	bl.Div()
	{
		bl.Id(hscrollId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.Width, node.Height, "")
		}, false)

		bl.Pos(state.Left_, state.Top_)
		bl.Dim(state.Width_, 50)

		hhandle.Id("handle").Width(40).Thickness(50 - 4*2).Left(4).Top(4).End()
	}
}

func End() {

	div()

	bl.End()
}