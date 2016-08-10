package go_dark_ux

import (
	"github.com/amortaza/go-ux"
)

var button_default *ux.Entity
var button_pressed *ux.Entity
var button_disabled *ux.Entity

func Init() {
	button_default = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-default.js"}
	button_pressed = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-pressed.js"}
	button_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-disabled.js"}
}

func DrawButtonDefault(x,y,w,h int, label string) {
	button_default.Draw(x,y,w,h,label)
}

func DrawButtonPressed(x,y,w,h int, label string) {
	button_pressed.Draw(x,y,w,h,label)
}

func DrawButtonDisabled(x,y,w,h int, label string) {
	button_disabled.Draw(x,y,w,h,label)
}
