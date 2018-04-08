package link

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-ux"
)

func init() {

	g_stateByLinkId = make(map[string] *State)
}

var g_state *State

func On(linkId string,
		onClick func(interface{})) *State {

	g_state = ensureState(linkId)

	g_state.onClick = onClick

	div()

	return g_state
}

func (state *State) Payload(payload string) *State {

	state.payload = payload

	return state
}

func div() {

	labelId := g_state.linkId
	state := g_state

	bl.Div()
	{
		bl.Id(labelId)

		if state.dirty {
			bl.Dirty()
			state.dirty = false
		}

		bl.CustomRenderer(func(node *bl.Node) {

			ux_default.SetInt("inState", state.state)

			ux_default.Draw(0, 0, node.Width(), node.Height(), state.label)

		}, false)
	}
}

func (state *State) Label(label string) (*State){

	if state.label == label {

		return state
	}

	bl.Dirty()

	state.label = label

	calcDims(state)

	return state
}

func (state *State) Left(left int) (*State){

	state.left = left

	return state
}

func (state *State) Top(top int) (*State){

	state.top = top

	return state
}

func (state *State) Width(w int) (*State){

	state.width = w

	return state
}

func (state *State) Height(h int) (*State){

	state.height = h

	return state
}

func End() {
	g_state.End()
}

func (state *State) End() {

	bl.Pos(state.left, state.top)

	if state.width == -1 {
		bl.Width(state.text_width + 20)
	} else {
		bl.Width(state.width)
	}

	if state.height == -1 {
		bl.Height(state.text_height+20)
	} else {
		bl.Height(state.height)
	}

	setup_hover(state)

	bl.End()
}

func calcDims(state *State) {

	if state.width > -1 && state.height > -1 {
		return
	}

	ux.Ctx.SetFontSize(float32(36))
	ux.Ctx.SetFontFace("sans")

	state.text_width = int(ux.GetTextWidth(state.label + "W"))
	state.text_height = int(ux.GetTextHeight(state.label))
}