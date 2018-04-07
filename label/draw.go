package label

import "github.com/amortaza/go-ux"

var ux_default *ux.Entity

func init() {
	ux_default = ux.NewEntity("github.com/amortaza/go-dark-ux/label/dark-label.js")
}
