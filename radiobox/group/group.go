package group

import (
	"github.com/amortaza/go-bellina"
	"container/list"
	"github.com/amortaza/go-dark-ux/radiobox/choice"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
)

func init() {
	g_stateById = make(map[string] *State)
}


// Shared variable across Div()/End()
var gCurState *State

func Id(postfixRadioId string) *State {
	radioId := bl.Current_Node.Id + "/" + postfixRadioId

	gCurState = ensureState(radioId)

	gCurState.ChoiceLabels = list.New()
	gCurState.ChoiceIds = list.New()

	return gCurState
}

func div() {

	groupId := gCurState.GroupId

	state := gCurState

	bl.Div()
	{
		bl.Id(groupId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_group.Draw(0, 0, node.Width, node.Height, state.Label_)
		}, true)

		height := 60

		for e := state.ChoiceLabels.Front(); e != nil; e = e.Next() {
		        label := e.Value.(string)

			choice.Id(label).Label(label).Left(10)/*.Top(top)*/.Width(240).Height(height)

			choiceId := choice.CurState.ChoiceId

			state.ChoiceIds.PushBack(choiceId)

			choice.OnClick(func() {
				for e := state.ChoiceIds.Front(); e != nil; e = e.Next() {
					choiceId := e.Value.(string)

					choiceState := choice.EnsureState(choiceId)
					choiceState.IsChecked = false
				}

				choiceState := choice.EnsureState(choiceId)
				choiceState.IsChecked = true
			})
			choice.End()
		}

		vert.Id("one").Top(50).Spacing(0).End()
	}
}

func End() {

	div()

	state := gCurState

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.ChoiceIds.Len() * 60 + 50 + 10)

	bl.End()
}

func OnClick(cb func()) {

	gCurState.OnClick = cb
}

func AddChoice(label string) {

	gCurState.AddChoice(label)
}


