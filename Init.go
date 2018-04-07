package go_dark_ux

import (
	"github.com/amortaza/go-ux"
	"github.com/amortaza/go-bellina"
)

func Init() {

	ux.Init()

	bl.Register_LifeCycle_After_UserTick_LongTerm( func() {

		if (ux.CheckJsFiles_throttled()) {

			// a JS file was updated
			bl.DirtyAllNodes()
		}
	} )
}

func Uninit() {
	ux.Uninit()
}

