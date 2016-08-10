package button

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux"
)

type State struct {
	ButtonId string
	IsHover  bool
	IsDown   bool
	IsDisabled	bool

	Label_   string

	Left_, Top_, Width_, Height_ int

	OnClick func()
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixButtonId string) *State {
	buttonId := bl.Current_Node.Id + "/" + postfixButtonId

	gCurState = ensureState(buttonId)

	div()

	return gCurState
}

func (s *State) On(cb func(interface{})) {

	gCurState = s
}

func div() {

	buttonId := gCurState.ButtonId
	state := gCurState

	bl.Div()
	{
		bl.Id(buttonId)

		bl.CustomRenderer(func(node *bl.Node) {
			if state.IsDisabled {
				go_dark_ux.DrawButtonDisabled(node.Left, node.Top, node.Width, node.Height, state.Label_)

			} else if state.IsDown {
				go_dark_ux.DrawButtonPressed(node.Left,  node.Top, node.Width, node.Height, state.Label_)

			} else {
				go_dark_ux.DrawButtonDefault(node.Left,  node.Top, node.Width, node.Height, state.Label_)
			}

		}, false)
	}
}

func (s *State) Label(label string) (*State){
	s.Label_ = label

	return s
}

func (s *State) Left(left int) (*State){
	s.Left_ = left

	return s
}

func (s *State) Top(top int) (*State){
	s.Top_ = top

	return s
}

func (s *State) Width(w int) (*State){
	s.Width_ = w

	return s
}

func (s *State) Height(h int) (*State){
	s.Height_ = h

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

func OnClick(cb func()) {

	gCurState.OnClick = cb
}

func End() {

	state := gCurState

/*	hover.On(func(i interface{}){
		e := i.(*hover.Event)

		if e.IsInEvent {
			state.IsHover = true
		} else {
			state.IsHover = false
		}
	})*/

	//fmt.Println(state.Left_, state.Top_, state.Width_, state.Height_)

	click.On2(

		// click
		func(i interface{}) {
			if state.IsDisabled {
				return
			}

			state.IsDown = false

			if state.OnClick != nil {
				state.OnClick()
			}
		},

		// on down
		func(i interface{}) {
			state.IsDown = true
		},

		// on miss
		func(i interface{}) {
			state.IsDown = false
		} )

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.Height_)

	bl.End()
}

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}

