package editbox

import (
	"github.com/amortaza/go-bellina"
	"math"
	"github.com/amortaza/go-bellina-plugins/focus"
	"github.com/amortaza/go-hal"
)

var g_cursorDirty = false

func init() {
	g_stateById = make(map[string] *State)

	bl.Register_LifeCycle_After_UserTick_LongTerm(longTermOnTick)
}

// Shared variable across Div()/End()
var g_curState *State
var g_cursorOpacity float64
var g_cursorCycle float64

func Id(postfixEditId string) *State {

	editId := bl.Current_Node.Id + "/" + postfixEditId

	g_curState = ensureState(editId)

	return g_curState
}

func div() {

	editId := g_curState.EditId

	state := g_curState

	bl.Div()
	{
		bl.Id(editId)

		bl.CustomRenderer(func(node *bl.Node) {

			var opacity = 0

			if state.HasFocus {
				opacity = int(g_cursorOpacity * 255)
			}

			ux_enabled.SetInt("v_cursorOpacity", opacity)
			ux_enabled.SetInt("v_cursorPos", state.CursorPos)
			ux_enabled.Draw(0, 0, node.Width(), node.Height(), state.Text_)
		}, true)
	}
}

func End() {

	div()

	state := g_curState

	focus.On_LifeCycle( func(focusEvent interface{}) {

		e := focusEvent.(focus.Event)

		if e.KeyEvent.Action == hal.Button_Action_Down {

			key := e.KeyEvent.Key

			processKeyDown(key, e.KeyEvent.Alt, e.KeyEvent.Ctrl, e.KeyEvent.Shift, state)

			state.Dirty = true
		}

	}, func(focusEvent interface{}) {
		state.HasFocus = true
		state.CursorPos = len(state.Text_)
		state.Dirty = true

	}, func(focusEvent interface{}) {
		state.HasFocus = false
		state.Dirty = true
	})

	bl.Pos(state.Left_, state.Top_)
	bl.Dim(state.Width_, 50)

	if state.Dirty || g_cursorDirty {
		bl.Dirty()
		state.Dirty = false
		g_cursorDirty = false
	}

	bl.End()
}


func longTermOnTick() {

	if g_curState != nil && g_curState.HasFocus {

		g_cursorOpacity = (math.Sin(g_cursorCycle) + 1 ) / 2
		g_cursorCycle += .15
		g_cursorDirty = true
	}
}