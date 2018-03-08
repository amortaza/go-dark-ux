package choice

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

func init() {
	gStateById = make(map[string] *State)
}

func Id(postChoiceId string) *State {
	choiceId := bl.Current_Node.Id + "/" + postChoiceId

	CurState = EnsureState(choiceId)

	div()

	return CurState
}

func div() {

	radioboxId := CurState.ChoiceId

	state := CurState

	bl.Div()
	{
		bl.Id(radioboxId)

		bl.CustomRenderer(func(node *bl.Node) {

			if state.IsChecked {
				if state.IsEnabled {
					ux_checked.Draw(0, 0, node.width, node.height, state.Label_)
				} else {
					//go_dark_ux.DrawCheckbox_Checked_Disabled(0, 0, node.width, node.height, state.Label_)
				}

			} else {
				if state.IsEnabled {
					ux_unchecked.Draw(0, 0, node.width, node.height, state.Label_)
				} else {
					//go_dark_ux.DrawCheckbox_Unchecked_Disabled(0, 0, node.width, node.height, state.Label_)
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

func (s *State) End() (*State){
	End()

	return s
}

func OnClick(cb func()) {

	CurState.OnClick = cb
}

func End() {

	state := CurState

	click.On( func(i interface{}) {
			state.IsChecked = !state.IsChecked

			if state.IsEnabled && state.OnClick != nil {
				state.OnClick()
			}
	} )

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.Height_)

	bl.End()
}

