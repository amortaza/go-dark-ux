package accordian_menu_tile

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/animation"
	"github.com/amortaza/go-dark-ux/checkbox"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/layout/pack"
)

var g_state *State

func Id(tileId string) *State {

	g_state = ensureState(tileId)

	div()

	return g_state
}

func div() {

	bl.Div()
	{
		bl.Id(g_state.tileId + "/red")

		bl.Pos(50,50)
		bl.Dim(640, 580)

		bl.Div()
		{
			bl.Id(g_state.tileId + "/body")
			bl.Dim(100, 550)
			bl.Pos(10, 100)

			//border.Wire(50, 200, 200)

			// items()
		}
	}
}

func header() {

	bl.Div()
	{
		bl.Id(g_state.tileId + "/header")
		bl.Height(100)

		checkbox.Id(g_state.tileId + "/1").Label(g_state.title).Left(10).Top(10).Width(250).Height(90).Checked(g_state.isOpen)

		state := g_state

		checkbox.OnClick(func() {

			// closedHeight is -5 actually
			closedHeight, openedHeight := state.headerHeight, state.headerHeight + state.bodyHeight + 5
			start, end := closedHeight, openedHeight

			if state.isOpen {
				start, end = openedHeight, closedHeight
			}

			animation.StartPath(state.tileId + "/red", state.tileId + "/hi", float32(start), float32(end), 50, func(shadow *bl.ShadowNode, value float32) {
				shadow.Height = int(value)
			})

			state.isOpen = !state.isOpen
		})

		checkbox.End()

		/*label.Id(g_state.tileId + "/first").Left(3).Top(3).Width(200).Height(90).Label(g_state.title)
		{
			label.BackColor4i(50, 100, 0, 255)
		}
		label.End()*/

		//border.Wire(50,100,200)
		//border.Fill(40,40,40)

		docker.Use().AnchorLeft(3).AnchorTop(0).AnchorRight(3).End()
		//horiz.Use().Left(3).Spacing(10).End()

		g_state.headerHeight = bl.EnsureShadow().Height
	}
	bl.End()
}

func End() {

	// Body
	{
		docker.Use().AnchorLeft(5).AnchorRight(5).AnchorBottom(5).End()
		vert.Use().Top(3).Spacing(5).End()
		pack.Use().Vert().End()

		g_state.bodyHeight = bl.EnsureShadow().Height
	}
	bl.End()

	header()

	// Red

	if g_state.isOpen {
		bl.EnsureShadow().Height = g_state.headerHeight + g_state.bodyHeight + 5
	} else {
		// should be -5 really
		bl.EnsureShadow().Height = g_state.headerHeight
	}

	// div...
	{
		border.Wire_TopsKids(100, 100, 100)
	}
	bl.End()
}


