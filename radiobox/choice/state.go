package choice

type State struct {
	ChoiceId                     string
	IsChecked                    bool
	IsEnabled                    bool

	Label_                       string

	Left_, Top_, Width_, Height_ int

	OnClick                      func()
}

// Shared variable across Div()/End()
var CurState *State

var gStateById  map[string] *State

func EnsureState(choiceId string) *State {
	state, ok := gStateById[choiceId]

	if !ok {
		state = &State{ChoiceId: choiceId, Label_: "OK", Width_: 95, Height_: 30, IsEnabled: true}

		gStateById[choiceId] = state
	}

	return state
}

