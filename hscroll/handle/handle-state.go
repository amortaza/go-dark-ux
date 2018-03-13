package hhandle

import "github.com/amortaza/go-bellina"

func init() {

	g_stateById = make(map[string] *State)
}

type State struct {

	HandleId string

	Left_, Top_, Width_, Thickness_ int

	onScrollCallback func(float32)

	shadow *bl.Node
}

// Shared variable across Div()/End()
var g_curState *State

var g_stateById  map[string] *State

func EnsureState(handleId string) *State {

	state, ok := g_stateById[handleId]

	if !ok {
		state = &State{
			HandleId: handleId,
			Width_: 95,
			Thickness_: 50}

		g_stateById[handleId] = state
	}

	return state
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

func (s *State) Thickness(thickness int) (*State){
	s.Thickness_ = thickness

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

