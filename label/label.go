package label

import (
	"github.com/amortaza/go-bellina"
)

func init() {
	g_stateByNode = make(map[string] *State)
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

	labelId := gCurState.S_LabelId
	state := gCurState

	bl.Div()
	{
		bl.Id(labelId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_default.SetFloat("FontSize", float32(state.S_FontSize))
			ux_default.Draw(0, 0, node.Width, node.Height, state.S_Label)
		}, false)

		//border.Draw()
	}
}

func (s *State) Label(label string) (*State){
	s.S_Label = label

	return s
}

func (s *State) FontSize(size int) (*State){
	s.S_FontSize = size

	return s
}

func (s *State) Left(left int) (*State){
	s.S_Left = left

	return s
}

func (s *State) Top(top int) (*State){
	s.S_Top = top

	return s
}

func (s *State) Width(w int) (*State){
	s.S_Width = w

	return s
}

func (s *State) Height(h int) (*State){
	s.S_Height = h

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

func End() {

	state := gCurState

	bl.Pos(state.S_Left, state.S_Top)
	bl.Dim(state.S_Width, state.S_Height)

	bl.End()
}

