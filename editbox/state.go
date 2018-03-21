package editbox

type State struct {

	EditId              string

	Text_              string

	HasFocus		bool
	CursorPos 		int

	Left_, Top_, Width_ int

	Dirty bool
}

var g_stateById  map[string] *State

func ensureState(EditId string) *State {

	state, ok := g_stateById[EditId]

	if !ok {
		state = &State{EditId: EditId, Text_: "Hi", Width_: 95}

		g_stateById[EditId] = state
	}

	return state
}

func (s *State) Text(text string) (*State){

	s.Text_ = text

	return s
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

func (s *State) End() (*State){
	End()

	return s
}

