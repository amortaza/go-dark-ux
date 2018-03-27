package border

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-ux"
)

// this is the default graphic or look
var ux_default *ux.Entity

func init() {
	ux_default = &ux.Entity{Filename:"github.com/amortaza/go-dark-ux/border/dark-border.js"}
}

func Wire(r, g, b int) {

	bl.CustomRenderer(func(node *bl.Node) {

		shadow := bl.EnsureShadowByNode(node)

		ux_default.SetInt("inFill", 0)
		ux_default.SetInt("inRed", r)
		ux_default.SetInt("inGreen", g)
		ux_default.SetInt("inBlue", b)
		ux_default.Draw(0, 0, shadow.Width, shadow.Height, "")
	}, false)
}

func Wire_TopsKids(r, g, b int) {

	bl.CustomRenderer(func(node *bl.Node) {

		ux_default.SetInt("inFill", 0)
		ux_default.SetInt("inRed", r)
		ux_default.SetInt("inGreen", g)
		ux_default.SetInt("inBlue", b)
		ux_default.Draw(0, 0, node.Width(), node.Height(), "")
	}, true)
}

func Fill(r, g, b int) {

	bl.CustomRenderer(func(node *bl.Node) {

		ux_default.SetInt("inFill", 1)
		ux_default.SetInt("inRed", r)
		ux_default.SetInt("inGreen", g)
		ux_default.SetInt("inBlue", b)
		ux_default.Draw(0, 0, node.Width(), node.Height(), "")
	}, false)
}
