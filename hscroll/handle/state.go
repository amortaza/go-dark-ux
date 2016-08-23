package hhandle

type State struct {
	HHandleId           string

	Left_, Top_, Width_, Height_ int
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {
	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{HHandleId: EditId, Width_: 95, Height_: 50}

		g_stateById[EditId] = state
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

func (s *State) Height(w int) (*State){
	s.Height_ = w

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

