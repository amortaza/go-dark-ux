package button

import "github.com/amortaza/go-ux"

var ux_default       *ux.Entity
var ux_pressed       *ux.Entity
var ux_disabled      *ux.Entity

func init() {
	ux_default = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button.js"}
	ux_pressed = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-pressed.js"}
	ux_disabled = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-button-disabled.js"}
}

