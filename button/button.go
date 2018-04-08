package button

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

type State struct {

	buttonId string

	isHover   bool
	isDown    bool
	isEnabled bool

	label string

	left, top, width, height int
	padding int

	onClick func()

	dirty bool
}

// Shared variable across Div()/End()
var g_state *State

func init() {

	g_stateByNode = make(map[string] *State)
}

func Id(postfixButtonId string) *State {
	
	buttonId := bl.Current_Node.Id + "/" + postfixButtonId

	g_state = ensureState(buttonId)

	div()

	return g_state
}

func (s *State) On(cb func(interface{})) {

	g_state = s
}

func div() {

	buttonId := g_state.buttonId
	state := g_state

	bl.Div()
	{
		bl.Id(buttonId)

		bl.CustomRenderer(func(node *bl.Node) {

			if !state.isEnabled {

				ux_disabled.SetInt("inPadding_left", state.padding)
				ux_disabled.SetInt("inPadding_top", state.padding)
				ux_disabled.Draw(0, 0, node.Width(), node.Height(), state.label)

			} else if state.isDown {

				ux_pressed.SetInt("inPadding_left", state.padding)
				ux_pressed.SetInt("inPadding_top", state.padding)
				ux_pressed.Draw(0,  0, node.Width(), node.Height(), state.label)

			} else {

				ux_default.SetInt("inPadding_left", state.padding)
				ux_default.SetInt("inPadding_top", state.padding)
				ux_default.Draw(0,  0, node.Width(), node.Height(), state.label)
			}

		}, false)
	}
}

func (s *State) Label(label string) (*State){

	s.label = label

	return s
}

func (s *State) Padding(padding int) (*State){

	s.padding = padding

	return s
}

func (s *State) Left(left int) (*State){

	s.left = left

	return s
}

func (s *State) Top(top int) (*State){

	s.top = top

	return s
}

func (s *State) Width(w int) (*State){

	s.width = w

	return s
}

func (s *State) Height(h int) (*State){

	s.height = h

	return s
}

func (s *State) End() (*State){

	End()

	return s
}

func OnClick(cb func()) {

	g_state.onClick = cb
}

func End() {

	state := g_state

/*	hover.On(func(i interface{}){
		e := i.(*hover.Event)

		if e.IsInEvent {
			state.isHover = true
		} else {
			state.isHover = false
		}
	})*/

	//fmt.Println(state.left, state.top, state.width, state.height)

	click.On_WithLifeCycle(

		// click
		func(i interface{}) {

			state.isDown = false
			state.dirty = true

			if state.isEnabled && state.onClick != nil {
				state.onClick()
			}
		},

		// on down
		func(i interface{}) {

			state.isDown = true
			state.dirty = true
		},

		// on miss
		func(i interface{}) {

			state.isDown = false
			state.dirty = true
		} )

	bl.Pos(state.left, state.top)
	bl.Dim(state.width + state.padding * 2, state.height + state.padding * 2)

	if state.dirty {
		bl.Dirty()
		state.dirty = false
	}

	bl.End()
}

