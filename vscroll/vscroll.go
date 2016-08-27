package vscroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/vscroll/handle"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	vscrollId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(vscrollId)

	return g_curState
}

func GetValue(vscrollId string) float32 {
	state := ensureState(vscrollId)

	handleId := vscrollId + "/handle"
	handleState := vhandle.EnsureState(handleId)

	height := state.Z_Height -  handleState.Height_

	value := float32(handleState.Left_) / float32(height)

	return value
}

func div() {

	vscrollId := g_curState.Z_VScrollId

	state := g_curState

	bl.Div()
	{
		bl.Id(vscrollId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.Width, node.Height, "")
		}, false)

		bl.Pos(state.Z_Left, state.Z_Top)
		bl.Dim( 20, state.Z_Height,)

		vhandle.Id("handle").Height(80).Thickness(20 - 2*2).Left(2).Top(2)
		vhandle.On(state.onScrollCallback)
		vhandle.End()
	}
}

func On(cb func(float32)) {
	g_curState.onScrollCallback = cb
}

func End() {

	div()

	bl.End()
}