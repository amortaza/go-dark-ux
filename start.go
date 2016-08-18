package go_dark_ux

import (
	"github.com/amortaza/go-ux"
)

var button *ux.Entity
var button_pressed *ux.Entity
var button_disabled *ux.Entity

var checkbox_checked *ux.Entity
var checkbox_unchecked *ux.Entity
var checkbox_unchecked_disabled *ux.Entity
var checkbox_checked_disabled *ux.Entity

var radiobox_checked *ux.Entity
var radiobox_unchecked *ux.Entity

var pane *ux.Entity

func Init() {
	ux.Init()

	button = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button.js"}
	button_pressed = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-pressed.js"}
	button_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-disabled.js"}

	checkbox_unchecked = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-unchecked.js"}
	checkbox_checked = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-checked.js"}
	checkbox_checked_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-checked-disabled.js"}
	checkbox_unchecked_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-checkbox-unchecked-disabled.js"}

	radiobox_checked = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-radiobox-checked.js"}
	radiobox_unchecked = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-radiobox-unchecked.js"}

	pane = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-pane.js"}
}

func Uninit() {
	ux.Uninit()
}

/* Button */
func DrawButton(x,y,w,h int, label string) {
	button.Draw(x,y,w,h,label)
}

func DrawButton_Pressed(x,y,w,h int, label string) {
	button_pressed.Draw(x,y,w,h,label)
}

func DrawButton_Disabled(x,y,w,h int, label string) {
	button_disabled.Draw(x,y,w,h,label)
}

/* Checkbox */

func DrawCheckbox_Unchecked(x,y,w,h int, label string) {
	checkbox_unchecked.Draw(x,y,w,h,label)
}

func DrawCheckbox_Checked(x,y,w,h int, label string) {
	checkbox_checked.Draw(x,y,w,h,label)
}

func DrawCheckbox_Checked_Disabled(x,y,w,h int, label string) {
	checkbox_checked_disabled.Draw(x,y,w,h,label)
}

func DrawCheckbox_Unchecked_Disabled(x,y,w,h int, label string) {
	checkbox_unchecked_disabled.Draw(x,y,w,h,label)
}

/* Radiobox */
func DrawRadiobox_Checked(x,y,w,h int, label string) {

	radiobox_checked.Draw(x,y,w,h,label)
}

func DrawRadiobox_Unchecked(x,y,w,h int, label string) {

	radiobox_unchecked.Draw(x,y,w,h,label)
}

/* Pane */

func DrawPane(x,y,w,h int, label string) {
	pane.Draw(x,y,w,h,label)
}
