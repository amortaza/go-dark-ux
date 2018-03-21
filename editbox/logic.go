package editbox

import (
	"math"
	"github.com/amortaza/go-hal"
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

func processKeyDown(key hal.KeyboardKey, alt, ctrl, shift bool, state *State) {

	g_cursorCycle = 0

	bt := hal.KeyToBehavior(key, false, true)

	if bt == hal.Key_Behavior_CHAR {
		state.Text_ = insertChar(state.Text_, state.CursorPos, hal.KeyToChar(key, shift, true))
		state.CursorPos++

	} else {
		if key == hal.Key_HOME {
			state.CursorPos = 0
		}
		if key == hal.Key_END {
			state.CursorPos = len(state.Text_)
		}
		if key == hal.Key_DELETE {
			state.Text_ = doDelete(state.Text_, state.CursorPos)
		}
		if key == hal.Key_BACKSPACE {
			state.Text_ = backspace(state.Text_, state.CursorPos)
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == hal.Key_LEFT {
			state.CursorPos = int(math.Max(0, float64(state.CursorPos-1)))
		}
		if key == hal.Key_RIGHT {
			state.CursorPos = int(math.Min(float64(state.CursorPos+1), float64(len(state.Text_))))
		}
	}
}

