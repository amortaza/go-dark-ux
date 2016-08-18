package radiobox

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux"
)

type State struct {
	RadioboxId                   string
	IsChecked                    bool
	IsEnabled                    bool

	Label_                       string

	Left_, Top_, Width_, Height_ int

	OnClick                      func()
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixRadioboxId string) *State {
	buttonId := bl.Current_Node.Id + "/" + postfixRadioboxId

	gCurState = ensureState(buttonId)

	div()

	return gCurState
}

func (s *State) On(cb func(interface{})) {

	gCurState = s
}

func div() {

	radioboxId := gCurState.RadioboxId

	//fmt.Println("checkbox Id " + checkboxId)
	state := gCurState

	bl.Div()
	{
		bl.Id(radioboxId)

		bl.CustomRenderer(func(node *bl.Node) {

			if state.IsChecked {
				if state.IsEnabled {
					go_dark_ux.DrawRadiobox_Checked(0, 0, node.Width, node.Height, state.Label_)
				} else {
					//go_dark_ux.DrawCheckbox_Checked_Disabled(0, 0, node.Width, node.Height, state.Label_)
				}

			} else {
				if state.IsEnabled {
					go_dark_ux.DrawRadiobox_Unchecked(0, 0, node.Width, node.Height, state.Label_)
				} else {
					//go_dark_ux.DrawCheckbox_Unchecked_Disabled(0, 0, node.Width, node.Height, state.Label_)
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

func init() {
	plugin = &Plugin{}
	bl.Plugin(plugin)
}

