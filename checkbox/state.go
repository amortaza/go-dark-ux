package checkbox

type State struct {

	checkboxId string
	isChecked  bool
	isEnabled  bool

	label string

	left, top, width, height int

	onClick func()

	// this is true only for the first time the state is created
	// after that the state value is modified by application flow
	values_come_from_source_code bool

	isDirty bool
	isPlusBox bool
}

var g_stateByNodeId map[string] *State

func ensureState(CheckboxId string) *State {

	state, ok := g_stateByNodeId[CheckboxId]

	if !ok {

		state = &State{checkboxId: CheckboxId, label: "OK", width: 95, height: 30, isEnabled: true, values_come_from_source_code: true}

		g_stateByNodeId[CheckboxId] = state
	}

	return state
}

func (s *State) Label(label string) (*State){
	s.label = label

	return s
}

func (s *State) Left(left int) (*State){
	s.left = left

	return s
}

func (s *State) Top(top int) (*State){
	s.top = top

	return s
}

func (s *State) Width(w int) (*State){
	s.width = w

	return s
}

func (s *State) Height(h int) (*State){
	s.height = h

	return s
}

func (s *State) Enabled(enabled bool) (*State){

	s.isEnabled = enabled

	return s
}

func (s *State) PlusBox() (*State){

	s.isPlusBox = true

	return s
}

func (s *State) Checked(checked bool) (*State){

	if !s.values_come_from_source_code {
		return s
	}

	s.isChecked = checked

	return s
}

func (s *State) End() (*State){
	End()

	return s
}
