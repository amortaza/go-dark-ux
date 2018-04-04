package checkbox

import "github.com/amortaza/go-ux"

var ux_checkbox__checked *ux.Entity
var ux_checkbox__unchecked *ux.Entity
var ux_checkbox__unchecked_disabled *ux.Entity
var ux_checkbox__checked_disabled *ux.Entity

var ux_plusbox__checked *ux.Entity
var ux_plusbox__unchecked *ux.Entity
var ux_plusbox__unchecked_disabled *ux.Entity
var ux_plusbox__checked_disabled *ux.Entity

func init() {

	ux_checkbox__unchecked = &ux.Entity{Filename:			"github.com/amortaza/go-dark-ux/checkbox/js_checkbox/unchecked.js"}
	ux_checkbox__checked = &ux.Entity{Filename:				"github.com/amortaza/go-dark-ux/checkbox/js_checkbox/checked.js"}
	ux_checkbox__checked_disabled = &ux.Entity{Filename:	"github.com/amortaza/go-dark-ux/checkbox/js_checkbox/checked-disabled.js"}
	ux_checkbox__unchecked_disabled = &ux.Entity{Filename:	"github.com/amortaza/go-dark-ux/checkbox/js_checkbox/unchecked-disabled.js"}

	ux_plusbox__unchecked = &ux.Entity{Filename:		 "github.com/amortaza/go-dark-ux/checkbox/js_plusbox/unchecked.js"}
	ux_plusbox__checked = &ux.Entity{Filename:			 "github.com/amortaza/go-dark-ux/checkbox/js_plusbox/checked.js"}
	ux_plusbox__checked_disabled = &ux.Entity{Filename:	 "github.com/amortaza/go-dark-ux/checkbox/js_plusbox/checked-disabled.js"}
	ux_plusbox__unchecked_disabled = &ux.Entity{Filename:"github.com/amortaza/go-dark-ux/checkbox/js_plusbox/unchecked-disabled.js"}
}








