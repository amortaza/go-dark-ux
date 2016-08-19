package label

import (
	"github.com/amortaza/go-bellina"
)

func init() {
	gStateByNode = make(map[string] *State)
}

type State struct {
	LabelId                     string
	Label_                       string

	Left_, Top_, Width_, Height_ int
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixLabelId string) *State {
	labelId := bl.Current_Node.Id + "/" + postfixLabelId

	gCurState = ensureState(labelId)

	div()

	return gCurState
}

func div() {

	labelId := gCurState.LabelId
	state := gCurState

	bl.Div()
	{
		bl.Id(labelId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_default.Draw(0, 0, node.Width, node.Height, state.Label_)
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

func End() {

	state := gCurState

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.Height_)

	bl.End()
}

