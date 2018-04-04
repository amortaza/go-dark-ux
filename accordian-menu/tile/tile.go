package accordian_menu_tile

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/layout/pack"
	"github.com/amortaza/go-dark-ux/checkbox"
	"github.com/amortaza/go-bellina-plugins/animation"
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

		bl.Dim(640, 580)

		//border.Wire(255, 255, 255)

		bl.Div()
		{
			bl.Id(g_state.tileId + "/body")
			bl.Dim(100, 550)
			bl.Pos(10, 100)

			//border.Wire(50, 200, 200)
		}
	}
}

func header() {

	bl.Div()
	{
		bl.Id(g_state.tileId + "/header")
		bl.Height(150)

		checkbox.Id(g_state.tileId + "/1").Label(g_state.title).Height(30).Checked(g_state.isOpen).PlusBox()
		{
			state := g_state

			checkbox.OnClick(func() {

				// closedHeight is -5 actually
				closedHeight, openedHeight := state.headerHeight, state.headerHeight+state.bodyHeight+5
				start, end := closedHeight, openedHeight

				if state.isOpen {
					start, end = openedHeight, closedHeight
				}

				animation.StartPath(state.tileId+"/red", state.tileId+"/hi", float32(start), float32(end), 50, func(shadow *bl.ShadowNode, value float32) {
					shadow.Height = int(value)
				})

				state.isOpen = !state.isOpen
			})

			//border.Wire(0, 100, 200)

			docker.Use().AnchorLeft(0).AnchorTop(0).AnchorRight(0).End()
		}
		checkbox.End()

		//border.Wire(0,100,200)
		//border.Fill(40,40,40)

		docker.Use().AnchorLeft(3).AnchorTop(0).AnchorRight(3).End()
		pack.Use()

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
		//border.Wire_TopsKids(100, 100, 100)
	}
	bl.End()
}


