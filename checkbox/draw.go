package checkbox

import "github.com/amortaza/go-ux"

var ux_checked *ux.Entity
var ux_unchecked *ux.Entity
var ux_unchecked_disabled *ux.Entity
var ux_checked_disabled *ux.Entity

func init() {
	ux_unchecked = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/checkbox/dark-checkbox-unchecked.js"}
	ux_checked = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/checkbox/dark-checkbox-checked.js"}
	ux_checked_disabled = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/checkbox/dark-checkbox-checked-disabled.js"}
	ux_unchecked_disabled = &ux.Entity{Filename:"c:/go-proj/src/github.com/amortaza/go-dark-ux/checkbox/dark-checkbox-unchecked-disabled.js"}
}








