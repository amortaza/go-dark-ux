package border

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-ux"
)

var ux_default *ux.Entity

func init() {
	ux_default = &ux.Entity{File:"c:/go-proj/src/github.com/amortaza/go-dark-ux/border/dark-border.js"}
}

func Wire() {
	bl.CustomRenderer(func(node *bl.Node) {
		ux_default.SetInt("inFill", 0)
		ux_default.SetInt("inRed", 50)
		ux_default.SetInt("inGreen", 0)
		ux_default.SetInt("inBlue", 0)
		ux_default.Draw(0, 0, node.Width, node.Height, "")
	}, false)
}

func Fill(r,g,b int) {
	bl.CustomRenderer(func(node *bl.Node) {
		ux_default.SetInt("inFill", 1)
		ux_default.SetInt("inRed", r)
		ux_default.SetInt("inGreen", g)
		ux_default.SetInt("inBlue", b)
		ux_default.Draw(0, 0, node.Width, node.Height, "")
	}, false)
}

func WireOrFill(r,g,b int, wire bool) {
	bl.CustomRenderer(func(node *bl.Node) {
		//bl.Disp(node)

		if wire {
			ux_default.SetInt("inFill", 0)
		} else {
			ux_default.SetInt("inFill", 1)
		}

		ux_default.SetInt("inRed", r)
		ux_default.SetInt("inGreen", g)
		ux_default.SetInt("inBlue", b)

		ux_default.Draw(0, 0, node.Width, node.Height, "")
	}, wire)
}

