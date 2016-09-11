package listpane

import "container/list"

type State struct {
	Z_ListPaneId                     string

	Z_Left, Z_Top, Z_Width, Z_Height int

	Z_ItemLabels                     *list.List
	Z_ItemHover                      map[string] bool

	offset int
	paneId string
}

var g_stateById  map[string] *State

func ensureState(listpaneId string) *State {
	state, ok := g_stateById[listpaneId]

	if !ok {
		state = &State{Z_ListPaneId: listpaneId, paneId: listpaneId + "/pane", Z_Width: 95, Z_Height: 150, Z_ItemLabels: list.New(), Z_ItemHover: make(map[string] bool)}

		g_stateById[listpaneId] = state
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

