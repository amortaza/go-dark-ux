package hhandle

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/drag"
	"math"
)

func Id(postfixId string) *State {

	handleId := bl.Current_Node.Id + "/" + postfixId

	g_curState = EnsureState(handleId)

	return g_curState
}

func div() {

	handleId := g_curState.HandleId

	state := g_curState

	bl.Div()
	{
		bl.Id(handleId)

		bl.TopOwner("handle")

		bl.Pos(state.Left_, state.Top_)
		bl.Dim(state.Width_, state.Thickness_)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_enabled.Draw(0, 0, node.Width(), node.Height(), "")
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

		totalWidth := handle.Parent.Width()

		handleLeft := int(math.Max(1, float64(handle.Left())))

		maxLeft := totalWidth - handle.Width() - 1
		handleLeft = int(math.Min(float64(maxLeft), float64(handleLeft)))

		//pctStart := float32(handle.left) / float32(totalWidth)
		//pctEnd := float32(handle.left + handle.width) / float32(totalWidth)

		//if cb != nil {
		//	cb(newEvent(pctStart, pctEnd))
		//}

		var pct = float32(handleLeft) / (float32(handle.Parent.Width()) - float32(handle.Width()))
		state.onScrollCallback(pct)
	})

	bl.End()
}