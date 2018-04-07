package label

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-ux"
)

var White = []int{255, 255, 255, 170}
var Orange = []int{234,134,60, 250}

func init() {

	g_stateByLabelId = make(map[string] *State)
}

var g_state *State

func Id(labelId string) *State {

	g_state = ensureState(labelId)

	g_state.Color1v(White)
	g_state.BackColor1v(White)

	div()

	return g_state
}

func div() {

	labelId := g_state.labelId
	state := g_state

	bl.Div()
	{
		bl.Id(labelId)

		bl.CustomRenderer(func(node *bl.Node) {

			hasBack := 0

			if state.hasBack {
				hasBack = 1
			}

			ux_default.SetInt("inTextRed", state.colorRed)
			ux_default.SetInt("inTextGreen", state.colorGreen)
			ux_default.SetInt("inTextBlue", state.colorBlue)
			ux_default.SetInt("inTextAlpha", state.colorAlpha)
			ux_default.SetInt("inHasBack", hasBack)
			ux_default.SetInt("inBackRed", state.back_colorRed)
			ux_default.SetInt("inBackGreen", state.back_colorGreen)
			ux_default.SetInt("inBackBlue", state.back_colorBlue)
			ux_default.SetInt("inBackAlpha", state.back_colorAlpha)

			ux_default.SetFloat("FontSize", float32(state.fontsize))

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

func (state *State) HasBack(hasBack bool) (*State){

	if hasBack {
		state.hasBack = true
	} else {
		state.hasBack = false
	}

	return state
}

func (state *State) Color4i(r,g,b,a int) (*State){

	state.colorRed, state.colorGreen, state.colorBlue, state.colorAlpha = r, g, b, a

	return state
}

func (state *State) Color1v(colors []int) (*State){

	state.colorRed, state.colorGreen, state.colorBlue, state.colorAlpha = colors[0], colors[1], colors[2], colors[3]

	return state
}

func (state *State) BackColor1v(colors []int) (*State){

	state.hasBack = true

	state.back_colorRed, state.back_colorGreen, state.back_colorBlue, state.back_colorAlpha = colors[0], colors[1], colors[2], colors[3]

	return state
}

func (state *State) BackColor4i(r,g,b,a int) (*State){

	state.hasBack = true

	state.back_colorRed, state.back_colorGreen, state.back_colorBlue, state.back_colorAlpha = r, g, b, a

	return state
}

func BackColor4i(r,g,b,a int) (*State){

	return g_state.BackColor4i(r,g,b,a)
}

func (state *State) FontSize(size int) (*State){

	if state.fontsize == size {

		return state
	}

	bl.Dirty()

	state.fontsize = size

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

func (state *State) End() (*State){

	bl.Pos(state.left, state.top)

	if state.width == -1 {
		bl.Width(state.text_width)
	} else {
		bl.Width(state.width)
	}

	if state.height == -1 {
		bl.Height(state.text_height)
	} else {
		bl.Height(state.height)
	}

	bl.End()

	return state
}

func calcDims(state *State) {

	if state.width > -1 && state.height > -1 {
		return
	}

	ux.Ctx.SetFontSize(float32(state.fontsize))
	ux.Ctx.SetFontFace("sans")

	state.text_width = int(ux.GetTextWidth(state.label + "W"))
	state.text_height = int(ux.GetTextHeight(state.label))
}