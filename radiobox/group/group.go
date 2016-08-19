package radio

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-dark-ux/radiobox/choice"
	"fmt"
	"github.com/amortaza/go-dark-ux/pane"
)

func init() {
	gStateByNode = make(map[string] *State)
}

type State struct {
	RadioId                      string
	IsEnabled                    bool

	Label_                       string

	Left_, Top_, Width_ 	     int

	OnClick                      func()
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixRadioId string) *State {
	radioId := bl.Current_Node.Id + "/" + postfixRadioId

	gCurState = ensureState(radioId)

	div()

	return gCurState
}

func (s *State) On(cb func(interface{})) {

	gCurState = s
}

func div() {

	radioId := gCurState.RadioId

	//state := gCurState

	bl.Div()
	{
		bl.Id(radioId)

		pane.Id("mypane2").Label("Cool").Left(0).Top(0).Width(100).Height(100).End()

		radiobox.Id("One").Label("One!").Left(30).Top(30).Width(160).Height(60)
		radiobox.OnClick(func() {
			fmt.Println("checkbox clicked")
		})
		radiobox.End()

		radiobox.Id("Two").Label("Two").Left(30).Top(110).Width(160).Height(60)
		radiobox.OnClick(func() {
			fmt.Println("checkbox clicked")
		})
		radiobox.End()
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

func (s *State) End() (*State){
	End()

	return s
}

func OnClick(cb func()) {

	gCurState.OnClick = cb
}

func End() {

	state := gCurState

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, 150)

	bl.End()
}

