package accordian_menu_tile

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/animation"
	"github.com/amortaza/go-bellina-plugins/layout/horiz"
	"github.com/amortaza/go-dark-ux/checkbox"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
)

var g_state *State

func Use() *State {

	g_state = ensureState(bl.Current_Node.Id)

	div()

	return g_state
}

func div() {

	bl.Div()
	{
		bl.Id("red")

		bl.Pos(50,50)
		bl.Dim(640, 580)

		bl.Div()
		{
			bl.Id("ace")
			bl.Dim(100,250)
			bl.Pos(10,100)

			border.Wire(50,200,200)

			items()

			docker.Use().AnchorLeft(5).AnchorBottom(5).AnchorRight(5).End()
			vert.Use().Top(3).Spacing(5).End()
		}
		bl.End()



		header()
	}
}

func items() {

	bl.Div()
	{
		bl.Id("i-1")
		bl.Dim(100,50)
		bl.Pos(10,100)

		border.Wire(150,200,20)
	}
	bl.End()

	bl.Div()
	{
		bl.Id("i-2")
		bl.Dim(100,50)
		bl.Pos(10,100)

		border.Wire(150,200,20)
	}
	bl.End()

	bl.Div()
	{
		bl.Id("i-3")
		bl.Dim(100,50)
		bl.Pos(10,100)

		border.Wire(150,200,20)
	}
	bl.End()

}

func header() {

	bl.Div()
	{
		bl.Id("header")
		bl.Height(150)

		checkbox.Id("1").Label("").Left(10).Top(10).Width(50).Height(90).Checked(g_state.isOpen)

		state := g_state

		checkbox.OnClick(func() {

			start, end := 180, 600

			if state.isOpen {
				start, end = 600, 180
			}

			animation.StartPath("red", "hi", float32(start), float32(end), 50, func(shadow *bl.ShadowNode, value float32) {
				shadow.Height = int(value)
			})

			state.isOpen = !state.isOpen
		})

		checkbox.End()

		label.Id("first").Left(3).Top(3).Width(200).Height(90).Label("Hello World!")
		{
			label.BackColor4i(50, 100, 0, 255)
		}
		label.End()

		border.Wire(50,100,200)

		docker.Use().AnchorLeft(3).AnchorTop(3).AnchorRight(3).End()
		horiz.Use().Left(3).Spacing(10).End()
	}
	bl.End()
}
func End() {

	// div...
	{
		border.Wire(255, 0, 0)
	}
	bl.End()
}


