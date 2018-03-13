package hscroll

import "github.com/amortaza/go-bellina"

func init() {
	g_stateById = make(map[string] *State)
}

type State struct {

	Z_HScrollId                         string

	Z_Left, Z_Top, Z_Width, Z_Thickness int

	onScrollCallback                    func(float32)
}

// Shared variable across Div()/End()
var g_curState *State

var g_stateById  map[string] *State

func ensureState(hscrollId string) *State {

	state, ok := g_stateById[hscrollId]

	if !ok {

		state = &State{Z_HScrollId: hscrollId, Z_Width: 95, Z_Thickness: 50}

		g_stateById[hscrollId] = state
	}

	return state
}

func (s *State) SettleBoundary() {

	state := g_curState

	{
		bl.Pos(state.Z_Left, state.Z_Top)
		bl.Dim(state.Z_Width, 20)
		bl.SettleBoundary()
	}
}

func (s *State) Left(left int) (*State){

	s.Z_Left = left

	return s
}

func (s *State) Top(top int) (*State){

	s.Z_Top = top

	return s
}

func (s *State) Width(w int) (*State){

	s.Z_Width = w

	return s
}

func (s *State) Thickness(thickness int) (*State){

	s.Z_Thickness = thickness

	return s
}

func (s *State) End() (*State){

	End()

	return s
}

