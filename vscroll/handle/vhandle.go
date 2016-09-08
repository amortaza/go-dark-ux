package vhandle

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/drag"
	"math"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	vhandleId := bl.Current_Node.Id + "/" + postfixId

	g_curState = EnsureState(vhandleId)

	return g_curState
}

func div() {

	vhandleId := g_curState.VHandleId

	state := g_curState

	bl.Div()
	{
		bl.Id(vhandleId)

		bl.Pos(state.Left_, state.Top_)
		bl.Dim(state.Thickness_, state.Height_)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.Width, node.Height, "")
		}, false)
	}
}

func On(cb func(float32)) {
	g_curState.onScrollCallback = cb
}

func End() {

	div()

	state := g_curState

	drag.On(func(v interface{}) {
		e := v.(drag.Event)
		handle := e.Target

		totalHeight := handle.Parent.Height

		handleTop := int(math.Max(1, float64(handle.Top)))

		maxTop := totalHeight - handle.Height - 1
		handleTop = int(math.Min(float64(maxTop), float64(handleTop)))

		bl.EnsureShadowByNode(handle).Left__Self_and_Node(2, "*")
		bl.EnsureShadowByNode(handle).Top__Self_and_Node(handleTop, "*")

		var pct float32 = float32(handleTop) / (float32(handle.Parent.Height) - float32(handle.Height))

		if state.onScrollCallback != nil {
			state.onScrollCallback(pct)
		}
	})

	shadow := bl.EnsureShadow()

	shadow.Left__Node_Only("*")
	shadow.Top__Node_Only("*")

	bl.End()
}