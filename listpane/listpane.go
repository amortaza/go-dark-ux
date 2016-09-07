package listpane

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/hover"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"fmt"
)

func fake() {
    var _ = fmt.Println
}

func init() {
	g_stateById = make(map[string] *State)
}

// Shared variable across Div()/End()
var g_curState *State

func Id(listpaneId string) *State {
	g_curState = ensureState(listpaneId)

	g_curState.Z_ItemLabels.Init()

	bl.Div()
	{
		bl.Id(listpaneId)
	}

	return g_curState
}

func Item(label string) {
	g_curState.Z_ItemLabels.PushBack(label)
}

func (s *State) SettleBoundary() {
	bl.Pos(s.Z_Left, s.Z_Top)
	bl.Dim(s.Z_Width, s.Z_Height)
	bl.SettleBoundary()

	border.Draw()
}

func SettleBoundary() {
	g_curState.SettleBoundary()
}

func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func SettleKids() {
	state := g_curState

	for e := state.Z_ItemLabels.Front(); e != nil; e = e.Next() {
		itemLabel := e.Value.(string)

		isHover, ok := state.Z_ItemHover[itemLabel]

		var textColor []int = label.White
		hasBack := false

		if ok && isHover {
			hasBack = true
			textColor = label.Orange
		}

		label.Id(itemLabel).Width(200).Height(40).Label(itemLabel).FontSize(30).Color1v(textColor).HasBack(hasBack)
		{

			if hasBack {
				label.BackColor4i(label.Orange[0], label.Orange[1], label.Orange[2], 50)
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

			bl.SettleBoundary()
			//border.Draw()

			//bl.PushFunc(func() {
			//	fmt.Println("(0) Label Docker!", )
			//	bl.DivId(itemLabel)
				docker.Id().AnchorLeft().AnchorRight().End()
				//bl.End()
			//})
		}
		label.End()
	}

	vert.Id().Spacing(0).Top(0).End()

	bl.SettleKids()
}

func End() {

	bl.RequireSettledBoundaries()
	bl.RequireSettledKids()

	//fmt.Println("(4) listpane width ", bl.Current_Node.Id, " : ", bl.Current_Node.Width)

	bl.End()
}