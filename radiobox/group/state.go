package group

import (
	"container/list"
)

type State struct {
	GroupId             string
	IsEnabled           bool

	Label_              string

	Left_, Top_, Width_ int

	OnClick             func()

	ChoiceLabels        *list.List
	ChoiceIds           *list.List
}

var g_stateById  map[string] *State

func ensureState(RadioId string) *State {
	state, ok := g_stateById[RadioId]

	if !ok {
		state = &State{GroupId: RadioId, Label_: "OK", Width_: 95, IsEnabled: true}

		g_stateById[RadioId] = state
	}

	return state
}

func (s *State) AddChoice(label string) (*State){
	s.ChoiceLabels.PushBack(label)

	return s
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

func (s *State) End() (*State){
	End()

	return s
}

