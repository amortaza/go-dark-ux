package listpane_scroll

type State struct {
	ListPaneScrollId                 string

	Z_Left, Z_Top, Z_Width, Z_Height int
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {
	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{ListPaneScrollId: EditId, Z_Width: 95, Z_Height: 150}

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

func (s *State) Width(w int) (*State){
	s.Z_Width = w

	return s
}

func (s *State) Height(h int) (*State){
	s.Z_Height = h

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

