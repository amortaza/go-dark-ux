package vscroll

type State struct {
	VScrollId           string

	Left_, Top_, Height_, Thickness_ int

	onScrollCallback func(float32)
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {
	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{VScrollId: EditId, Height_: 95, Thickness_: 50}

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

func (s *State) Height(h int) (*State){
	s.Height_ = h

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

