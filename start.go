package go_dark_ux

import (
	"github.com/amortaza/go-ux"
)

var button_default *ux.Entity
var button_pressed *ux.Entity
var button_disabled *ux.Entity

var checkbox_checked_enabled *ux.Entity
var checkbox_unchecked_enabled *ux.Entity
var checkbox_unchecked_disabled *ux.Entity
var checkbox_checked_disabled *ux.Entity

var pane *ux.Entity

func Init() {
	button_default = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-default.js"}
	button_pressed = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-pressed.js"}
	button_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-disabled.js"}

	checkbox_unchecked_enabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-unchecked-enabled.js"}
	checkbox_checked_enabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-checked-enabled.js"}
	checkbox_checked_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-checked-disabled.js"}
	checkbox_unchecked_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-unchecked-disabled.js"}

	pane = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-pane.js"}
}

func DrawButton_Default(x,y,w,h int, label string) {
	button_default.Draw(x,y,w,h,label)
}

func DrawButton_Pressed(x,y,w,h int, label string) {
	button_pressed.Draw(x,y,w,h,label)
}

func DrawButton_Disabled(x,y,w,h int, label string) {
	button_disabled.Draw(x,y,w,h,label)
}



func DrawCheckbox_Unchecked_Enabled(x,y,w,h int, label string) {
	checkbox_unchecked_enabled.Draw(x,y,w,h,label)
}

func DrawCheckbox_Checked_Enabled(x,y,w,h int, label string) {
	checkbox_checked_enabled.Draw(x,y,w,h,label)
}

func DrawCheckbox_Checked_Disabled(x,y,w,h int, label string) {
	checkbox_checked_disabled.Draw(x,y,w,h,label)
}

func DrawCheckbox_Unchecked_Disabled(x,y,w,h int, label string) {
	checkbox_unchecked_disabled.Draw(x,y,w,h,label)
}

func DrawPane(x,y,w,h int, label string) {
	pane.Draw(x,y,w,h,label)
}
