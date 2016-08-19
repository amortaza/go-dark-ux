package checkbox

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

func init() {
	gStateByNode = make(map[string] *State)
}

type State struct {
	CheckboxId                   string
	IsChecked                    bool
	IsEnabled                    bool

	Label_                       string

	Left_, Top_, Width_, Height_ int

	OnClick                      func()
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixCheckboxId string) *State {
	buttonId := bl.Current_Node.Id + "/" + postfixCheckboxId

	gCurState = ensureState(buttonId)

	div()

	return gCurState
}

func (s *State) On(cb func(interface{})) {

	gCurState = s
}

func div() {

	checkboxId := gCurState.CheckboxId

	state := gCurState

	bl.Div()
	{
		bl.Id(checkboxId)

		bl.CustomRenderer(func(node *bl.Node) {
			if state.IsChecked {
				if state.IsEnabled {
					ux_checked.Draw(0, 0, node.Width, node.Height, state.Label_)
				} else {
					ux_checked_disabled.Draw(0, 0, node.Width, node.Height, state.Label_)
				}

			} else {
				if state.IsEnabled {
					ux_unchecked.Draw(0, 0, node.Width, node.Height, state.Label_)
				} else {
					ux_unchecked_disabled.Draw(0, 0, node.Width, node.Height, state.Label_)
				}
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

func (s *State) Enabled(enabled bool) (*State){
	s.IsEnabled = enabled

	return s
}

func (s *State) Checked(checked bool) (*State){
	s.IsChecked = checked

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

	click.On2(

		// click
		func(i interface{}) {
			state.IsChecked = !state.IsChecked

			if state.IsEnabled && state.OnClick != nil {
				state.OnClick()
			}
		},

		// on down
		func(i interface{}) {
		},

		// on miss
		func(i interface{}) {
		} )

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.Height_)

	bl.End()
}

