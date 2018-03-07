package listpane_scroll

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/listpane"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-dark-ux/vscroll"
	"github.com/amortaza/go-bellina-plugins/layout/side-glue"
)

func init() {
	g_stateById = make(map[string] *State)
}

var g_curState *State

func Id(scrollId string) *State {

	g_curState = ensureState(scrollId, scrollId + "/listpane", scrollId + "/vscroll" )

	bl.Div()
	{
		bl.Id(scrollId)
	}

	return g_curState
}

func (s *State) SettleBoundary() {
	{
		bl.Pos(s.Z_Left, s.Z_Top)
		bl.Dim(s.Z_Width, s.Z_Height)
		//bl.Dim(500,500)
		bl.SettleBoundary()

		//border.Fill(0,0,100)
		border.Wire(255, 255, 0)()
	}
}

func SettleBoundary() {
	g_curState.SettleBoundary()
}


func On(cb func(float32)) {
	//g_curState.onScrollCallback = cb
}

func SettleKids() {
	// todo
	// check that if settle kids is getting called 2x blow up

	listPaneAndVScroll()

	bl.SettleKids()

	//resize.Use()

	//padpush.Id().Right(0).End()
}

func listPaneAndVScroll() {
	listpaneId := g_curState.listpaneId
	vscrollId := g_curState.vscrollId

	listpane.Id( listpaneId ).Width(300).Height(40).Left(10).Top(10).SettleBoundary()
	{
		listpane.Item("1 One")
		listpane.Item("2 Two")
		listpane.Item("3 Three")
		listpane.Item("4 Four")
		listpane.Item("5 Five")
		listpane.Item("6 Six")
		listpane.Item("7 Seven")
		listpane.Item("8 Eight")
		listpane.Item("9 Nine")
		listpane.Item("10 Ten")

		listpane.SettleKids()

		docker.Id().AnchorLeft(0).AnchorBottom(0).AnchorTop(0).End()
		//border.Fill(50,0,0)
	}
	listpane.End()

	vscroll.Id(vscrollId).Left(10).Top(10).SettleBoundary()
	{

		vscroll.On(func(value float32) {
			listpane.ScrollToPercent(listpaneId, value)
		})

		docker.Id().AnchorRight(0).AnchorBottom(0).AnchorTop(0).End()
	}
	vscroll.End()

	sideglue.Id().LeftNodeId(listpaneId).RightNodeId(vscrollId).End()
}

func End() {

	// todo
	// do not call SettleKids here because it could get called 2x
	// should check if !settle kids and THEN call settle kids!!!!!!

	bl.RequireSettledBoundary()
	bl.RequireSettledKids()

	bl.End()
}