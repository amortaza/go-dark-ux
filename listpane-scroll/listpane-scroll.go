package listpane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/listpane"
	"github.com/amortaza/go-dark-ux/vscroll"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	listpaneId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(listpaneId)

	return g_curState
}

func div() {

	listpaneId := g_curState.ListPaneId

	state := g_curState

	bl.Div()
	{
		bl.Id(listpaneId)

		bl.Pos(state.Left_, state.Top_)
		bl.Dim(state.Width_, state.Height_)

		border.Draw()

		listpane.Id("hi").Width(100).Height(200).Left(10).Top(10)
		{
			listpane.Item("One")
			listpane.Item("Two")
			listpane.Item("Three")
		}
		listpane.End()

		vscroll.Id("yo").Left(10).Top(10)
		vscroll.End()


	}
}

func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func End() {

	div()

	//state := g_curState

	bl.End()
}