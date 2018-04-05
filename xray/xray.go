package xray

import (
	"github.com/amortaza/go-bellina"
)

func Use() {

	bl.CustomRenderer(func(node *bl.Node) {

		shadow := bl.EnsureShadowByNode(node)

		ux_xray.SetInt("inFill", 0)
		ux_xray.SetInt("inRed", 255)
		ux_xray.SetInt("inGreen", 255)
		ux_xray.SetInt("inBlue", 0)

		ux_xray.SetInt("inViewWidth", shadow.Width)
		ux_xray.SetInt("inViewHeight", shadow.Height)

		for i := node.Kids.Front(); i != nil; i = i.Next() {

			kid := i.Value.(*bl.Node)

			kidShadow := bl.EnsureShadowByNode(kid)

			ux_xray.Draw(kidShadow.Left, kidShadow.Top, kidShadow.Width, kidShadow.Height, kid.Id)
		}

	}, true)

}

