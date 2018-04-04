package checkbox

import (
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina"
)

func init() {
	g_stateByNodeId = make(map[string] *State)
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

	checkboxId := gCurState.checkboxId

	state := gCurState

	bl.Div()
	{
		bl.Id(checkboxId)

		bl.CustomRenderer(func(node *bl.Node) {

			if state.isChecked {

				if state.isEnabled {
					if state.isPlusBox {
						ux_plusbox__checked.Draw(0, 0, node.Width(), node.Height(), state.label)
					} else {
						ux_checkbox__checked.Draw(0, 0, node.Width(), node.Height(), state.label)
					}

				} else {

					if state.isPlusBox {
						ux_plusbox__checked_disabled.Draw(0, 0, node.Width(), node.Height(), state.label)
					} else {
						ux_checkbox__checked_disabled.Draw(0, 0, node.Width(), node.Height(), state.label)
					}
				}

			} else {
				if state.isEnabled {
					if state.isPlusBox {
						ux_plusbox__unchecked.Draw(0, 0, node.Width(), node.Height(), state.label)
					} else {
						ux_checkbox__unchecked.Draw(0, 0, node.Width(), node.Height(), state.label)
					}

				} else {
					if state.isPlusBox {
						ux_plusbox__unchecked_disabled.Draw(0, 0, node.Width(), node.Height(), state.label)
					} else {
						ux_checkbox__unchecked_disabled.Draw(0, 0, node.Width(), node.Height(), state.label)
					}
				}
			}

		}, false)
	}
}

func OnClick(cb func()) {

	gCurState.onClick = cb
}

func End() {

	state := gCurState

	click.On_WithLifeCycle(

		// click
		func(i interface{}) {

			if state.isEnabled {

				state.isChecked = !state.isChecked
				state.isDirty = true

				if state.onClick != nil {
					state.onClick()
				}
			}
		},

		// on down
		func(i interface{}) {
		},

		// on miss
		func(i interface{}) {
		} )

	bl.Pos(state.left, state.top)
	bl.Dim(state.width, state.height)

	if state.isDirty {
		bl.Dirty()
		state.isDirty = false
	}

	bl.End()

	state.values_come_from_source_code = false
}

