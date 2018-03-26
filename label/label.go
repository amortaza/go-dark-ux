package label

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

var White = []int{255, 255, 255, 170}
var Orange = []int{234,134,60, 250}

func init() {

	g_stateByNode = make(map[string] *State)
}

var g_curState *State

func Id(labelId string) *State {

	g_curState = ensureState(labelId)

	g_curState.Color1v(White)
	g_curState.BackColor1v(White)

	div()

	return g_curState
}

func div() {

	labelId := g_curState.S_LabelId
	state := g_curState

	bl.Div()
	{
		bl.Id(labelId)

		bl.CustomRenderer(func(node *bl.Node) {
			ux_default.SetInt("inTextRed", state.S_ColorRed)
			ux_default.SetInt("inTextGreen", state.S_ColorGreen)
			ux_default.SetInt("inTextBlue", state.S_ColorBlue)
			ux_default.SetInt("inTextAlpha", state.S_ColorAlpha)

			ux_default.SetInt("inHasBack", state.S_HasBack)
			ux_default.SetInt("inBackRed", state.S_BackColorRed)
			ux_default.SetInt("inBackGreen", state.S_BackColorGreen)
			ux_default.SetInt("inBackBlue", state.S_BackColorBlue)
			ux_default.SetInt("inBackAlpha", state.S_BackColorAlpha)

			ux_default.SetFloat("FontSize", float32(state.S_FontSize))

			ux_default.Draw(0, 0, node.Width(), node.Height(), state.S_Label)
		}, false)
	}
}

func (s *State) Label(label string) (*State){

	if s.S_Label != label {
		fmt.Println(s.S_Label, " vs ", label)
		bl.Dirty()
	}

	s.S_Label = label

	return s
}

func (s *State) HasBack(has bool) (*State){
	if has {
		s.S_HasBack = 1
	} else {
		s.S_HasBack = 0
	}

	return s
}

func (s *State) Color4i(r,g,b,a int) (*State){
	s.S_ColorRed, s.S_ColorGreen, s.S_ColorBlue, s.S_ColorAlpha = r, g, b, a

	return s
}

func (s *State) Color1v(colors []int) (*State){
	s.S_ColorRed, s.S_ColorGreen, s.S_ColorBlue, s.S_ColorAlpha = colors[0], colors[1], colors[2], colors[3]

	return s
}

func (s *State) BackColor1v(colors []int) (*State){
	s.S_HasBack = 1
	s.S_ColorRed, s.S_ColorGreen, s.S_ColorBlue, s.S_ColorAlpha = colors[0], colors[1], colors[2], colors[3]

	return s
}

func (s *State) BackColor4i(r,g,b,a int) (*State){
	s.S_HasBack = 1
	s.S_BackColorRed, s.S_BackColorGreen, s.S_BackColorBlue, s.S_BackColorAlpha = r, g, b, a

	return s
}

func BackColor4i(r,g,b,a int) (*State){
	return g_curState.BackColor4i(r,g,b,a)
}

func BackColor1v(colors []int) (*State){
	return g_curState.BackColor1v(colors)
}

func (s *State) FontSize(size int) (*State){
	s.S_FontSize = size

	return s
}

func (s *State) Left(left int) (*State){
	s.S_Left = left

	bl.Pos(s.S_Left, s.S_Top)

	return s
}

func (s *State) Top(top int) (*State){
	s.S_Top = top

	bl.Pos(s.S_Left, s.S_Top)

	return s
}

func (s *State) Width(w int) (*State){
	s.S_Width = w

	bl.Dim(s.S_Width, s.S_Height)

	return s
}

func (s *State) Height(h int) (*State){
	s.S_Height = h

	bl.Dim(s.S_Width, s.S_Height)

	return s
}

func (s *State) End() (*State){
	End()

	return s
}

func End() {
	bl.End()
}

