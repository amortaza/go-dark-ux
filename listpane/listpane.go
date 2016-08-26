package listpane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/hover"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	listpaneId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(listpaneId)

	g_curState.ItemLabels.Init()

	return g_curState
}

func Item(label string) {
	g_curState.ItemLabels.PushBack(label)
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
	}
}

func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func End() {

	div()

	state := g_curState

	for e := g_curState.ItemLabels.Front(); e != nil; e = e.Next() {
		itemLabel := e.Value.(string)

		isHover, ok := state.ItemHover[itemLabel]

		size := 20

		if ok && isHover {
			size = 36
		}

		label.Id(itemLabel).Width(100).Height(40).Label(itemLabel).FontSize(size).B
		{
			hover.On(func(v interface{}) {
				e := v.(*hover.Event)

				if e.IsInEvent {
					state.ItemHover[ itemLabel ] = true

				} else {
					state.ItemHover[ itemLabel ] = false
				}
			})
		}
		label.End()
	}

	vert.Id().Spacing(10).Top(10).End()

	bl.End()
}