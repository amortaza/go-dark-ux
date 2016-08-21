package border

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-ux"
)

var ux_default *ux.Entity

func init() {
	ux_default = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/js/dark-border.js"}
}

func Draw() {
	bl.CustomRenderer(func(node *bl.Node) {
		ux_default.Draw(0, 0, node.Width, node.Height, "")
	}, true)
}

