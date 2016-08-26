package listpane

import "container/list"

type State struct {
	ListPaneId                   string

	Left_, Top_, Width_, Height_ int

	ItemLabels        *list.List
	ItemHover         map[string] bool
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {
	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{ListPaneId: EditId, Width_: 95, Height_: 150, ItemLabels: list.New(), ItemHover: make(map[string] bool)}

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

func (s *State) Height(h int) (*State){
	s.Height_ = h

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

