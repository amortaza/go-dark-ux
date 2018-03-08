package hscroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/hscroll/handle"
	"fmt"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(hscrollId string) *State {

	g_curState = ensureState(hscrollId)

	bl.Div()
	{
		bl.Id(hscrollId)
	}

	return g_curState
}

func GetValue(hscrollId string) float32 {
	state := ensureState(hscrollId)

	handleId := hscrollId + "/handle"
	handleState := hhandle.EnsureState(handleId)

	width := state.Z_Width -  handleState.Width_

	value := float32(handleState.Left_) / float32(width)

	fmt.Println("value ", value)

	return value
}

func (s *State) SettleBoundary() {

	state := g_curState

	{
		bl.Pos(state.Z_Left, state.Z_Top)
		bl.Dim(state.Z_Width, 20)
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
	bl.RequireSettledBoundary()

	state := g_curState

	{
		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.width, node.height, "")
		}, false)

		hhandle.Id("handle").Width(80).Thickness(20 - 2*2).Left(1).Top(2)
		hhandle.On(state.onScrollCallback)
		hhandle.End()
	}
	bl.End()
}