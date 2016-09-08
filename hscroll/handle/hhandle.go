package hhandle

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
	hhandleId := bl.Current_Node.Id + "/" + postfixId

	g_curState = EnsureState(hhandleId)

	return g_curState
}

func div() {

	hhandleId := g_curState.HHandleId

	state := g_curState

	bl.Div()
	{
		bl.Id(hhandleId)

		bl.Pos(state.Left_, state.Top_)
		bl.Dim(state.Width_, state.Thickness_)

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

		totalWidth := handle.Parent.Width

		handleLeft := int(math.Max(1, float64(handle.Left)))

		maxLeft := totalWidth - handle.Width - 1
		handleLeft = int(math.Min(float64(maxLeft), float64(handleLeft)))

		bl.EnsureShadowByNode(handle).Left__Self_and_Node(handleLeft, "hhandle")
		bl.EnsureShadowByNode(handle).Top__Self_and_Node(2, "hhandle")

		//pctStart := float32(handle.Left) / float32(totalWidth)
		//pctEnd := float32(handle.Left + handle.Width) / float32(totalWidth)

		//if cb != nil {
		//	cb(newEvent(pctStart, pctEnd))
		//}

		var pct float32 = float32(handleLeft) / (float32(handle.Parent.Width) - float32(handle.Width))
		state.onScrollCallback(pct)
	})

	bl.EnsureShadow().Top__Node_Only("hhandle")

	bl.End()
}