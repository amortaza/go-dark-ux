package button

import "github.com/amortaza/go-ux"

var ux_default       *ux.Entity
var ux_pressed       *ux.Entity
var ux_disabled      *ux.Entity

func init() {
	ux_default = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/button/dark-button.js"}
	ux_pressed = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/button/dark-button-pressed.js"}
	ux_disabled = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/button/dark-button-disabled.js"}
}

