package listpane_scroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/listpane"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-dark-ux/vscroll"
	"github.com/amortaza/go-bellina-plugins/layout/side-glue"
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

func Id(postfixId string) *State {
	lpScrollId := bl.Current_Node.Id + "/" + postfixId

	g_curState = ensureState(lpScrollId)

	bl.Div()
	{
		bl.Id(lpScrollId)
	}

	return g_curState
}

func (s *State) SettleBoundary() {
	{
		bl.Pos(s.Z_Left, s.Z_Top)
		bl.Dim(s.Z_Width, s.Z_Height)
		bl.SettleBoundary()

		border.Draw()
	}
}

func SettleBoundary() {
	g_curState.SettleBoundary()
}


func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func SettleKids() {
	{


		listpane.Id("BBB").Width(300).Height(200).Left(10).Top(10).SettleBoundary()
		{
			listpane.Item("1 One")
			listpane.Item("2 Two")
			listpane.Item("3 Three")
			listpane.Item("4 Four")
			listpane.Item("5 Five")
			listpane.Item("6 Six")

			listpane.SettleKids()

			sideglue.Id().LeftNodeId("BBB").RightNodeId("CCC").End()

			docker.Id().AnchorLeft().AnchorBottom().AnchorTop().End()
		}
		listpane.End()

		vscroll.Id("CCC").Left(10).Top(10).SettleBoundary()
		{
			docker.Id().AnchorRight().AnchorBottom().AnchorTop().End()
		}
		vscroll.End()


		// to do
		//parentNode := bl.GetNodeById("<listpane-scroll : listpane>")
		//fmt.Println("(8) Width ", parentNode.Id, " : ", parentNode.Width)
	}

	bl.SettleKids()
}

func End() {

	SettleKids()

	bl.RequireSettledBoundaries()
	bl.RequireSettledKids()

	bl.End()
}