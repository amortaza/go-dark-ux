package listpane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/hover"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
)

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(postfixId string) *State {
	listpaneId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(listpaneId)

	g_curState.Z_ItemLabels.Init()

	return g_curState
}

func Item(label string) {
	g_curState.Z_ItemLabels.PushBack(label)
}

func div() {

	listpaneId := g_curState.Z_ListPaneId

	state := g_curState

	bl.Div()
	{
		bl.Id(listpaneId)

		bl.Pos(state.Z_Left, state.Z_Top)
		bl.Dim(state.Z_Width, state.Z_Height)

		border.Draw()
	}
}

func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func End() {

	div()

	state := g_curState

	for e := g_curState.Z_ItemLabels.Front(); e != nil; e = e.Next() {
		itemLabel := e.Value.(string)

		isHover, ok := state.Z_ItemHover[itemLabel]

		var textColor []int = label.White
		hasBack := false

		if ok && isHover {
			hasBack = true
			textColor = label.Orange
		}

		label.Id(itemLabel).Height(40).Label(itemLabel).FontSize(36).Color1v(textColor).HasBack(hasBack)

		if hasBack {
			label.BackColor4i(label.Orange[0],label.Orange[1],label.Orange[2], 50)
		}

		{
			hover.On(func(v interface{}) {
				e := v.(*hover.Event)

				if e.IsInEvent {
					state.Z_ItemHover[ itemLabel ] = true

				} else {
					state.Z_ItemHover[ itemLabel ] = false
				}
			})
		}

		docker.Id().AnchorLeft().AnchorRight().End()

		label.End()
	}

	vert.Id().Spacing(0).Top(0).End()


	bl.End()
}