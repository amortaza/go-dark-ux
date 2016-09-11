package listpane_scroll

type State struct {
	ScrollId                         string

	listpaneId string
	vscrollId string

	Z_Left, Z_Top, Z_Width, Z_Height int
}

var g_stateById  map[string] *State

func ensureState(scrollId, listpaneId, vscrollId string) *State {
	state, ok := g_stateById[scrollId]

	if !ok {
		state = &State{ScrollId: scrollId, listpaneId: listpaneId, vscrollId: vscrollId, Z_Width: 95, Z_Height: 150}

		g_stateById[scrollId] = state
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

