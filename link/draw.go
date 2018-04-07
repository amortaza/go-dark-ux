package link

import "github.com/amortaza/go-ux"

var ux_default *ux.Entity

func init() {

	ux_default = ux.NewEntity("github.com/amortaza/go-dark-ux/link/dark-link.js")
}
