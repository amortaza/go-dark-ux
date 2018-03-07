package checkbox

type State struct {

	CheckboxId                   string
	IsChecked                    bool
	IsEnabled                    bool

	Label_                       string

	Left_, Top_, Width_, Height_ int

	OnClick                      func()

	// this is true only for the first time the state is created
	// after that the state value is modified by application flow
	ValuesComeFromSourceCode	bool

	Dirty bool
}

var gStateByNode  map[string] *State

func ensureState(CheckboxId string) *State {

	state, ok := gStateByNode[CheckboxId]

	if !ok {

		state = &State{CheckboxId: CheckboxId, Label_: "OK", Width_: 95, Height_: 30, IsEnabled: true, ValuesComeFromSourceCode: true}

		gStateByNode[CheckboxId] = state
	}

	return state
}

func (s *State) Label(label string) (*State){
	s.Label_ = label

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

func (s *State) Height(h int) (*State){
	s.Height_ = h

	return s
}

func (s *State) Enabled(enabled bool) (*State){

	s.IsEnabled = enabled

	return s
}

func (s *State) Checked(checked bool) (*State){

	if !s.ValuesComeFromSourceCode {
		return s
	}

	s.IsChecked = checked

	return s
}

func (s *State) End() (*State){
	End()

	return s
}
