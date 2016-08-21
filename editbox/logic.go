package editbox

import (
	"math"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-xel2"
)

func insertChar(src string, pos int, char string) string {

	if pos >= len(src) {
		return src + char
	}

	p1 := src[0:pos]
	p2 := src[pos:]

	return p1 + char + p2
}

func backspace(src string, pos int) string {

	if pos < 1 {
		return src
	}

	p1 := src[0:pos-1]
	p2 := src[pos:]

	return p1 + p2
}

func doDelete(src string, pos int) string {

	if pos >= len(src) {
		return src
	}

	p1 := src[0:pos]
	p2 := src[pos+1:]

	return p1 + p2
}

func processKeyDown(key bl.KeyboardKey, alt, ctrl, shift bool, state *State) {
	g_cursorCycle = 0

	bt := xel.KeyToBehavior(key, false, true)

	if bt == xel.Key_Behavior_CHAR {
		state.Text_ = insertChar(state.Text_, state.CursorPos, xel.KeyToChar(key, shift, true))
		state.CursorPos++

	} else {
		if key == xel.Key_HOME {
			state.CursorPos = 0
		}
		if key == xel.Key_END {
			state.CursorPos = len(state.Text_)
		}
		if key == xel.Key_DELETE {
			state.Text_ = doDelete(state.Text_, state.CursorPos)
		}
		if key == xel.Key_BACKSPACE {
			state.Text_ = backspace(state.Text_, state.CursorPos)
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == xel.Key_LEFT {
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == xel.Key_RIGHT {
			state.CursorPos = int(math.Min(float64(state.CursorPos+1), float64(len(state.Text_))))
		}
	}
}

