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

	g_curState = ensureState(hhandleId)

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

func End() {

	div()

	drag.On(func(v interface{}) {
		e := v.(drag.Event)
		handle := e.Target

		totalWidth := handle.Parent.Width

		handleLeft := int(math.Max(4, float64(handle.Left)))

		maxLeft := totalWidth - handle.Width - 4
		handleLeft = int(math.Min(float64(maxLeft), float64(handleLeft)))

		bl.EnsureShadowByNode(handle).PosLeft__Self_and_Node(handleLeft)
		bl.EnsureShadowByNode(handle).PosTop__Self_and_Node(4)

		//pctStart := float32(handle.Left) / float32(totalWidth)
		//pctEnd := float32(handle.Left + handle.Width) / float32(totalWidth)

		//if cb != nil {
		//	cb(newEvent(pctStart, pctEnd))
		//}
	})

	bl.EnsureShadow().PosTop__Node_Only()

	bl.End()
}