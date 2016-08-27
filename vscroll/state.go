package vscroll

type State struct {
	Z_VScrollId                          string

	Z_Left, Z_Top, Z_Height, Z_Thickness int

	onScrollCallback                     func(float32)
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {
	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{Z_VScrollId: EditId, Z_Height: 95, Z_Thickness: 50}

		g_stateById[EditId] = state
	}

	return state
}

func (s *State) Left(left int) (*State){
	s.Z_Left = left

	return s
}

func (s *State) Top(top int) (*State){
	s.Z_Top = top

	return s
}

func (s *State) Height(h int) (*State){
	s.Z_Height = h

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

