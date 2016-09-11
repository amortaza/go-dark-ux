package vscroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/vscroll/handle"
)

func init() {
	g_stateById = make(map[string] *State)
}

var g_curState *State

func Id(vscrollId string) *State {

	g_curState = ensureState(vscrollId)

	bl.Div()
	{
		bl.Id(vscrollId)
	}

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

func (s *State) SettleBoundary() {
	{
		bl.Pos(s.Z_Left, s.Z_Top)
		bl.Dim( 20, s.Z_Height,)
		bl.SettleBoundary()
	}
}

func SettleBoundary() {
	g_curState.SettleBoundary()
}

func On(cb func(float32)) {
	g_curState.onScrollCallback = cb
}

func End() {
	bl.RequireSettledBoundaries()

	state := g_curState

	{
		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.Width, node.Height, "")
		}, false)

		vhandle.Id("handle").Height(80).Thickness(20 - 2 * 2).Left(2).Top(2)
		vhandle.On(state.onScrollCallback)
		vhandle.End()
	}
	bl.End()
}