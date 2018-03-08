package checkbox

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

func init() {
	gStateByNode = make(map[string] *State)
}

// Shared variable across Div()/End()
var gCurState *State

func Id(postfixCheckboxId string) *State {

	buttonId := bl.Current_Node.Id + "/" + postfixCheckboxId

	gCurState = ensureState(buttonId)

	div()

	return gCurState
}

func div() {

	checkboxId := gCurState.CheckboxId

	state := gCurState

	bl.Div()
	{
		bl.Id(checkboxId)

		bl.CustomRenderer(func(node *bl.Node) {

			if state.IsChecked {

				if state.IsEnabled {
					ux_checked.Draw(0, 0, node.width, node.height, state.Label_)

				} else {
					ux_checked_disabled.Draw(0, 0, node.width, node.height, state.Label_)
				}

			} else {
				if state.IsEnabled {
					ux_unchecked.Draw(0, 0, node.width, node.height, state.Label_)

				} else {
					ux_unchecked_disabled.Draw(0, 0, node.width, node.height, state.Label_)
				}
			}

		}, false)
	}
}

func OnClick(cb func()) {

	gCurState.OnClick = cb
}

func End() {

	state := gCurState

	click.On_WithLifeCycle(

		// click
		func(i interface{}) {

			if state.IsEnabled {

				state.IsChecked = !state.IsChecked
				state.Dirty = true

				if state.OnClick != nil {
					state.OnClick()
				}
			}
		},

		// on down
		func(i interface{}) {
		},

		// on miss
		func(i interface{}) {
		} )

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, state.Height_)

	if state.Dirty {
		bl.Dirty()
		state.Dirty = false
	}

	bl.End()

	state.ValuesComeFromSourceCode = false
}

